package meta

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var FilterKeys = struct {
	Query    string
	QueryInt string
	Param    string
	ParamInt string
}{
	Query:    "query",
	QueryInt: "queryInt",
	Param:    "param",
	ParamInt: "paramInt",
}

type Filter struct {
	Keys   map[string]any
	Values map[string]any
}

// NewFilter membuat instance baru Filter dengan key yang diinginkan
func NewFilter(keys map[string]any) *Filter {
	return &Filter{
		Keys:   keys,
		Values: make(map[string]any),
	}
}

// Parse membaca value dari query atau param berdasarkan konfigurasi FilterKeys
func (f *Filter) Parse(ctx *gin.Context) {
	if len(f.Keys) == 0 {
		return
	}

	for key, mode := range f.Keys {
		switch mode {
		case FilterKeys.Query:
			f.Values[key] = ctx.Query(key)

		case FilterKeys.QueryInt:
			if valStr := ctx.Query(key); valStr != "" {
				if val, err := strconv.Atoi(valStr); err == nil {
					f.Values[key] = val
				}
			}

		case FilterKeys.Param:
			f.Values[key] = ctx.Param(key)

		case FilterKeys.ParamInt:
			if valStr := ctx.Param(key); valStr != "" {
				if val, err := strconv.Atoi(valStr); err == nil {
					f.Values[key] = val
				}
			}
		default:
			fmt.Printf("unknown filter mode for key %s: %v\n", key, mode)
		}
	}
}

func (f *Filter) Filterize(query *gorm.DB) *gorm.DB {
	for key, value := range f.Values {
		if value == nil || value == "" {
			continue
		}
		query = query.Where(fmt.Sprintf("%s = ?", key), value)
	}
	return query
}
