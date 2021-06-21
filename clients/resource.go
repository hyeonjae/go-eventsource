package clients

import (
	"context"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Resource struct {
	service *Service
}

func NewResource(service *Service) *Resource {
	return &Resource{
		service: service,
	}
}

func (r Resource) CreateClient() func(echo.Context) error {
	return func(c echo.Context) error {
		fmt.Println(c.Path())
		var client Client
		if err := c.Bind(&client); err != nil {
			return err
		}

		command := EnrollCommand{
			Name:  client.Name,
			Email: client.Email,
		}

		ctx := context.Background()
		id, err := r.service.Process(ctx, command)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		response := Client{
			ID:    id,
			Name:  client.Name,
			Email: client.Email,
		}
		return c.JSON(http.StatusOK, response)
	}
}

func (Resource) ListClients() func(echo.Context) error {
	return func(c echo.Context) error {
		fmt.Println(c.Path())
		return nil
	}
}

func (r Resource) GetClient() func(echo.Context) error {
	return func(c echo.Context) error {
		fmt.Println(c.Path())
		var client Client
		if err := c.Bind(&client); err != nil {
			return err
		}
		query := GetQuery{
			ID: client.ID,
		}

		ctx := context.Background()
		aggregate, err := r.service.Query(ctx, &query)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, aggregate)
	}
}

func (r Resource) UpdateClient() func(echo.Context) error {
	return func(c echo.Context) error {
		fmt.Println(c.Path())
		var client Client
		if err := c.Bind(&client); err != nil {
			return err
		}
		command := UpdateCommand{
			ID: client.ID,
			Name:  client.Name,
			Email: client.Email,
		}

		ctx := context.Background()
		id, err := r.service.Process(ctx, command)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		response := Client{
			ID:    id,
			Name:  client.Name,
			Email: client.Email,
		}
		return c.JSON(http.StatusOK, response)
	}
}
