package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Create Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Gateway routes
	e.Any("/student/*", func(c echo.Context) error {
		target := "http://api-student:8082" + c.Request().URL.Path
		return proxyRequest(c, target)
	})

	e.Any("/class/*", func(c echo.Context) error {
		target := "http://api-class:8081" + c.Request().URL.Path
		return proxyRequest(c, target)
	})

	// Start gateway on port 8080
	e.Logger.Fatal(e.Start(":8080"))
}

// proxyRequest forwards the request to the correct microservice
func proxyRequest(c echo.Context, target string) error {
	req, err := http.NewRequest(c.Request().Method, target, c.Request().Body)
	if err != nil {
		return err
	}

	// Copy the headers
	for k, v := range c.Request().Header {
		req.Header.Set(k, v[0])
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Copy response to the client
	return c.Stream(resp.StatusCode, resp.Header.Get("Content-Type"), resp.Body)
}
