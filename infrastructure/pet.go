package infrastructure

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
	"github.com/nao50/ddd-api/domain/model"
	"github.com/nao50/ddd-api/domain/repository"
)

type PetRepository struct {
	DbConnection *sql.DB
}

func checkTable(conn *sql.DB) error {
	cmd := `CREATE TABLE IF NOT EXISTS pet(
		id STRING,
		name STRING,
		tag STRING
		)`
	_, err := conn.Exec(cmd)
	if err != nil {
		fmt.Println("checkTable err", err)
		return err
	}
	return nil
}

func NewPetRepository(conn *sql.DB) repository.PetRepository {
	return &PetRepository{DbConnection: conn}
}

func (pr *PetRepository) CreatePet(pet *model.Pet) (*model.Pet, error) {
	err := checkTable(pr.DbConnection)
	if err != nil {
		return nil, err
	}

	cmd := "INSERT INTO pet (id, name, tag) VALUES (?, ?, ?)"

	_, err = pr.DbConnection.Exec(cmd, &pet.ID, &pet.Name, &pet.Tag)
	if err != nil {
		return nil, err
	}

	return pet, nil
}

func (pr *PetRepository) ListPets(context.Context) ([]*model.Pet, error) {
	err := checkTable(pr.DbConnection)
	if err != nil {
		return nil, err
	}

	cmd := "SELECT * FROM pet"
	rows, err := pr.DbConnection.Query(cmd)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()

	var pets []*model.Pet
	for rows.Next() {
		var pet model.Pet
		err := rows.Scan(&pet.ID, &pet.Name, &pet.Tag)
		if err != nil {
			fmt.Println(err)
		}
		pets = append(pets, &pet)
	}
	return pets, nil
}

//
func (pr *PetRepository) GetPet(ctx context.Context, id string) (*model.Pet, error) {
	err := checkTable(pr.DbConnection)
	if err != nil {
		return nil, err
	}

	cmd := "SELECT * FROM pet WHERE id=?"
	row := pr.DbConnection.QueryRow(cmd, id)

	var pet model.Pet
	err = row.Scan(&pet.ID, &pet.Name, &pet.Tag)
	if err != nil {
		if err == sql.ErrNoRows {
			return &pet, nil
		} else {
			return nil, err
		}
	}
	return &pet, nil
}

//
func (pr *PetRepository) UpdatePet(ctx context.Context, pet *model.Pet) (*model.Pet, error) {
	err := checkTable(pr.DbConnection)
	if err != nil {
		return nil, err
	}

	cmd := "UPDATE pet SET name=?,tag=?  WHERE id=?"
	_, err = pr.DbConnection.Exec(cmd, &pet.Name, &pet.Tag, &pet.ID)
	if err != nil {
		return nil, err
	}

	return pet, nil
}

//
func (pr *PetRepository) DeletePet(ctx context.Context, id string) error {
	err := checkTable(pr.DbConnection)
	if err != nil {
		return err
	}

	cmd := "DELETE FROM pet WHERE id = ?"
	_, err = pr.DbConnection.Exec(cmd, id)
	if err != nil {
		return err
	}
	return nil
}
