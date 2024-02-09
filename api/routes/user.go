package user

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/zachatrocity/htmx-hyperscript-starter/api/types"
)

func GetUser(c echo.Context) error {
	u := &types.User{
		Id:    c.Param("id"),
		Name:  "Zach",
		Email: "email@email.com",
	}
	// simulate network request
	time.Sleep(3 * time.Second)
	return c.JSON(http.StatusOK, u)
}
