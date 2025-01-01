package echoUtils

import (
	"github.com/bytedance/sonic"
	"github.com/labstack/echo/v4"
	"io"
	"net/http"
)

type SonicJSONSerializer struct{}

// Serialize converts an interface into JSON and writes it to the response.
// Uses indentation if specified.
func (d SonicJSONSerializer) Serialize(c echo.Context, i interface{}, indent string) error {
	if indent != "" {
		prettyJSON, err := sonic.MarshalIndent(i, "", indent)
		if err != nil {
			return err
		}
		c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		_, err = c.Response().Write(prettyJSON)
		return err
	}

	jsonBytes, err := sonic.Marshal(i)
	if err != nil {
		return err
	}
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	_, err = c.Response().Write(jsonBytes)
	return err
}

// Deserialize reads a JSON from a request body and converts it into an interface using Sonic.
func (d SonicJSONSerializer) Deserialize(c echo.Context, i interface{}) error {
	body, err := io.ReadAll(c.Request().Body) // Thay thế bằng io.ReadAll
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Failed to read request body").SetInternal(err)
	}
	// Sử dụng Sonic để giải mã JSON
	err = sonic.Unmarshal(body, i)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Failed to unmarshal request body").SetInternal(err)
	}
	return err
}
