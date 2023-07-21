package group

import (
	"time"

	"github.com/program-world-labs/DDDGo/internal/adapter/http/v1/user"
	application_group "github.com/program-world-labs/DDDGo/internal/application/group"
)

type Response struct {
	ID          string          `json:"id"`
	Name        string          `json:"name"`
	Description string          `json:"description"`
	OwnerID     string          `json:"ownerId"`
	Users       []user.Response `json:"users"`
	CreatedAt   time.Time       `json:"createdAt"`
	UpdatedAt   time.Time       `json:"updatedAt"`
	DeletedAt   time.Time       `json:"deletedAt"`
}

type ResponseList struct {
	Offset int64      `json:"offset"`
	Limit  int64      `json:"limit"`
	Total  int64      `json:"total"`
	Items  []Response `json:"items"`
}

func NewResponse(model *application_group.Output) Response {
	userList := make([]user.Response, len(model.Users))

	for i, v := range model.Users {
		value := v
		userList[i] = user.NewResponse(&value)
	}

	return Response{
		ID:          model.ID,
		Name:        model.Name,
		Description: model.Description,
		OwnerID:     model.OwnerID,
		Users:       userList,
		CreatedAt:   model.CreatedAt,
		UpdatedAt:   model.UpdatedAt,
		DeletedAt:   model.DeletedAt,
	}
}

func NewResponseList(modelList *application_group.OutputList) ResponseList {
	responseList := make([]Response, len(modelList.Items))

	for i, v := range modelList.Items {
		value := v
		responseList[i] = NewResponse(&value)
	}

	return ResponseList{
		Offset: modelList.Offset,
		Limit:  modelList.Limit,
		Total:  modelList.Total,
		Items:  responseList,
	}
}
