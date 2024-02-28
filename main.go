package main

import (
	"context"
	"log"
	"movie_services/api"
	"net"
	"net/http"
	"os"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	//prepare database
	dbHost := os.Getenv("DATABASE_HOST")
	dbPort := os.Getenv("DATABASE_PORT")
	dbUser := os.Getenv("DATABASE_USER")
	dbPass := os.Getenv("DATABASE_PASS")
	dbName := os.Getenv("DATABASE_NAME")

	connStr := "user=" + dbUser + " dbname=" + dbName + " sslmode=disable password=" + dbPass + " host=" + dbHost + " port=" + dbPort

	db, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		log.Fatalln(err)
	}

	defer db.Close()

	// Test the connection to the database
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	} else {
		log.Println("Successfully Connected")
	}

	//	querier := store.New(db)

	// [Set up tcp]
	s := grpc.NewServer()
	lis, err := net.Listen("tcp", ":"+os.Getenv("GRPC_PORT"))
	if err != nil {
		log.Fatalln("Failed to listen", err)
	}

	// [grpc Server]
	// Launch grpc Server
	go func() {
		log.Println("Serving gRPC on " + os.Getenv("HOST") + ":" + os.Getenv("GRPC_PORT"))
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Failed to serve GRPC server %v", err)
		}
	}()

	go func() {
		log.Println("Serving gRPC on " + os.Getenv("HOST") + ":" + os.Getenv("GRPC_PORT"))
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Failed to serve gRPC server: %v", err)
		}
	}()

	// Test connection
	conn, err := grpc.DialContext(
		context.Background(),
		os.Getenv("HOST")+":"+os.Getenv("GRPC_PORT"),
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalln("Failed to dial server:", err)
		log.Println(conn)
	}

	// [HTTP Server conf]
	gwmux := runtime.NewServeMux()

	api.RegisterHealthCheckServiceHandler(context.Background(), gwmux, conn)
	// err = api.RegisterUserServiceHandler(context.Background(), gwmux, conn)
	// if err != nil {
	// 	log.Fatalln("Failed to register gateway:", err)
	// }

	gwServer := &http.Server{
		Addr:    ":" + os.Getenv("HTTP_PORT"),
		Handler: gwmux,
	}

	log.Println("Serving gRPC-Gateway for REST on " + os.Getenv("HOST") + ":" + os.Getenv("HTTP_PORT"))
	if err := gwServer.ListenAndServe(); err != nil {
		log.Fatalf("Failed to serve gRPC-Gateway server: %v", err)
	}

}
