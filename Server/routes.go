package main

import (
	"github.com/labstack/echo"
	"net/http"
)

// Root is the root route to test service health
func Root(c echo.Context) error {
	println(" ")
	println("Base URL executed")
	return c.NoContent(http.StatusOK)
}