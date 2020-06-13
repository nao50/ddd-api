package repository

import (
	"context"

	"github.com/nao50/ddd-api/domain/model"
)

// PetRepository の interface
type PetRepository interface {
	CreatePet(pet *model.Pet) (*model.Pet, error)
	ListPets(ctx context.Context) ([]*model.Pet, error)
	GetPet(ctx context.Context, id string) (*model.Pet, error)
	UpdatePet(ctx context.Context, pet *model.Pet) (*model.Pet, error)
	DeletePet(ctx context.Context, id string) error
}
