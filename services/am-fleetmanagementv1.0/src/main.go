package main

import (
	"fleetmanagement/infrastructure/database"
	"fleetmanagement/infrastructure/external"
	"fleetmanagement/logic/operations"
	"fmt"
	"net"
	"net/http"
	"os"
	"strings"

	"fleetmanagement/api/controller"
	"fleetmanagement/api/controller/pb"

	"fleetmanagement/logic/model"

	pbRentalManagement "fleetmanagement/infrastructure/external/am-rentalmanagement/client/pb"

	"github.com/improbable-eng/grpc-web/go/grpcweb"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func addTestData(repo *database.PostgresRepository) {

	testFleet := model.Fleet{
		Cars:         []model.Car{},
		Location:     "Karlsruhe",
		FleetManager: "Fred",
		FleetId:      "1",
	}

	repo.AddFleet(testFleet)
}

func main() {
	// Create SQLite in-memory database

	db, err := database.ConnectDB(false)
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

	repo := database.NewPostgresRepository(db)

	addTestData(repo)

	portEnv := os.Getenv("GRPC_PORT")
	if portEnv == "" {
		portEnv = "9001"
	}

	// Create a TCP listener on the specified port
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", portEnv))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	carAPIEnv := os.Getenv("CAR_API_URL")
	if carAPIEnv == "" {
		carAPIEnv = "http://dm-car:80"
	}

	// Rental Management API

	rentalManagementEnv := os.Getenv("RENTAL_MANAGEMENT_URL")
	if rentalManagementEnv == "" {
		rentalManagementEnv = "localhost:9001"
	}

	rentalManagementConn, err := grpc.Dial(rentalManagementEnv, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer rentalManagementConn.Close()

	c := pbRentalManagement.NewRentableCarsCollectionServiceClient(rentalManagementConn)

	rentalManagementApi := external.NewRentalManagementAPI(c)

	carApi := external.NewCarAPI(carAPIEnv)

	fleetOps := operations.NewFleetOperations(repo, carApi, rentalManagementApi)
	carOps := operations.NewCarOperations(repo, carApi)

	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)
	pb.RegisterFleetServiceServer(grpcServer, controller.NewFleetController(fleetOps))
	pb.RegisterCarServiceServer(grpcServer, controller.NewCarController(carOps))

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

	// Get the wrapped gRPC server port from the environment variable
	wrappedPortEnv := os.Getenv("WRAPPED_GRPC_PORT")
	if wrappedPortEnv == "" {
		wrappedPortEnv = "80"
	}

	// Create http server for wrapped gRPC
	httpServer := http.Server{
		Addr:    fmt.Sprintf(":%s", wrappedPortEnv),
		Handler: http.HandlerFunc(wrappedServer.ServeHTTP),
	}

	// Start serving incoming gRPC requests
	go func() {
		log.Println(fmt.Sprintf("gRPC server is running on port %s", portEnv))
		if err := grpcServer.Serve(listener); err != nil {
			log.Fatalf("Failed to serve: %v", err)
		}
	}()

	// Start listening and serving using wrapped server
	log.Println(fmt.Sprintf("Wrapped gRPC server is running on port %s", wrappedPortEnv))
	if err := httpServer.ListenAndServe(); err != nil {
		log.Fatalf("failed starting http server: %v", err)
	}

}
