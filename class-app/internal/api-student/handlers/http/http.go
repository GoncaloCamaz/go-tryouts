package http

import (
	classpb "class-app/internal/api-class/handlers/grpc/proto"
	StudentDataModel "class-app/internal/api-student/datamodel"
	StudentDTO "class-app/internal/api-student/dto/requests"
	"github.com/go-pg/pg/v10"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"strconv"
	"time"
)

// Handler struct will hold the DB connection to access the database
type Handler struct {
	DB  *pg.DB
	RPC classpb.ClassServiceClient
}

// Create Student handles the HTTP POST request to create a student
func (h *Handler) CreateStudent(c echo.Context) error {
	studentDTO := new(StudentDTO.StudentCreate)
	if err := c.Bind(studentDTO); err != nil {
		log.Printf("Error binding student DTO: %v", err)
		return c.JSON(http.StatusBadRequest, err)
	}

	//todo get class info to know if class exists

	student := &StudentDataModel.Student{
		Name:    studentDTO.Name,
		Email:   studentDTO.Email,
		ClassId: studentDTO.ClassId,
		Created: time.Now(),
		Updated: time.Now(),
	}

	_, err := h.DB.Model(student).Insert()
	if err != nil {
		log.Printf("Error inserting student: %v", err)
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusCreated, student)
}

func (h *Handler) GetStudent(c echo.Context) error {
	id := c.Param("id")
	var student StudentDataModel.Student
	err := h.DB.Model(&student).Where("id = ?", id).Select()

	if err != nil {
		log.Printf("Database error: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Database error"})
	}

	return c.JSON(http.StatusOK, student)
}

func (h *Handler) UpdateStudent(c echo.Context) error {
	id := c.Param("id")
	studentDTO := new(StudentDTO.StudentUpdate)
	if err := c.Bind(studentDTO); err != nil {
		log.Printf("Error binding student DTO: %v", err)
		return c.JSON(http.StatusBadRequest, err)
	}

	student := &StudentDataModel.Student{
		Name:    studentDTO.Name,
		Email:   studentDTO.Email,
		ClassId: studentDTO.ClassId,
		Updated: time.Now(),
	}

	_, err := h.DB.Model(student).Where("id = ?", id).Update()
	if err != nil {
		log.Printf("Error updating student: %v", err)
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, student)
}

func (h *Handler) DeleteStudent(c echo.Context) error {
	id := c.Param("id")

	num, _ := strconv.Atoi(id)

	student := &StudentDataModel.Student{ID: num}
	_, err := h.DB.Model(student).WherePK().Delete()
	if err != nil {
		log.Printf("Error deleting student: %v", err)
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "Student deleted"})
}

func (h *Handler) ListStudents(c echo.Context) error {
	var students []StudentDataModel.Student
	err := h.DB.Model(&students).Select()

	if err != nil {
		log.Printf("Database error: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Database error"})
	}

	return c.JSON(http.StatusOK, students)
}
