package context

import (
	"api.com/go-echo-rest-api/src/infrastructure/database"
	"github.com/labstack/echo/v4"
)

type CustomContext struct {
	echo.Context
	DB *database.SqlHandler
}
