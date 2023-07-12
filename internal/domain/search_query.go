package domain

import (
	"fmt"
	"strings"
)

type Page struct {
	Limit  int `json:"limit" form:"limit" validate:"min=1,max=1000"`
	Offset int `json:"offset" form:"offset" validate:"min=0"`
}

// Filter -.
type Filter struct {
	FilterField string      `json:"filter_field" form:"filter_field" validate:"required,ne="`
	Operator    string      `json:"operate" form:"operator" validate:"required,oneof=== != > >= < <= in not-in like between"`
	Value       interface{} `json:"value" form:"value" validate:"required"`
}

// Sort -.
type Order struct {
	OrderField string `json:"order_field" form:"order_field" validate:"required,ne="`
	Dir        string `json:"dir" form:"dir" validate:"omitempty,oneof=asc desc"`
}

type SearchQuery struct {
	Page    Page     `json:"page" binding:"required" form:"page" validate:"required"`
	Filters []Filter `json:"filters" form:"filters"`
	Orders  []Order  `json:"orders" form:"orders"`
}

func (sq *SearchQuery) GetWhere() string {
	var conditions []string

	for _, filter := range sq.Filters {
		switch filter.Operator {
		case "==":
			conditions = append(conditions, fmt.Sprintf("%s = ?", filter.FilterField))
		case "!=":
			conditions = append(conditions, fmt.Sprintf("%s != ?", filter.FilterField))
		case ">":
			conditions = append(conditions, fmt.Sprintf("%s > ?", filter.FilterField))
		case ">=":
			conditions = append(conditions, fmt.Sprintf("%s >= ?", filter.FilterField))
		case "<":
			conditions = append(conditions, fmt.Sprintf("%s < ?", filter.FilterField))
		case "<=":
			conditions = append(conditions, fmt.Sprintf("%s <= ?", filter.FilterField))
		case "in":
			conditions = append(conditions, fmt.Sprintf("%s IN (?)", filter.FilterField))
		case "not-in":
			conditions = append(conditions, fmt.Sprintf("%s NOT IN (?)", filter.FilterField))
		case "like":
			conditions = append(conditions, fmt.Sprintf("%s LIKE ?", filter.FilterField))
		case "between":
			conditions = append(conditions, fmt.Sprintf("%s BETWEEN ? AND ?", filter.FilterField))
		}
	}

	return strings.Join(conditions, " AND ")
}

func (sq *SearchQuery) GetArgs() []interface{} {
	var args []interface{}
	for _, filter := range sq.Filters {
		args = append(args, filter.Value)
	}

	return args
}

func (sq *SearchQuery) GetOrder() string {
	var orders []string
	for _, sort := range sq.Orders {
		orders = append(orders, fmt.Sprintf("%s %s", sort.OrderField, sort.Dir))
	}

	return strings.Join(orders, ", ")
}

func (sq *SearchQuery) GetKey() string {
	return fmt.Sprintf("%v-%v-%v", sq.Page, sq.Filters, sq.Orders)
}
