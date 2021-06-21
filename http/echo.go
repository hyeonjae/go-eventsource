package http

import (
	"fmt"
	"net/http"

	"github.com/hyeonjae/go-eventsource/accounts"
	"github.com/hyeonjae/go-eventsource/clients"
	"github.com/labstack/echo/v4"
)

func Start(cr *clients.Resource, ar *accounts.Resource) {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		fmt.Println(c.Path())
		return c.String(http.StatusOK, "Hello, World!")
	})

	clientRouter := e.Group("/clients")
	clientRouter.POST("", cr.CreateClient())
	clientRouter.GET("", cr.ListClients())
	clientRouter.GET("/:id", cr.GetClient())
	clientRouter.PUT("/:id", cr.UpdateClient())

	accountRouter := e.Group("/accounts")
	accountRouter.POST("", ar.CreateAccount())
	accountRouter.GET("", ar.ListAccounts())
	accountRouter.GET("/:id", ar.GetAccount())

	e.Logger.Fatal(e.Start(":8080"))
}
