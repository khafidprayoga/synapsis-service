package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	"github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/khafidprayoga/synapsis-service/common/config"
	"github.com/khafidprayoga/synapsis-service/common/conn"
	synapsisv1 "github.com/khafidprayoga/synapsis-service/gen/synapsis/v1"
	"github.com/khafidprayoga/synapsis-service/repository/mongoRepository"
	"github.com/khafidprayoga/synapsis-service/seed"
	"github.com/khafidprayoga/synapsis-service/service"
	"go.uber.org/zap"
	"google.golang.org/grpc/credentials/insecure"
	"net"
	"net/http"

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
		log = config.GetZapLogger()
	)

	log.Info("starting synapsis service")
	log.Info(
		"socket",
		zap.Int("http port", httpPort),
		zap.Int("rpc", rpcPort),
	)

	log.Info("connecting to postgres")
	psqlDB, errConPsql := conn.PostgresConnect()
	if errConPsql != nil {
		log.Fatal("failed to connect to postgres", zap.Error(errConPsql))
	}

	// checks
	psqlDB.Raw("SELECT 1 = 1").Debug()

	if migrateSchema {
		psqlDB.AutoMigrate(
			&synapsisv1.ProductCategory{},
		)
	}

	log.Info("connecting to mongo")
	mongoClient, errConMongo := conn.MongoConnect(context.Background())
	if errConMongo != nil {
		log.Fatal("failed to connect to mongo", zap.Error(errConMongo))
	}
	dbName := mongoClient.Database("synapsis")

	mongoRep, errInitRepo := mongoRepository.NewRepository(log, dbName)
	if errInitRepo != nil {
		log.Fatal("failed to initialize repository", zap.Error(errInitRepo))
	}

	if seedDataSource {
		// seed to postgres data
		productCategory := seed.NewProductCategory()

		log.Info("seeding product category data")
		r := psqlDB.Create(&productCategory)
		if r.Error != nil {
			log.Fatal("failed to seed product category", zap.Error(r.Error))
		}

		// seed to mongo data
		// user domain
		{
			user := seed.NewUser()
			log.Info("seeding user data")
			for _, u := range user {
				_, errInsert := mongoRep.CreateUser(context.Background(), &synapsisv1.CreateUserRequest{
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

	handler := service.NewSynapsisService(log, mongoRep)

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
	flag.IntVar(&rpcPort, "rpc", config.GRPCAddress, "rpc port")
	flag.BoolVar(&migrateSchema, "migrate", false, "migrate schema")
	flag.BoolVar(&seedDataSource, "seed", false, "seed data")

	flag.Parse()
}
