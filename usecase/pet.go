package usecase

import (
	"context"

	"github.com/nao50/ddd-api/domain/model"
	"github.com/nao50/ddd-api/domain/repository"
)

type PetUsecase interface {
	CreatePet(name, tag string) (*model.Pet, error)
	ListPets(context.Context) ([]*model.Pet, error)
	GetPet(ctx context.Context, id string) (*model.Pet, error)
	UpdatePet(ctx context.Context, id, name, tag string) (*model.Pet, error)
	DeletePet(ctx context.Context, id string) error
}

type petUsecase struct {
	petRepository repository.PetRepository
}

func NewPetUsecase(pr repository.PetRepository) PetUsecase {
	return &petUsecase{
		petRepository: pr,
	}
}

//
func (pu petUsecase) ListPets(ctx context.Context) (pets []*model.Pet, err error) {
	pets, err = pu.petRepository.ListPets(ctx)
	if err != nil {
		return nil, err
	}
	return pets, nil
}

//
func (pu petUsecase) CreatePet(name, tag string) (pet *model.Pet, err error) {
	p, err := model.NewPet(name, tag)
	if err != nil {
		return nil, err
	}

	pet, err = pu.petRepository.CreatePet(p)
	if err != nil {
		return nil, err
	}
	return pet, nil
}

//
func (pu petUsecase) GetPet(ctx context.Context, id string) (pet *model.Pet, err error) {
	pet, err = pu.petRepository.GetPet(ctx, id)
	if err != nil {
		return nil, err
	}
	return pet, nil
}

//
// func (pu petUsecase) UpdatePet(ctx context.Context, p *model.Pet) (pet *model.Pet, err error) {
func (pu petUsecase) UpdatePet(ctx context.Context, id, name, tag string) (pet *model.Pet, err error) {
	pet, err = pu.petRepository.UpdatePet(ctx, &model.Pet{ID: id, Name: name, Tag: tag})
	if err != nil {
		return nil, err
	}
	return pet, nil
}

//
func (pu petUsecase) DeletePet(ctx context.Context, id string) (err error) {
	err = pu.petRepository.DeletePet(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
