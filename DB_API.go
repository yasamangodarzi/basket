package main

import (
	_ "net/http"

	"github.com/yourusername/project/bussines"
)

func main() {
	e := echo.New()

	e.GET("/basket/", bussines.GetBaskets)
	e.POST("/basket/", bussines.CreateBasket)
	e.PATCH("/basket/:id", bussines.UpdateBasket)
	e.GET("/basket/:id", bussines.GetBasketByID)
	e.DELETE("/basket/:id", bussines.DeleteBasket)

	e.Logger.Fatal(e.Start(":8080"))
}
