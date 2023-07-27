package wallet

import (
	"errors"
	"time"

	"github.com/go-playground/validator/v10"

	"github.com/program-world-labs/DDDGo/internal/application/utils"
	"github.com/program-world-labs/DDDGo/internal/domain"
	"github.com/program-world-labs/DDDGo/internal/domain/entity"
)

type CreatedInput struct {
	Name        string       `json:"name"`
	Description string       `json:"description" validate:"omitempty"`
	Chain       entity.Chain `json:"chain"`
	UserID      string       `json:"userId" validate:"required,alphanum,len=20"`
}

func (c *CreatedInput) Validate() error {
	validate := validator.New()

	// 提取錯誤訊息
	err := validate.Struct(c)
	if err != nil {
		var e validator.ValidationErrors
		if errors.As(err, &e) {
			return utils.HandleValidationError("Wallet", e)
		}

		return err
	}

	return nil
}

func (c *CreatedInput) ToEntity() *entity.Wallet {
	return &entity.Wallet{
		Name:        c.Name,
		Description: c.Description,
		Chain:       c.Chain,
		UserID:      c.UserID,
	}
}

type UpdatedInput struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (c *UpdatedInput) Validate() error {
	validate := validator.New()

	// 提取錯誤訊息
	err := validate.Struct(c)
	if err != nil {
		var e validator.ValidationErrors
		if errors.As(err, &e) {
			return utils.HandleValidationError("Wallet", e)
		}

		return err
	}

	return nil
}

func (c *UpdatedInput) ToEntity() *entity.Role {
	return &entity.Role{
		ID:          c.ID,
		Name:        c.Name,
		Description: c.Description,
	}
}

type DeletedInput struct {
	ID string `json:"id" validate:"required,alphanum,len=20"`
}

func (i *DeletedInput) Validate() error {
	err := validator.New().Struct(i)
	// 提取錯誤訊息
	if err != nil {
		var e validator.ValidationErrors
		if errors.As(err, &e) {
			return utils.HandleValidationError("Wallet", e)
		}

		return err
	}

	return nil
}

type DetailGotInput struct {
	ID string `json:"id" validate:"required,alphanum,len=20"`
}

func (i *DetailGotInput) Validate() error {
	err := validator.New().Struct(i)
	// 提取錯誤訊息
	if err != nil {
		var e validator.ValidationErrors
		if errors.As(err, &e) {
			return utils.HandleValidationError("Wallet", e)
		}

		return err
	}

	return nil
}

type ListGotInput struct {
	Limit      int      `json:"limit"`
	Offset     int      `json:"offset"`
	FilterName string   `json:"filterName" validate:"omitempty,oneof=name"`
	SortFields []string `json:"sortFields" validate:"dive,oneof=id name updated_at created_at"`
	Dir        string   `json:"dir" validate:"oneof=asc desc"`
}

func (i *ListGotInput) Validate() error {
	err := validator.New().Struct(i)
	// 提取錯誤訊息
	if err != nil {
		var e validator.ValidationErrors
		if errors.As(err, &e) {
			return utils.HandleValidationError("Wallet", e)
		}

		return err
	}

	return nil
}

func (i *ListGotInput) ToSearchQuery() *domain.SearchQuery {
	sq := &domain.SearchQuery{
		Page: domain.Page{
			Limit:  i.Limit,
			Offset: i.Offset,
		},
	}
	if i.FilterName != "" {
		sq.Filters = append(sq.Filters, domain.Filter{
			FilterField: "name",
			Operator:    "like",
			Value:       i.FilterName,
		})
	}

	if len(i.SortFields) > 0 {
		for _, sortField := range i.SortFields {
			sq.Orders = append(sq.Orders, domain.Order{
				OrderField: sortField,
				Dir:        i.Dir,
			})
		}
	}

	return sq
}

func (i *ListGotInput) ToEntity() *entity.Role {
	return &entity.Role{}
}

type Output struct {
	ID          string       `json:"id"`
	Name        string       `json:"name"`
	Description string       `json:"description"`
	Chain       entity.Chain `json:"chain"`
	UserID      string       `json:"userId"`
	CreatedAt   time.Time    `json:"createdAt"`
	UpdatedAt   time.Time    `json:"updatedAt"`
	DeletedAt   time.Time    `json:"deletedAt"`
}

type OutputList struct {
	Offset int64    `json:"offset"`
	Limit  int64    `json:"limit"`
	Total  int64    `json:"total"`
	Items  []Output `json:"items"`
}

func NewOutput(e *entity.Wallet) *Output {
	return &Output{
		ID:          e.ID,
		Name:        e.Name,
		Description: e.Description,
		Chain:       e.Chain,
		UserID:      e.UserID,
		CreatedAt:   e.CreatedAt,
		UpdatedAt:   e.UpdatedAt,
		DeletedAt:   e.DeletedAt,
	}
}

func NewListOutput(e *domain.List) *OutputList {
	var outputList []Output
	for _, v := range e.Data {
		outputList = append(outputList, *NewOutput(v.(*entity.Wallet)))
	}

	return &OutputList{
		Offset: e.Offset,
		Limit:  e.Limit,
		Total:  e.Total,
		Items:  outputList,
	}
}
