package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"targetApi/internal/delivery"
	"targetApi/internal/transport"
	"targetApi/internal/db"
	"targetApi/internal/listener"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Initialize the database connection
	postgresConnection := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_DB"),
	)
	pg, err := sql.Open("postgres", postgresConnection)
	if err != nil {
		log.Fatalf("Postgres connection failed: %v", err)
	}

	// use the connection pool + redis cache
	cache := db.NewCache(os.Getenv("REDIS_HOST"))

	repo := db.NewPostgresRepo(pg)
	campaigns, _ := repo.GetActiveCampaigns()
	rules, _ := repo.GetTargetingRules()
	ctx := context.Background()
	cache.SetCampaigns(ctx, campaigns)
	cache.SetRules(ctx, rules)

	// Listen for changes in the database
	go listener.ListenForCampaignChanges(postgresConnection, repo, cache)

	// Initialize the repository
	svc := delivery.NewService(cache)
	handler := transport.NewHTTPHandler(svc)

	log.Printf("api listening on :%s", port)
	log.Fatal(http.ListenAndServe(":"+port, handler))
}
