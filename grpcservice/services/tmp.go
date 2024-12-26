package services

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/encoding/protojson"
	"grpcservice/database"
	"grpcservice/middleware"
	mdpb "grpcservice/proto/model"
	nwpb "grpcservice/proto/network"
	tspb "grpcservice/proto/transaction"
	"grpcservice/services/model"
	"grpcservice/services/network"
	"grpcservice/services/transaction"
	"log"
	"net"
	"net/http"
	"strings"
)

func StartServer() {
	if err := database.InitDB(&database.Config{
		DBPath: "./models.db",
		Debug:  true,
	}); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer func() {
		err := database.CloseDB()
		if err != nil {
			log.Fatalf("Failed to close database: %v", err)
		}
	}()
	go runRpcGatewayServers()
}
func runRpcGatewayServers() {
	//go startGRPCServer()
	//go startHTTPGateway()
	// 创建 gRPC 服务器
	// 初始化数据库

	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(middleware.AuthInterceptor),
	)
	mdpb.RegisterModelServiceServer(grpcServer, &model.Server{})

	// 创建 HTTP Gateway
	ctx := context.Background()
	gwmux := runtime.NewServeMux(
		runtime.WithMarshalerOption(
			runtime.MIMEWildcard,
			&runtime.JSONPb{
				MarshalOptions: protojson.MarshalOptions{
					UseProtoNames:   true,
					EmitUnpopulated: true,
				},
				UnmarshalOptions: protojson.UnmarshalOptions{
					DiscardUnknown: true,
				},
			},
		),
		// 添加认证元数据注入器
		runtime.WithMetadata(func(ctx context.Context, r *http.Request) metadata.MD {
			md := make(map[string]string)
			if auth := r.Header.Get("Authorization"); auth != "" {
				md["authorization"] = auth
			}
			return metadata.New(md)
		}),
	)

	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	err := mdpb.RegisterModelServiceHandlerFromEndpoint(ctx, gwmux, ":50051", opts)
	if err != nil {
		log.Fatalf("Failed to register gateway: %v", err)
	}

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.ProtoMajor == 2 && strings.Contains(r.Header.Get("Content-Type"), "application/grpc") {
			grpcServer.ServeHTTP(w, r)
			return
		}
		gwmux.ServeHTTP(w, r)
	})
	handlerWithAuth := middleware.AuthMiddleware(handler)

	h2Handler := h2c.NewHandler(handlerWithAuth, &http2.Server{
		MaxConcurrentStreams: 1000,
	})

	log.Printf("Server listening on :50051")
	if err := http.ListenAndServe(":50051", h2Handler); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
func startGRPCServer() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	mdpb.RegisterModelServiceServer(grpcServer, &model.Server{})
	nwpb.RegisterNetworkServiceServer(grpcServer, &network.Server{})
	tspb.RegisterTransactionServiceServer(grpcServer, &transaction.Server{})

	log.Println("All gRPC services running on 0.0.0.0:50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
func startHTTPGateway() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// 创建 gRPC 客户端连接
	conn, err := grpc.NewClient(
		"localhost:50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("Failed to dial server: %v", err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {

		}
	}(conn)

	mux := runtime.NewServeMux(
		runtime.WithMarshalerOption(
			runtime.MIMEWildcard,
			&runtime.JSONPb{
				MarshalOptions: protojson.MarshalOptions{
					UseProtoNames:   true,
					EmitUnpopulated: true,
				},
				UnmarshalOptions: protojson.UnmarshalOptions{
					DiscardUnknown: true,
				},
			},
		),
	)

	// 注册服务处理器
	err = mdpb.RegisterModelServiceHandlerClient(ctx, mux, mdpb.NewModelServiceClient(conn))
	if err != nil {
		log.Fatalf("Failed to register gateway for ModelService: %v", err)
	}

	//err = nwpb.RegisterNetworkServiceHandlerClient(ctx, mux, nwpb.NewNetworkServiceClient(conn))
	//if err != nil {
	//	log.Fatalf("Failed to register gateway for NetworkService: %v", err)
	//}
	//
	//err = tspb.RegisterTransactionServiceHandlerClient(ctx, mux, tspb.NewTransactionServiceClient(conn))
	//if err != nil {
	//	log.Fatalf("Failed to register gateway for TransactionService: %v", err)
	//}

	// 启动 HTTP 服务器
	httpServer := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	log.Println("HTTP gateway listening on :8080")
	if err := httpServer.ListenAndServe(); err != nil {
		log.Fatalf("Failed to listen and serve: %v", err)
	}
}
