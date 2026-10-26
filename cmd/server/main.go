package main

import (
	"context"
	"log"
	"net"
	"net/url"
	"os"
	"time"

	grpcserver "github.com/AmIAFreeMan/Go-HR-managment-exercise/internal/delivery/grpc"
	"github.com/AmIAFreeMan/Go-HR-managment-exercise/internal/repository"
	"github.com/AmIAFreeMan/Go-HR-managment-exercise/internal/usecase"
	hrv1 "github.com/AmIAFreeMan/Go-HR-managment-exercise/proto/hr/v1"
	jwtgrpc "github.com/auth0/go-jwt-middleware/v3/integrations/grpc"
	"github.com/auth0/go-jwt-middleware/v3/jwks"
	"github.com/auth0/go-jwt-middleware/v3/validator"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, relying on system environment variables")
	}

	ctx := context.Background()

	// Database Connection URL setup
	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		databaseURL = "postgres://postgres:postgres@localhost:5432/hr_system?sslmode=disable"
	}

	dbPool, err := repository.NewPostgresPool(ctx, databaseURL)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer dbPool.Close()
	log.Println("Connected to PostgreSQL successfully!")

	// Auth0 Configuration & JWKS Provider Setup
	issuerString := os.Getenv("AUTH0_ISSUER_URL")
	if issuerString == "" {
		issuerString = "placeholder" // Fallback placeholder
	}
	audience := os.Getenv("AUTH0_AUDIENCE")

	issuerURL, err := url.Parse(issuerString)
	if err != nil {
		log.Fatalf("Failed to parse Auth0 issuer URL: %v", err)
	}

	// Create caching JWKS provider from the jwks package
	provider, err := jwks.NewCachingProvider(
		jwks.WithIssuerURL(issuerURL),
		jwks.WithCacheTTL(15*time.Minute),
	)
	if err != nil {
		log.Fatalf("Failed to create JWKS caching provider: %v", err)
	}

	// Set up the JWT validator using v3
	jwtValidator, err := validator.New(
		validator.WithKeyFunc(provider.KeyFunc),
		validator.WithAlgorithm(validator.RS256),
		validator.WithIssuer(issuerURL.String()),
		validator.WithAudience(audience),
	)
	if err != nil {
		log.Fatalf("Failed to set up JWT validator: %v", err)
	}

	// Create Auth0 gRPC Interceptor
	authInterceptor, err := jwtgrpc.New(
		jwtgrpc.WithValidator(jwtValidator),
	)
	if err != nil {
		log.Fatalf("Failed to create gRPC JWT middleware: %v", err)
	}

	// Initialize Repository & gRPC Service Handler
	employeeRepo := repository.NewEmployeeRepository(dbPool)
	employeeUsecase := usecase.NewEmployeeUsecase(employeeRepo)
	employeeSvc := grpcserver.NewEmployeeServiceServer(employeeUsecase)

	// Create TCP Listener on port 50051
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen on port 50051: %v", err)
	}

	// gRPC Server Auth0 Unary Interceptor
	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(authInterceptor.UnaryServerInterceptor()),
	)

	// Register service
	hrv1.RegisterEmployeeServiceServer(grpcServer, employeeSvc)

	// Enable gRPC Reflection for local debugging
	reflection.Register(grpcServer)

	log.Println("Secure gRPC server is running and listening on port 50051...")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve gRPC: %v", err)
	}
}
