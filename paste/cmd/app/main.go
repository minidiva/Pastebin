package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"pastebin/config"
	"pastebin/internal/db"
	"pastebin/internal/handlers"
	"pastebin/internal/repo"
	"pastebin/internal/service"
	"pastebin/internal/storage"

	configAws "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func main() {

	cfg := config.InitConfig()

	db, err := db.SetupDB(cfg)
	if err != nil {
		log.Fatalf("error connect to DB: %v", err)
	}

	awsCfg, err := configAws.LoadDefaultConfig(
		context.TODO(),
		configAws.WithRegion(cfg.Region),
		configAws.WithCredentialsProvider(
			credentials.NewStaticCredentialsProvider(
				cfg.AccessKey,
				cfg.SecretKey,
				"",
			),
		),
	)

	if err != nil {
		log.Fatalf("failed to load AWSconfig")
	}

	client := s3.NewFromConfig(awsCfg, func(o *s3.Options) {
		o.EndpointResolver = s3.EndpointResolverFromURL(cfg.Endpoint) // https://storage.yandexcloud.kz
		o.UsePathStyle = true                                         // обязательно для Yandex S3
	})

	storage := storage.NewS3Storage(client, cfg.Bucket)

	r := repo.NewPasteRepo(db)

	s := service.NewPasteService(r, storage)

	PasteHandler := handlers.NewPasteHandler(s)

	http.HandleFunc("/health", PasteHandler.CheckHealth)

	http.HandleFunc("/upload", PasteHandler.CreatePaste)

	http.HandleFunc("/download", PasteHandler.GetPaste)

	fmt.Println("Starting server on localhost:8081...")
	if err = http.ListenAndServe(":8081", nil); err != nil {
		log.Fatalf("error starting server... %v", err)
	}
}
