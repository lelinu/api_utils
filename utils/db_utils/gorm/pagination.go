package gorm

import (
	"github.com/jinzhu/gorm"
)

//pagination struct
type Pagination struct {
	DB      *gorm.DB
	OrderBy []string
	Page    int64
	PerPage int64
}

func NewPagination(db *gorm.DB, orderBy []string, page int64, perPage int64) *Pagination {
	return &Pagination{
		DB:      db,
		OrderBy: orderBy,
		Page:    page,
		PerPage: perPage,
	}
}

//PaginationResult struct
type PaginationResult struct {
	TotalRecords int64       `json:"total_records"`
	Records      interface{} `json:"records"`
	CurrentPage  int64       `json:"current_page"`
	TotalPages   int64       `json:"total_pages"`
}

//Paginate function
func (p *Pagination) Paginate(dataSource interface{}) *PaginationResult {
	db := p.DB

	if len(p.OrderBy) > 0 {
		for _, o := range p.OrderBy {
			db = db.Order(o)
		}
	}

	done := make(chan bool, 1)
	var output PaginationResult
	var count int64
	var offset int64

	go countRecords(db, dataSource, done, &count)

	if p.Page == 1 {
		offset = 0
	} else {
		tmpPage := p.Page
		tmpPerPage := p.PerPage
		offset = (tmpPage - 1) * tmpPerPage
	}

	db.Limit(p.PerPage).Offset(offset).Find(dataSource)
	<-done

	output.TotalRecords = count
	output.Records = dataSource
	output.CurrentPage = p.Page
	output.TotalPages = getTotalPages(p.PerPage, count)

	return &output
}

func countRecords(db *gorm.DB, countDataSource interface{}, done chan bool, count *int64) {
	db.Model(countDataSource).Count(count)
	done <- true
}

func getTotalPages(perPage int64, totalRecords int64) int64 {
	totalPages := float64(totalRecords) / float64(perPage)
	return int64(totalPages)
}
