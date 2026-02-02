package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/jirawatchn/resume-backend/internal/resume"
	"github.com/jirawatchn/resume-backend/internal/route"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	// ตัด cmd/app ออก เพราะเราย้าย internal ออกมาไว้ที่ root แล้ว
)

func ConnectDB() (*mongo.Client, error) {
	_ = godotenv.Load()

	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatal("MONGODB_URI is not set in .env file")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	if err := client.Ping(ctx, nil); err != nil {
		return nil, err
	}

	log.Println("Connected to MongoDB Atlas successfully!")
	return client, nil
}

func main() {
	mongoClient, err := ConnectDB()
	if err != nil {
		log.Fatal("Could not connect to MongoDB:", err)
	}

	// เรียกใช้ผ่านชื่อ package resume และ route ตรงๆ
	repo := resume.NewRepository(mongoClient)
	svc := resume.NewService(repo)
	h := resume.NewHandler(svc)

	r := route.NewRouter(h)

	log.Println("Server starting on :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
