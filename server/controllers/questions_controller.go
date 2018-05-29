package controllers

import (
	"github.com/labstack/echo"
	"net/http"
)

func QueIndex() echo.HandlerFunc {
	return func(c echo.Context) error {
		jsonMap := map[string]string{
			"question": "index",
		}
		return c.JSON(http.StatusOK, jsonMap)
	}
}

func QueShow() echo.HandlerFunc {
	return func(c echo.Context) error {
		jsonMap := map[string]string{
			"question": "show",
		}
		return c.JSON(http.StatusOK, jsonMap)
	}
}
