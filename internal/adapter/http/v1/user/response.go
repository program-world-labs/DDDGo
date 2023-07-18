package user

import (
	application_user "github.com/program-world-labs/DDDGo/internal/application/user"
)

type Response struct {
	ID        string `json:"id"`
	Username  string `json:"username"`
	EMail     string `json:"email"`
	Avatar    string `json:"avatar"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type ResponseList struct {
	Offset int64      `json:"offset"`
	Limit  int64      `json:"limit"`
	Total  int64      `json:"total"`
	Items  []Response `json:"items"`
}

func NewResponse(model *application_user.Output) Response {
	return Response{
		ID:       model.ID,
		Username: model.Username,
		EMail:    model.EMail,
		Avatar:   model.Avatar,
	}
}

func NewResponseList(modelList *application_user.OutputList) ResponseList {
	responseList := make([]Response, len(modelList.Items))
	for i := range modelList.Items {
		responseList[i] = NewResponse(&modelList.Items[i])
	}

	return ResponseList{
		Offset: modelList.Offset,
		Limit:  modelList.Limit,
		Total:  modelList.Total,
		Items:  responseList,
	}
}
