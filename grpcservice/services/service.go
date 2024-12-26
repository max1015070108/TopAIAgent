package services

import (
	"context"
	"fmt"
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
	"grpcservice/services/model"
	"log"
	"net/http"
	"strings"
)

// Options 服务配置选项
type Options struct {
	Address     string            // 服务地址
	DBConfig    *database.Config  // 数据库配置
	AuthEnabled bool              // 是否启用认证
	Debug       bool              // 是否启用调试模式
	MaxStreams  uint32            // 最大并发流
	Metadata    map[string]string // 额外的元数据
}

func DefaultOptions() *Options {
	return &Options{
		Address: ":50051",
		DBConfig: &database.Config{
			DBPath: "./models.db",
			Debug:  false,
		},
		AuthEnabled: true,
		MaxStreams:  1000,
	}
}

type Server struct {
	opts       *Options
	grpcServer *grpc.Server
	gwMux      *runtime.ServeMux
	httpServer *http.Server
}

func NewServer(opts *Options) *Server {
	if opts == nil {
		opts = DefaultOptions()
	}
	return &Server{
		opts: opts,
	}
}

// Start 启动服务
func (s *Server) Start(ctx context.Context) error {
	// 初始化数据库
	if err := database.InitDB(s.opts.DBConfig); err != nil {
		return fmt.Errorf("failed to initialize database: %w", err)
	}

	// 创建 gRPC 服务器
	var serverOpts []grpc.ServerOption
	if s.opts.AuthEnabled {
		serverOpts = append(serverOpts, grpc.UnaryInterceptor(middleware.AuthInterceptor))
	}
	s.grpcServer = grpc.NewServer(serverOpts...)
	mdpb.RegisterModelServiceServer(s.grpcServer, &model.Server{})

	// 创建 HTTP Gateway
	s.gwMux = runtime.NewServeMux(
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
		runtime.WithMetadata(s.metadataAnnotator),
	)

	// 注册 gateway handler
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	if err := mdpb.RegisterModelServiceHandlerFromEndpoint(ctx, s.gwMux, s.opts.Address, opts); err != nil {
		return fmt.Errorf("failed to register gateway: %w", err)
	}

	s.httpServer = &http.Server{
		Addr:    s.opts.Address,
		Handler: s.createHandler(),
	}

	// 启动服务器
	go func() {
		log.Printf("Server listening on %s", s.opts.Address)
		if err := s.httpServer.ListenAndServe(); err != http.ErrServerClosed {
			log.Printf("Failed to serve: %v", err)
		}
	}()

	// 监听上下文取消
	go func() {
		<-ctx.Done()
		s.Stop()
	}()

	return nil
}

// Stop 停止服务
func (s *Server) Stop() error {
	if s.httpServer != nil {
		if err := s.httpServer.Shutdown(context.Background()); err != nil {
			log.Printf("Failed to shutdown HTTP server: %v", err)
		}
	}

	if s.grpcServer != nil {
		s.grpcServer.GracefulStop()
	}

	return database.CloseDB()
}

// metadataAnnotator 处理元数据
func (s *Server) metadataAnnotator(ctx context.Context, r *http.Request) metadata.MD {
	md := make(map[string]string)

	// 处理认证头
	if auth := r.Header.Get("Authorization"); auth != "" {
		md["authorization"] = auth
	}

	// 添加配置中的额外元数据
	for k, v := range s.opts.Metadata {
		md[k] = v
	}

	return metadata.New(md)
}

func (s *Server) createHandler() http.Handler {
	// 主处理器
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 检测是否为 gRPC 请求
		if r.ProtoMajor == 2 && strings.Contains(r.Header.Get("Content-Type"), "application/grpc") {
			s.grpcServer.ServeHTTP(w, r)
			return
		}

		// 将 gwMux 包装为 http.Handler
		var httpHandler http.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			s.gwMux.ServeHTTP(w, r)
		})

		// 应用中间件
		if s.opts.AuthEnabled {
			httpHandler = middleware.AuthMiddleware(httpHandler)
		}

		httpHandler.ServeHTTP(w, r)
	})

	// 使用 h2c 支持 HTTP/2 的非 TLS 连接
	return h2c.NewHandler(handler, &http2.Server{
		MaxConcurrentStreams: s.opts.MaxStreams,
	})
}
