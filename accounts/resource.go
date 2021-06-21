package accounts

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

type Resource struct {
}

func NewResource() *Resource {
	return &Resource{}
}

func (Resource) CreateAccount() func(c echo.Context) error {
	return func(c echo.Context) error {
		fmt.Println(c.Path())
		return nil
	}
}

func (Resource) ListAccounts() func(c echo.Context) error {
	return func(c echo.Context) error {
		fmt.Println(c.Path())
		return nil
	}
}

func (Resource) GetAccount() func(c echo.Context) error {
	return func(c echo.Context) error {
		fmt.Println(c.Path())
		return nil
	}
}
