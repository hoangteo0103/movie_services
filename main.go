package main

import (
	"log"
	"net"
	"os"
	"store"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
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

	querier := store.New(db)

	// [Set up tcp]
	s := grpc.NewServer()
	lis, err := net.Listen("tcp", ":"+os.Getenv("GRPC_PORT"))
	if err != nil {
		log.Fatalln("Failed to listen", err)
	}

	// [grpc Server]
	// Launch grpc Server
	go func() {
		log.Println("Serving gRPC on " + os.Getenv("SERVER_HOST") + ":" + os.Getenv("GRPC_PORT"))
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Failed to serve GRPC server %v", err)
		}
	}()

}
