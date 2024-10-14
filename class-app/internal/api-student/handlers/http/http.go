package http

import (
	classpb "class-app/internal/api-class/handlers/grpc/proto"
	StudentDataModel "class-app/internal/api-student/datamodel"
	requestDTO "class-app/internal/api-student/dto/requests"
	responseDTO "class-app/internal/api-student/dto/responses"
	"class-app/pkg/utils"
	"github.com/go-pg/pg/v10"
	"github.com/labstack/echo/v4"
	"google.golang.org/protobuf/types/known/emptypb"
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
	studentDTO := new(requestDTO.StudentCreate)
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

	resp, err := h.RPC.GetClassInfo(c.Request().Context(), &classpb.ClassRequest{ClassId: int64(int32(student.ClassId))})
	if err != nil {
		log.Printf("Error getting class info: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Error getting class info"})
	}

	var studentDTO = responseDTO.Student{
		ID:          student.ID,
		Name:        student.Name,
		Email:       student.Email,
		ClassId:     student.ClassId,
		ClassNumber: int(resp.Number),
		ClassYear:   resp.Year,
		Created:     student.Created,
		Updated:     student.Updated,
	}

	return c.JSON(http.StatusOK, studentDTO)
}

func (h *Handler) UpdateStudent(c echo.Context) error {
	id := c.Param("id")
	studentDTO := new(requestDTO.StudentUpdate)
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

	resp, err := h.RPC.GetClassList(c.Request().Context(), &emptypb.Empty{})
	if err != nil {
		log.Printf("Error getting class list: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Error getting class list"})
	}

	classList := resp.Classes
	classMap := utils.Reduce(classList, func(classMap map[int32]*classpb.ClassResponse, class *classpb.ClassResponse) map[int32]*classpb.ClassResponse {
		classMap[class.Id] = class
		return classMap
	}, make(map[int32]*classpb.ClassResponse))

	var studentsDTO []responseDTO.Student
	for _, student := range students {
		class, ok := classMap[int32(student.ClassId)]
		if !ok {
			log.Printf("Class with ID: %d not found", student.ClassId)
			return c.JSON(http.StatusNotFound, map[string]string{"error": "Class not found"})
		}

		studentDTO := responseDTO.Student{
			ID:          student.ID,
			Name:        student.Name,
			Email:       student.Email,
			ClassId:     student.ClassId,
			ClassNumber: int(class.Number),
			ClassYear:   class.Year,
		}
		studentsDTO = append(studentsDTO, studentDTO)
	}
	return c.JSON(http.StatusOK, studentsDTO)
}
