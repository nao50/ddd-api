package handler

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/nao50/ddd-api/usecase"
)

type PetHandler interface {
	ListPets(http.ResponseWriter, *http.Request)
	CreatePet(http.ResponseWriter, *http.Request)
	GetPet(http.ResponseWriter, *http.Request)
	UpdatePet(http.ResponseWriter, *http.Request)
	DeletePet(http.ResponseWriter, *http.Request)
}

type petHandler struct {
	petUsecase usecase.PetUsecase
}

func NewPetHandler(pu usecase.PetUsecase) PetHandler {
	return &petHandler{
		petUsecase: pu,
	}
}

type petResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Tag  string `json:"tag"`
}

type petsResponse struct {
	Pets []petResponse `json:"pets"`
}

////////////////////////////////////////////////////////////////////////////////////////////////
func (ph petHandler) ListPets(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// Call Usecase
	pets, err := ph.petUsecase.ListPets(ctx)
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		return
	}

	// no content
	if len(pets) == 0 {
		w.WriteHeader(http.StatusNoContent)
	}

	//
	res := new(petsResponse)
	for _, pet := range pets {
		var pr petResponse
		pr = petResponse(*pet)
		res.Pets = append(res.Pets, pr)
	}

	//
	w.Header().Set("Content-Type", "application/json")
	if err = json.NewEncoder(w).Encode(res); err != nil {
		http.Error(w, "Internal Server Error", 500)
		return
	}
}

////////////////////////////////////////////////////////////////////////////////////////////////
func (ph petHandler) CreatePet(w http.ResponseWriter, r *http.Request) {
	var pr petResponse

	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Bad Request", 400)
		return
	}

	err := json.NewDecoder(r.Body).Decode(&pr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	pet, err := ph.petUsecase.CreatePet(pr.Name, pr.Tag)
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err = json.NewEncoder(w).Encode(pet); err != nil {
		http.Error(w, "Internal Server Error", 500)
		return
	}
}

////////////////////////////////////////////////////////////////////////////////////////////////
func (ph petHandler) GetPet(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := chi.URLParam(r, "id")

	var pr petResponse

	pet, err := ph.petUsecase.GetPet(ctx, id)
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		return
	}
	pr = petResponse(*pet)

	w.Header().Set("Content-Type", "application/json")
	if err = json.NewEncoder(w).Encode(pr); err != nil {
		http.Error(w, "Internal Server Error", 500)
		return
	}
}

////////////////////////////////////////////////////////////////////////////////////////////////
func (ph petHandler) UpdatePet(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := chi.URLParam(r, "id")

	var pr petResponse

	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Bad Request", 400)
		return
	}

	err := json.NewDecoder(r.Body).Decode(&pr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	pet, err := ph.petUsecase.UpdatePet(ctx, id, pr.Name, pr.Tag)
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		return
	}

	pr = petResponse(*pet)
	w.Header().Set("Content-Type", "application/json")
	if err = json.NewEncoder(w).Encode(pr); err != nil {
		http.Error(w, "Internal Server Error", 500)
		return
	}
}

////////////////////////////////////////////////////////////////////////////////////////////////
func (ph petHandler) DeletePet(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := chi.URLParam(r, "id")

	err := ph.petUsecase.DeletePet(ctx, id)
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
