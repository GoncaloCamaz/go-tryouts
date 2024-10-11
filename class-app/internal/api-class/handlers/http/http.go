package http

import (
	ClassDataModel "class-app/internal/api-class/datamodel"
	"database/sql"
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
	class := new(ClassDataModel.Class)
	if err := c.Bind(class); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	class.Created = time.Now()
	class.Updated = time.Now()
	_, err := h.DB.Model(class).Insert()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusCreated, class)
}

// GetClass handles the HTTP GET request to retrieve a class by ID
func (h *Handler) GetClass(c echo.Context) error {
	id := c.Param("id")

	// Query the database to get class info
	var name string

	return c.JSON(http.StatusOK, map[string]string{"class_id": id, "class_name": name})
}

// GetClassList handles the HTTP GET request to retrieve a class by ID
func (h *Handler) GetClassList(c echo.Context) error {
	// Query the database to get class info
	var classes []ClassDataModel.Class
	err := h.DB.Model(&classes).Select()
	if err != nil {
		if err == sql.ErrNoRows {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "Class not found"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Database error"})
	}

	return c.JSON(http.StatusOK, classes)
}

// UpdateClass handles the HTTP PUT request to update a class by ID
func (h *Handler) UpdateClass(c echo.Context) error {
	id := c.Param("id")
	// Simulate updating a class
	// You would update class info in the database here
	return c.JSON(http.StatusOK, map[string]string{"message": "Class updated", "class_id": id})
}

// DeleteClass handles the HTTP DELETE request to delete a class by ID
func (h *Handler) DeleteClass(c echo.Context) error {
	id := c.Param("id")
	// Simulate deleting a class
	// You would delete class info from the database here
	return c.JSON(http.StatusOK, map[string]string{"message": "Class deleted", "class_id": id})
}
