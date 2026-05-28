package storage

import (
	"context"
	"log"
	"study/models"

	"github.com/jackc/pgx/v5"
)

type Storage struct {
	posgres *pgx.Conn
}

func NewStorage(connStr string) *Storage {
	conn, err := pgx.Connect(context.Background(), connStr)

	if err != nil {
		log.Fatal(err)
	}

	return &Storage{posgres: conn}
}

func (s *Storage) CreateUser(user models.User) error {
	_, err := s.posgres.Exec(context.Background(),
		`INSERT INTO users (name,age) VALUES ($1,$2)`,
		user.Name,
		user.Age,
	)
	if err != nil {
		return err
	}

	return nil
}

func (s *Storage) GetUser(name string) (models.User, error) {

	var user models.User

	err := s.posgres.QueryRow(context.Background(),
		`SELECT name, age FROM users WHERE name=$1`,
		name,
	).Scan(&user.Name, &user.Age)

	if err != nil {
		return user, err
	}

	return user, nil
}

func (s *Storage) GetUsers() ([]models.User, error) {
	users := []models.User{}

	rows, err := s.posgres.Query(context.Background(),
		`SELECT name, age FROM users`,
	)

	if err != nil {
		return users, err
	}

	defer rows.Close()

	for rows.Next() {
		var user models.User
		rows.Scan(&user.Name, &user.Age)

		users = append(users, user)
	}

	return users, nil
}
