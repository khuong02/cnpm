package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectMongo(uri string) *mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Thay bằng thông tin từ docker-compose nếu cần
	clientOptions := options.Client().ApplyURI(uri)

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal("Kết nối Mongo thất bại:", err)
	}

	// Ping thử
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("Mongo không phản hồi:", err)
	}

	fmt.Println("✅ Kết nối Mongo thành công!")
	return client
}
