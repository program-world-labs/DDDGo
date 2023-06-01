package webapi

import (
	translator "github.com/Conight/go-googletrans"

	entity "github.com/program-world-labs/DDDGo/internal/domain/user/entity"
)

// UserWebAPI -.
type UserWebAPI struct {
	conf translator.Config
}

// New -.
func New() *UserWebAPI {
	conf := translator.Config{
		UserAgent:   []string{"Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:15.0) Gecko/20100101 Firefox/15.0.1"},
		ServiceUrls: []string{"translate.google.com"},
	}

	return &UserWebAPI{
		conf: conf,
	}
}

// Translate -.
func (t *UserWebAPI) Translate(User entity.User) (entity.User, error) {
	// trans := translator.New(t.conf)

	// result, err := trans.Translate(User.Original, User.Source, User.Destination)
	// if err != nil {
	// 	return entity.User{}, fmt.Errorf("UserWebAPI - Translate - trans.Translate: %w", err)
	// }

	// User.User = result.Text

	return User, nil
}
