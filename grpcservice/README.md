### run

```bash
protoc --go_out=. --go-grpc_out=. proto/*.proto
protoc --proto_path=. --proto_path=./google --go_out=. --go-grpc_out=. --grpc-gateway_out=. proto/model.proto
```


```go
package main

import (
	"context"
	"google.golang.org/grpc/metadata"
	"grpcservice/database"
	"grpcservice/services"
	"grpcservice/services/model"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func startTest() {
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
func main() {
	go startTest()

	client, err := model.NewClient("localhost:50051")
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	defer client.Close()

	md := metadata.New(map[string]string{
		"authorization": "Bearer 123456",
	})
	ctx := metadata.NewOutgoingContext(context.Background(), md)

	models, err := client.GetModels(ctx)
	if err != nil {
		log.Fatalf("GetModels failed: %v", err)
	}
	log.Println("Models:", models.Models)

	// stream
	stream, err := client.StreamGetModels(ctx)
	if err != nil {
		log.Fatalf("Failed to call StreamModels: %v", err)
	}

	log.Println("Receiving models...")
	for {
		model, err := stream.Recv()
		if err != nil {
			// 检查是否是流结束
			if err.Error() == "EOF" {
				log.Println("Stream ended")
				break
			}
			log.Fatalf("Failed to receive model: %v", err)
		}
		log.Printf("Model received: ID=%s, Name=%s", model.Id, model.Name)
	}
	select {}

}
```