package user

import (
	"errors"

	"github.com/go-playground/validator/v10"

	"github.com/program-world-labs/DDDGo/internal/application/utils"
	"github.com/program-world-labs/DDDGo/internal/domain"
	"github.com/program-world-labs/DDDGo/internal/domain/entity"
)

type CreatedInput struct {
	Username    string   `json:"username" validate:"required,lte=30,alphanum"`
	Password    string   `json:"password" validate:"required"`
	EMail       string   `json:"email" validate:"required,email"`
	DisplayName string   `json:"displayName" validate:"required,lte=30"`
	Avatar      string   `json:"avatar" validate:"required"`
	RoleIDs     []string `json:"roleIds"`
	GroupID     string   `json:"groupId"`
}

func (i *CreatedInput) ToEntity() *entity.User {
	roles := make([]*entity.Role, 0)
	for _, v := range i.RoleIDs {
		roles = append(roles, &entity.Role{ID: v})
	}

	return &entity.User{
		Username:    i.Username,
		Password:    i.Password,
		EMail:       i.EMail,
		DisplayName: i.DisplayName,
		Avatar:      i.Avatar,
		Roles:       roles,
	}
}

func (i *CreatedInput) Validate() error {
	err := validator.New().Struct(i)
	// 提取錯誤訊息
	if err != nil {
		var e validator.ValidationErrors
		if errors.As(err, &e) {
			return utils.HandleValidationError("User", e)
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

func (i *ListGotInput) Validate() error {
	err := validator.New().Struct(i)
	// 提取錯誤訊息
	if err != nil {
		var e validator.ValidationErrors
		if errors.As(err, &e) {
			return utils.HandleValidationError("User", e)
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
			return utils.HandleValidationError("User", e)
		}

		return err
	}

	return nil
}

type UpdatedInput struct {
	ID          string `json:"id" validate:"required,alphanum,len=20"`
	DisplayName string `json:"displayName" validate:"required,lte=30"`
	Avatar      string `json:"avatar" validate:"required"`
}

func (i *UpdatedInput) Validate() error {
	err := validator.New().Struct(i)
	// 提取錯誤訊息
	if err != nil {
		var e validator.ValidationErrors
		if errors.As(err, &e) {
			return utils.HandleValidationError("User", e)
		}

		return err
	}

	return nil
}

func (i *UpdatedInput) ToEntity() *entity.User {
	return &entity.User{
		ID:          i.ID,
		DisplayName: i.DisplayName,
		Avatar:      i.Avatar,
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
			return utils.HandleValidationError("User", e)
		}

		return err
	}

	return nil
}

type Output struct {
	ID          string `gorm:"primary_key;"`
	Username    string `json:"username"`
	EMail       string `json:"email"`
	DisplayName string `json:"display_name"`
	Avatar      string `json:"avatar"`
}

func NewOutput(e *entity.User) *Output {
	return &Output{
		ID:          e.ID,
		Username:    e.Username,
		EMail:       e.EMail,
		DisplayName: e.DisplayName,
		Avatar:      e.Avatar,
	}
}

type OutputList struct {
	Offset int64    `json:"offset"`
	Limit  int64    `json:"limit"`
	Total  int64    `json:"total"`
	Items  []Output `json:"items"`
}

func NewListOutput(e *domain.List) *OutputList {
	var outputList []Output
	for _, v := range e.Data {
		outputList = append(outputList, *NewOutput(v.(*entity.User)))
	}

	return &OutputList{
		Offset: e.Offset,
		Limit:  e.Limit,
		Total:  e.Total,
		Items:  outputList,
	}
}
