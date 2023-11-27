package main

import (
	"net/http"

	"github.com/devfullcycle/go-intensivo-jul/internal/entity"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/order", OrderHandler)
	e.Logger.Fatal(e.Start(":8888"))
}

func OrderHandler(c echo.Context) error {
	order, _ := entity.NewOrder("1", 10, 1)
	err := order.CalculateFinalPrice()
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, order)
}
