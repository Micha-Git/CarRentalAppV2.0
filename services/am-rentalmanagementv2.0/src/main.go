package main

import (
	"fmt"
	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"github.com/joho/godotenv"
	"net"
	"net/http"
	"os"
	"rentalmanagement/api/controller"
	"rentalmanagement/api/controller/pb"
	"rentalmanagement/infrastructure/database"
	"rentalmanagement/infrastructure/external"
	"rentalmanagement/logic/model"
	"rentalmanagement/logic/operations"
	"strings"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	// Load environment
	err := godotenv.Load()
	if err != nil {
		log.Warnf("Couldn't load .env file: %v", err)
	}

	// Connect to postgres database
	_, err = database.ConnectDB()
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
		return
	}

	// Migrate models to database tables
	err = database.Migrate()
	if err != nil {
		log.Fatalf("Failed to migrate tables: %v", err)
		return
	}

	// Get the gRPC server port from the environment variable
	portEnv := os.Getenv("GRPC_PORT")
	if portEnv == "" {
		portEnv = "8080"
	}

	// Create a TCP listener on the specified port
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", portEnv))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// Get the Car API URL from the environment variable
	carAPIEnv := os.Getenv("CAR_API_URL")
	if carAPIEnv == "" {
		carAPIEnv = "http://dm-car:80"
	}

	// Initialize repositories
	postgresRepository := database.NewPostgresRepository(database.DB)
	var rentalRepository model.RentalRepositoryInterface = &postgresRepository
	var rentableCarRepository model.RentableCarRepositoryInterface = &postgresRepository
	carRepository := external.NewCarAPI(carAPIEnv)

	// Create operations
	customerOps := operations.NewCustomerOperations(rentalRepository, rentableCarRepository)
	rentalsOps := operations.NewRentalsCollectionOperations(rentalRepository, rentableCarRepository)
	rentableCarOps := operations.NewRentableCarsCollectionOperations(rentalRepository, rentableCarRepository, carRepository)

	// Create a new gRPC server
	grpcServer := grpc.NewServer()

	// Register reflection and gRPC services
	reflection.Register(grpcServer)
	pb.RegisterRentalsCollectionServiceServer(grpcServer, controller.NewRentalsCollectionController(rentalsOps))
	pb.RegisterCustomerServiceServer(grpcServer, controller.NewCustomerController(customerOps))
	pb.RegisterRentableCarsCollectionServiceServer(grpcServer, controller.NewRentableCarsCollectionController(rentableCarOps))

	// Get the wrapped gRPC server port from the environment variable
	wrappedPortEnv := os.Getenv("WRAPPED_GRPC_PORT")

	if wrappedPortEnv != "" {
		// Get the allowed origins list from the environment variable
		allowedOriginsEnv := os.Getenv("ALLOWED_ORIGINS")
		if allowedOriginsEnv == "" {
			allowedOriginsEnv = "*"
		}

		// Create allowed origins map
		allowedOrigins := strings.Split(allowedOriginsEnv, ",")
		allowedOriginsMap := make(map[string]bool)
		for _, allowedOrigin := range allowedOrigins {
			allowedOriginsMap[allowedOrigin] = true
		}

		// Create allowed origin func for wrapped gRPC server
		allowedOriginFunc := grpcweb.WithOriginFunc(func(origin string) bool {
			return allowedOriginsMap[origin] || allowedOriginsMap["*"]
		})

		// Wrap gRPC server
		wrappedServer := grpcweb.WrapServer(grpcServer, allowedOriginFunc)

		// Create http server for wrapped gRPC
		httpServer := http.Server{
			Addr:    fmt.Sprintf(":%s", wrappedPortEnv),
			Handler: http.HandlerFunc(wrappedServer.ServeHTTP),
		}

		// Start listening and serving using wrapped server
		go func() {
			log.Println(fmt.Sprintf("Wrapped gRPC server is running on port %s", wrappedPortEnv))
			if err := httpServer.ListenAndServe(); err != nil {
				log.Fatalf("Failed starting http server: %v", err)
			}
		}()
	}

	// Start serving incoming gRPC requests
	log.Println(fmt.Sprintf("gRPC server is running on port %s", portEnv))
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
