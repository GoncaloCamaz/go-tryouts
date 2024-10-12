package main

import (
	grpc2 "class-app/internal/api-class/handlers/grpc"
	handler "class-app/internal/api-class/handlers/http"
	utils "class-app/pkg/database"
	"fmt"
	"github.com/go-pg/pg/v10"
	"github.com/labstack/echo/v4"
	"log"
	"os"
)

// Create necessary tables if they do not exist
func createTables(db *pg.DB) {
	if exists, err := utils.TableExists(db, "classes"); err != nil {
		log.Fatalf("Error checking classes table existence: %v", err)
	} else if !exists {
		createClassesTable := `
        CREATE TABLE classes (
            id SERIAL PRIMARY KEY,
            year VARCHAR(256) NOT NULL,
            number INT NOT NULL,
            updated TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
            created TIMESTAMP DEFAULT CURRENT_TIMESTAMP
        );`
		if _, err := db.Exec(createClassesTable); err != nil {
			log.Fatalf("Error creating classes table: %v", err)
		}
		log.Println("Classes table created.")
	} else {
		log.Println("Classes table already exists.")
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
	// Start gRPC server for class API
	go func() {
		grpc2.StartServer(":50051")
	}()

	db := ConnectDB()
	go createTables(db)
	defer db.Close()

	// Create handler instance with DB connection
	h := handler.Handler{DB: db}
	// Initialize Echo server
	e := echo.New()
	// Define HTTP routes
	e.POST("/class/create", h.CreateClass)
	e.GET("/class/:id/get", h.GetClass)
	e.GET("/class/list", h.GetClassList)
	e.PUT("/class/:id/update", h.UpdateClass)
	e.DELETE("/class/:id/delete", h.DeleteClass)

	log.Fatal(e.Start(":8081"))
}
