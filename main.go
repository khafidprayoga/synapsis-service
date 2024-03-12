package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	"github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/khafidprayoga/synapsis-service/common/config"
	synapsisv1 "github.com/khafidprayoga/synapsis-service/gen/synapsis/v1"
	"github.com/khafidprayoga/synapsis-service/service"
	"github.com/khafidprayoga/synapsis-service/service/repository/mongoRepository"
	"go.uber.org/zap"
	"google.golang.org/grpc/credentials/insecure"
	"net"
	"net/http"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var (
	httpPort int
	rpcPort  int
)

func main() {
	flagParse()
	var (
		log = config.GetZapLogger()
	)

	log.Info("starting Synpsis service")
	log.Info(
		"socket",
		zap.Int("http port", httpPort),
		zap.Int("rpc", rpcPort),
	)
	listen, err := net.Listen("tcp", fmt.Sprintf(":%v", rpcPort))
	if err != nil {
		panic(err)
	}

	bootMsg := fmt.Sprintf(
		"[RPC] service synapsis running on %v", rpcPort,
	)

	log.Info(bootMsg)

	rep, errInitRepo := mongoRepository.NewRepository(log)
	if errInitRepo != nil {
		log.Fatal("failed to initialize repository", zap.Error(errInitRepo))
	}

	handler := service.NewSynapsisService(log, rep)
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
		ctx, cancel = context.WithCancel(context.Background())
		mux         = runtime.NewServeMux()
	)
	defer cancel()

	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	log.Info("registering grpc service for HTTP endpoint")
	errSetupGW := synapsisv1.RegisterSynapsisServiceHandlerFromEndpoint(
		ctx,
		mux, fmt.Sprintf(":%v", httpPort),
		opts,
	)

	if errSetupGW != nil {
		panic(errSetupGW)
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
	flag.Parse()
}
