package entity

import (
	"time"
)

type Department struct {
	ID          string    `json:"id"`          // 部門ID
	Name        string    `json:"name"`        // 部門名稱
	Description string    `json:"description"` // 部門描述
	Leader      string    `json:"leader"`      // 部門領導人
	Employees   []User    `json:"employees"`   // 部門員工列表
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	DeletedAt   time.Time `json:"deletedAt"`
}
