package meta

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Pagination struct {
	PerPage   int `json:"perPage"`
	Total     int `json:"total"`
	Page      int `json:"page"`
	TotalPage int `json:"totalPage"`
	NextPage  int `json:"nextPage"`
	PrevPage  int `json:"prevPage"`
}

// NewPagination membaca query param ?page= dan ?perPage= dari request
func NewPagination(ctx *gin.Context, pagination ...Pagination) *Pagination {
	var p Pagination
	if len(pagination) > 0 {
		p = pagination[0]
	}

	// Ambil dari query parameter
	if val := ctx.Query("page"); val != "" {
		if page, ok := parseInt(val); ok {
			p.Page = page
		}
	}
	if val := ctx.Query("perPage"); val != "" {
		if perPage, ok := parseInt(val); ok {
			p.PerPage = perPage
		}
	}

	p.PerPage = zeroDefaultTo(p.PerPage, 10)
	p.Page = zeroDefaultTo(p.Page, 1)

	return &p
}

// Paginate menghitung total dan menambahkan limit + offset ke query GORM
func (p *Pagination) Paginate(query *gorm.DB) *gorm.DB {
	var total int64
	count := query.Count(&total)
	if count.Error != nil {
		return count
	}

	p.Total = int(total)
	p.TotalPage = (p.Total + p.PerPage - 1) / p.PerPage

	if p.Page <= 1 {
		p.PrevPage = 1
	} else {
		p.PrevPage = p.Page - 1
	}

	if p.Page >= p.TotalPage {
		p.NextPage = p.TotalPage
	} else {
		p.NextPage = p.Page + 1
	}

	offset := (p.Page - 1) * p.PerPage
	return query.Offset(offset).Limit(p.PerPage)
}

// Helper
func zeroDefaultTo(actual, def int) int {
	if actual <= 0 {
		return def
	}
	return actual
}

func parseInt(s string) (int, bool) {
	var i int
	_, err := fmt.Sscanf(s, "%d", &i)
	return i, err == nil
}
