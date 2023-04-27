package utils

import (
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetIdFromParam(e echo.Context) (int, error) {
	id, err := strconv.Atoi(e.Param("id"))
	return id, err
}
