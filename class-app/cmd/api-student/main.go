package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"

	classpb "class-app/internal/api-class/handlers/grpc/proto" // Update with the correct import path
	"github.com/labstack/echo/v4"
	"google.golang.org/grpc"
)

var classClient classpb.ClassServiceClient

func main() {
	// Set up gRPC connection to api-class
	conn, err := grpc.Dial("api-class:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect to api-class: %v", err)
	}
	defer conn.Close()
	classClient = classpb.NewClassServiceClient(conn)

	// Start HTTP server for student API
	e := echo.New()

	e.POST("/student", func(c echo.Context) error {
		// Handle create student
		return c.JSON(http.StatusCreated, map[string]string{"message": "Student created"})
	})

	e.GET("/student/:id", func(c echo.Context) error {
		// Handle get student by ID and fetch class info via gRPC
		id, _ := strconv.Atoi(c.Param("id"))
		classReq := &classpb.ClassRequest{ClassId: int64(id)}

		// Call gRPC to get class info
		classRes, err := classClient.GetClassInfo(context.Background(), classReq)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"student_id":   id,
			"student_name": fmt.Sprintf("Student %d", id),
			"class_id":     classRes.Id,
			"class_name":   classRes.Number,
		})
	})

	e.PUT("/student/:id", func(c echo.Context) error {
		// Handle update student by ID
		return c.JSON(http.StatusOK, map[string]string{"message": "Student updated"})
	})

	e.DELETE("/student/:id", func(c echo.Context) error {
		// Handle delete student by ID
		return c.JSON(http.StatusOK, map[string]string{"message": "Student deleted"})
	})

	log.Fatal(e.Start(":8081"))
}
