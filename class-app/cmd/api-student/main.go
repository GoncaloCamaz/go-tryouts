package main

import (
	classpb "class-app/internal/api-class/handlers/grpc/proto"
	handler "class-app/internal/api-student/handlers/http"
	utils "class-app/pkg/database"
	"fmt"
	"github.com/go-pg/pg/v10"
	"github.com/labstack/echo/v4"
	"google.golang.org/grpc"
	"log"
	"os"
)

var classClient classpb.ClassServiceClient

// Create necessary tables if they do not exist
func createTables(db *pg.DB) {
	if exists, err := utils.TableExists(db, "students"); err != nil {
		log.Fatalf("Error checking students table existence: %v", err)
	} else if !exists {
		createClassesTable := `
        CREATE TABLE students (
			id SERIAL PRIMARY KEY,
			name VARCHAR(64) NOT NULL,
			email VARCHAR(64) NOT NULL,
			class_id INT REFERENCES classes(id) ON DELETE CASCADE,
			updated TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			created TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		);`
		if _, err := db.Exec(createClassesTable); err != nil {
			log.Fatalf("Error creating students table: %v", err)
		}
		log.Println("Students table created.")
	} else {
		log.Println("Students table already exists.")
	}
}

func ConnectDB() *pg.DB {
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("POSTGRES_DB")
	addr := os.Getenv("POSTGRES_HOST")

	opts := &pg.Options{
		User:     user,
		Password: password,
		Addr:     fmt.Sprintf("%v:%d", addr, 5432),
		Database: dbName,
	}
	db := pg.Connect(opts)
	return db
}

func main() {
	// Set up gRPC connection to api-class
	conn, err := grpc.Dial("api-class:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect to api-class: %v", err)
	}
	
	classClient = classpb.NewClassServiceClient(conn)
	defer conn.Close()

	db := ConnectDB()
	go createTables(db)

	h := handler.Handler{DB: db, RPC: classClient}

	// Start HTTP server for student API
	e := echo.New()

	e.GET("/student/list", h.ListStudents)
	e.GET("/student/:id/get", h.GetStudent)
	e.POST("/student/create", h.CreateStudent)
	e.PUT("/student/:id/update", h.UpdateStudent)
	e.DELETE("/student/:id/delete", h.DeleteStudent)

	log.Fatal(e.Start(":8082"))
}
