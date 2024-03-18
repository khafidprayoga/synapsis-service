package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	"github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	commonConf "github.com/khafidprayoga/synapsis-service/common/config"
	commonConn "github.com/khafidprayoga/synapsis-service/common/conn"
	synapsisv1 "github.com/khafidprayoga/synapsis-service/gen/synapsis/v1"
	"github.com/khafidprayoga/synapsis-service/repository/mongoRepository"
	"github.com/khafidprayoga/synapsis-service/repository/postgresRepository"
	"github.com/khafidprayoga/synapsis-service/repository/redisRepository"
	"github.com/khafidprayoga/synapsis-service/seed"
	"github.com/khafidprayoga/synapsis-service/service"
	"go.uber.org/zap"
	"google.golang.org/grpc/credentials/insecure"
	"net"
	"net/http"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var (
	migrateSchema  bool
	seedDataSource bool
	httpPort       int
	rpcPort        int
)

func main() {
	flagParse()
	var (
		log = commonConf.GetZapLogger()
	)

	log.Info("starting synapsis service")
	log.Info(
		"socket",
		zap.Int("http port", httpPort),
		zap.Int("rpc", rpcPort),
	)

	log.Info("connecting to postgres")
	psqlDB, errConPsql := commonConn.PostgresConnect()
	if errConPsql != nil {
		log.Fatal("failed to connect to postgres", zap.Error(errConPsql))
	}

	psqlDBDebug := psqlDB.Debug()
	// checks
	psqlDBDebug.Raw("SELECT 1 = 1")

	if migrateSchema {
		psqlDBDebug.Raw(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp"`)
		log.Info("migrating postgres schema")
		psqlDBDebug.AutoMigrate(
			&synapsisv1.ProductCategory{},
			&synapsisv1.Product{},
			&synapsisv1.ProductCategoryRelation{},
		)
	}

	log.Info("connecting to mongo")
	mongoClient, errConMongo := commonConn.MongoConnect(context.Background())
	if errConMongo != nil {
		log.Fatal("failed to connect to mongo", zap.Error(errConMongo))
	}
	dbName := mongoClient.Database("synapsis")

	///// Repository
	userRepo, errInitUserRepo := mongoRepository.NewRepository(log, dbName)
	if errInitUserRepo != nil {
		log.Fatal("failed to initialize user repository", zap.Error(errInitUserRepo))
	}

	productRepo, errInitProductRepo := postgresRepository.NewRepository(log, psqlDB)
	if errInitProductRepo != nil {
		log.Fatal("failed to initialize product repository", zap.Error(errInitProductRepo))
	}

	authRepo, errInitAuthRepo := redisRepository.NewRepository(log)
	if errInitProductRepo != nil {
		log.Fatal("failed to initialize auth repository", zap.Error(errInitAuthRepo))
	}

	repositoryDependency := service.ServiceRepository{
		User:    userRepo,
		Product: productRepo,
		Auth:    authRepo,
	}
	///// Repository

	if seedDataSource {
		// seed to postgres data
		productCategory := seed.NewProductCategory()

		log.Info("seeding product category data")
		r := psqlDBDebug.Create(&productCategory)

		if r.Error != nil {
			if !strings.Contains(r.Error.Error(), "duplicate key value violates unique constraint") {
				log.Fatal("failed to seed product category", zap.Error(r.Error))
			}
		}

		// seed to mongo data
		// user domain
		{
			user := seed.NewUser()
			log.Info("seeding user data")
			for _, u := range user {
				_, errInsert := userRepo.CreateUser(context.Background(), &synapsisv1.CreateUserRequest{
					FullName: u.GetFullName(),
					Email:    u.GetEmail(),
					Password: u.GetPassword(),
					Dob:      u.GetDob(),
				}, true)
				if errInsert != nil {
					log.Fatal("failed to seed user", zap.Error(errInsert))
				}
			}
		}
	}

	handler := service.NewSynapsisService(log, repositoryDependency)

	listen, err := net.Listen("tcp", fmt.Sprintf(":%v", rpcPort))
	if err != nil {
		panic(err)
	}

	bootMsg := fmt.Sprintf(
		"[RPC] service synapsis running on %v", rpcPort,
	)
	log.Info(bootMsg)

	s := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			grpc_zap.UnaryServerInterceptor(log),
			grpc_recovery.UnaryServerInterceptor(),
		),
	)

	synapsisv1.RegisterSynapsisServiceServer(s, handler)
	reflection.Register(s)

	serveRpc := func() {
		if e := s.Serve(listen); e != nil {
			log.Fatal("failed to serve", zap.Error(e))
		}
	}

	if httpPort > 0 {
		go serveRpc()
	} else {
		serveRpc()
	}

	// run http service gateway
	var (
		mux = runtime.NewServeMux()
	)

	rpcAddress := fmt.Sprintf(":%v", rpcPort)
	rpcConn, errDial := grpc.
		DialContext(context.Background(), rpcAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if errDial != nil {
		log.Fatal("failed to dial", zap.Error(errDial))
	}

	log.Info("registering grpc service for HTTP endpoint")
	errRegister := synapsisv1.RegisterSynapsisServiceHandler(context.Background(), mux, rpcConn)
	if errRegister != nil {
		log.Fatal("failed to register grpc service", zap.Error(errRegister))
	}

	startProxyMsg := fmt.Sprintf(
		"[HTTP] serving proxy on %v ", httpPort,
	)

	log.Info(startProxyMsg)
	serveErr := http.ListenAndServe(fmt.Sprintf(":%v", httpPort), mux)
	if serveErr != nil {
		log.Fatal("failed to serve", zap.Error(serveErr))
	}
}

func flagParse() {
	flag.IntVar(&httpPort, "http", 0, "http port")
	flag.IntVar(&rpcPort, "rpc", commonConf.GRPCAddress, "rpc port")
	flag.BoolVar(&migrateSchema, "migrate", false, "migrate schema")
	flag.BoolVar(&seedDataSource, "seed", false, "seed data")

	flag.Parse()
}
