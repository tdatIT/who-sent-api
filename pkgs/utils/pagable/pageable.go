package pagable

import (
	"math"
	"strconv"
)

const (
	defaultSize = 10
	maxSize     = 100
	defaultPage = 1
)

type PageableQuery struct {
	Page string `json:"page"`
	Size string `json:"size"`
}

type Query struct {
	Page              int      `json:"page"`
	Size              int      `json:"size"`
	ExpressionFilters []Filter `json:"filters"`
}

type ListResponse struct {
	Items   interface{} `json:"items"`
	Total   int         `json:"total"`
	Page    int         `json:"page"`
	Size    int         `json:"size"`
	HasMore bool        `json:"has_more"`
}

// SetSize Set page size
func (q *Query) SetSize(sizeQuery string) error {
	if sizeQuery == "" {
		q.Size = defaultSize
		return nil
	}

	n, err := strconv.ParseUint(sizeQuery, 10, 32)
	if err != nil {
		return err
	}

	q.Size = int(n)
	if q.Size > maxSize {
		q.Size = maxSize
	}

	return nil
}

// SetPage Set page number
func (q *Query) SetPage(pageQuery string) error {
	if pageQuery == "" {
		q.Page = defaultPage
		return nil
	}
	n, err := strconv.ParseUint(pageQuery, 10, 32)
	if err != nil {
		return err
	}
	q.Page = int(n)

	return nil
}

// GetOffset Get offset
func (q *Query) GetOffset() int {
	if q.Page == 0 {
		return 0
	}
	return (q.Page - 1) * q.Size
}

func (q *Query) GetLimit() int {
	return q.Size
}

func (q *Query) GetPage() int {
	return q.Page
}

func (q *Query) GetSize() int {
	if q.Size == 0 {
		return defaultSize
	}
	return q.Size
}

func (q *Query) GetTotalPages(totalCount int) int {
	d := float64(totalCount) / float64(q.GetSize())
	return int(math.Ceil(d))
}

func (q *Query) GetHasMore(total int) bool {
	return q.Page < total/q.GetSize()
}
