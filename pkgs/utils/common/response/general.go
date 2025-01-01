package responses

import "github.com/labstack/echo/v4"

type General struct {
	Status  int         `json:"-"`
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func (g *General) JSON(c echo.Context) error {
	return c.JSON(200, g)
}
