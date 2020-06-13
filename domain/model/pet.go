package model

import (
	"errors"

	"github.com/google/uuid"
)

type Pet struct {
	ID   string
	Name string
	Tag  string
}

//
func NewPet(name, tag string) (*Pet, error) {
	if name == "" {
		return nil, errors.New("name required")
	}

	uuidObj, _ := uuid.NewRandom()

	pet := &Pet{
		ID:   uuidObj.String(),
		Name: name,
		Tag:  tag,
	}

	return pet, nil
}
