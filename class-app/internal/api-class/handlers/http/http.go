package http

import (
	ClassDataModel "class-app/internal/api-class/datamodel"
	ClassDTO "class-app/internal/api-class/dto/requests"
	"database/sql"
	"errors"
	"github.com/go-pg/pg/v10"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

// Handler struct will hold the DB connection to access the database
type Handler struct {
	DB *pg.DB
}

// CreateClass handles the HTTP POST request to create a class
func (h *Handler) CreateClass(c echo.Context) error {
	classDTO := new(ClassDTO.ClassCreate)
	if err := c.Bind(classDTO); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	class := &ClassDataModel.Class{
		Year:    classDTO.Year,
		Number:  classDTO.Number,
		Created: time.Now(),
		Updated: time.Now(),
	}

	_, err := h.DB.Model(class).Insert()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusCreated, class)
}

// GetClass handles the HTTP GET request to retrieve a class by ID
func (h *Handler) GetClass(c echo.Context) error {
	id := c.Param("id")
	var class ClassDataModel.Class
	err := h.DB.Model(&class).Where("id = ?", id).Select()

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "Class not found"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Database error"})
	}

	return c.JSON(http.StatusOK, class)
}

// GetClassList handles the HTTP GET request to retrieve a class by ID
func (h *Handler) GetClassList(c echo.Context) error {
	// Query the database to get class info
	var classes []ClassDataModel.Class
	err := h.DB.Model(&classes).Select()
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "Class not found"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Database error"})
	}

	return c.JSON(http.StatusOK, classes)
}

// UpdateClass handles the HTTP PUT request to update a class by ID
func (h *Handler) UpdateClass(c echo.Context) error {
	id := c.Param("id")
	classDTO := new(ClassDTO.ClassUpdate)

	if err := c.Bind(classDTO); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	var class ClassDataModel.Class
	err := h.DB.Model(&class).Where("id = ?", id).Select()

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "Class not found"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Database error"})
	}

	class.Year = classDTO.Year
	class.Number = classDTO.Number
	class.Updated = time.Now()

	_, err = h.DB.Model(&class).Where("id = ?", id).Update()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, class)
}

// DeleteClass handles the HTTP DELETE request to delete a class by ID
func (h *Handler) DeleteClass(c echo.Context) error {
	id := c.Param("id")

	_, err := h.DB.Model(&ClassDataModel.Class{}).Where("id = ?", id).Delete()
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "Class not found"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Database error"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Class deleted", "class_id": id})
}
