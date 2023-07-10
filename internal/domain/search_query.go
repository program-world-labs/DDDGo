package domain

type Page struct {
	Limit  int `json:"limit" form:"limit" validate:"min=1,max=1000"`
	Offset int `json:"offset" form:"offset" validate:"min=0"`
}

// Filter -.
type Filter struct {
	FilterField string      `json:"filter_field" form:"filter_field" validate:"required,ne="`
	Operator    string      `json:"operate" form:"operator" validate:"required,oneof=== != > >= < <= in not-in array-contains array-contains-any"`
	Value       interface{} `json:"value" form:"value" validate:"required"`
	Dir         string      `json:"dir" form:"dir" validate:"omitempty,oneof=asc desc"`
}

// Sort -.
type Sort struct {
	SortField string `json:"sort_field" form:"sort_field" validate:"required,ne="`
	Dir       string `json:"dir" form:"dir" validate:"omitempty,oneof=asc desc"`
}

type SearchQuery struct {
	Page   Page     `json:"page" binding:"required" form:"page" validate:"required"`
	Filter []Filter `json:"filter" form:"filter"`
	Sort   Sort     `json:"sort" form:"sort"`
}

func (*SearchQuery) GetWhere() string {
	return ""
}

func (*SearchQuery) GetArgs() []interface{} {
	return nil
}

func (*SearchQuery) GetOrder() string {
	return ""
}
