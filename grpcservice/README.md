### run

```bash
protoc --go_out=. --go-grpc_out=. proto/*.proto
protoc --proto_path=. --proto_path=./google --go_out=. --go-grpc_out=. --grpc-gateway_out=. proto/model.proto
```


```go
func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// 创建服务实例
	srv := services.NewServer(&services.Options{
		Address: ":50051",
		DBConfig: &database.Config{
			DBPath: "./models.db",
			Debug:  true,
		},
		AuthEnabled: true,
		MaxStreams:  1000,
		Metadata: map[string]string{
			"service-version": "1.0.0",
		},
	})

	// 启动服务
	if err := srv.Start(ctx); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

	// 处理优雅关闭
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	// 给服务一些时间来完成正在处理的请求
	_, shutdownCancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer shutdownCancel()

	cancel() // 触发服务关闭
	if err := srv.Stop(); err != nil {
		log.Printf("Error during shutdown: %v", err)
	}
}
```