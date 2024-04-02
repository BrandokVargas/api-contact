package repository

import (
	"database/sql"
	"fmt"

	"github.com/BrandokVargas/api-contact/model"
)

const (
	migrateUser = `CREATE TABLE IF NOT EXISTS users(
		id SERIAL NOT NULL,
		name VARCHAR(100) NOT NULL,
		gmail VARCHAR(255) NOT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT now(),
		updated_at TIMESTAMP,
		CONSTRAINT product_id_pk PRIMARY KEY (id)
	)`

	createUser = `INSERT INTO users(name,gmail,created_at)
	VALUES($1,$2,$3) RETURNING id`
)

type RepositoryUser struct {
	db *sql.DB
}

func NewRepositoryUser(db *sql.DB) *RepositoryUser {
	return &RepositoryUser{db}
}

func (ru *RepositoryUser) Migrate() error {
	stmt, err := ru.db.Prepare(migrateUser)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec()

	if err != nil {
		return err
	}

	fmt.Println("Migracion de user ejecutado correctamente")
	return nil
}

func (ru *RepositoryUser) Create(mu *model.User) error {

	stmt, err := ru.db.Prepare(createUser)
	if err != nil {
		return nil
	}

	defer stmt.Close()

	err = stmt.QueryRow(
		mu.Name,
		mu.Gmail,
		mu.CreatedAt,
	).Scan(&mu.ID)

	if err != nil {
		return err
	}

	return nil
}
