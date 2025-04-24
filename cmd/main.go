package main

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"os"

	"targetApi/internal/db"
	"targetApi/internal/delivery"
	"targetApi/internal/listener"
	"targetApi/internal/transport"

	"github.com/go-kit/log"
	"github.com/sony/gobreaker"
	"golang.org/x/time/rate"

	"targetApi/internal/middleware"
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
		fmt.Printf("Postgres connection failed: %v", err)
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

	// Initialize the service and HTTP handler
	logger := log.NewJSONLogger(os.Stdout)
	svc := delivery.NewService(cache)
	svc = middleware.LoggingMiddleware(logger)(svc)
	svc = middleware.CircuitBreakerMiddleware(gobreaker.NewCircuitBreaker(gobreaker.Settings{Name: "GetCampaigns"}))(svc)
	rateLimiter := rate.NewLimiter(10, 20) // 10 req/sec, burst of 20
	handler := middleware.RecoveryMiddleware(
		middleware.RateLimiterMiddleware(rateLimiter)(
			transport.NewHTTPHandler(svc)))

	fmt.Printf("api listening on :%s \n", port)
	fmt.Println(http.ListenAndServe(":"+port, handler))
}
