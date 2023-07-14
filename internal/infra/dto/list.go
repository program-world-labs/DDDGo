package dto

import (
	"github.com/program-world-labs/DDDGo/internal/domain"
)

type List struct {
	Limit  int         `json:"limit"`
	Offset int         `json:"offset"`
	Total  int         `json:"total"`
	Data   interface{} `json:"data"`
}

func (l *List) BackToDomain(model IRepoEntity) (*domain.List, error) {
	// Cast Data to []domain.IEntity
	var result []domain.IEntity

	for _, item := range l.Data.([]interface{}) {
		c, ok := item.(map[string]interface{})
		if !ok {
			return nil, nil
		}

		d, err := model.ParseMap(c)

		if err != nil {
			return nil, err
		}

		e, err := d.BackToDomain()
		if err != nil {
			return nil, err
		}

		result = append(result, e)
	}

	return &domain.List{
		Limit:  int64(l.Limit),
		Offset: int64(l.Offset),
		Total:  int64(l.Total),
		Data:   result,
	}, nil
}
