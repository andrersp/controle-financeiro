package database

import (
	"math"
	"net/http"
	"strconv"

	"gorm.io/gorm"
)

// Pagination for return pagination results
type Pagination struct {
	Limit       int         `json:"limit,omitempty;query:limit"`
	Page        int         `json:"page,omitempty;query:page"`
	TotalResult int64       `json:"total_results"`
	TotalPages  int         `json:"total_pages,omitempty"`
	Data        interface{} `json:"data"`
	Success     bool        `json:"success"`
}

// SinglePage for return single result
type SinglePage struct {
	Data    interface{} `json:"data,omitempty"`
	Success bool        `json:"success"`
}

// Paginator  return all results paginate
func Paginator(value interface{}, pagination *Pagination, db *gorm.DB, r *http.Request) func(db *gorm.DB) *gorm.DB {

	q := r.URL.Query()

	page, _ := strconv.Atoi(q.Get("page"))

	if page == 0 {
		page = 1
	}

	limit, _ := strconv.Atoi(q.Get("limit"))

	switch {
	case limit > 100:
		limit = 100
	case limit <= 0:
		limit = 10
	}

	pagination.Data = make([]string, 0)

	pagination.Page = page
	pagination.Limit = limit

	var totalRows int64
	db.Model(value).Count(&totalRows)

	pagination.TotalResult = totalRows
	totalPages := int(math.Ceil(float64(totalRows) / float64(pagination.Limit)))

	pagination.TotalPages = totalPages
	pagination.Success = true

	offset := (pagination.Page - 1) * pagination.Limit

	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(offset).Limit(pagination.Limit)
	}
}

// SingleResult return single result
func SingleResult(value interface{}, result *SinglePage, db *gorm.DB) func(db *gorm.DB) *gorm.DB {

	result.Success = true

	return func(db *gorm.DB) *gorm.DB {
		return db
	}
}
