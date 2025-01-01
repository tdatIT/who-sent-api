package pagable

import (
	"github.com/labstack/echo/v4"
)

// GetQueryFromEchoCtx Get pagination query struct from
func GetQueryFromEchoCtx(e echo.Context) (*Query, error) {
	q := &Query{}

	if err := q.SetPage(e.QueryParam("page")); err != nil {
		return nil, err
	}

	if err := q.SetSize(e.QueryParam("size")); err != nil {
		return nil, err
	}

	queryString := e.Request().URL.String()

	filters, err := FilterBinding(queryString)
	if err != nil {
		return q, err
	}

	q.ExpressionFilters = filters

	return q, nil
}
