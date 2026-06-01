package service

import (
	"fmt"
	"study/models"
)

type Postgres interface {
	CreateUser(user models.User) error
	GetUser(name string) (models.User, error)
	GetUsers() ([]models.User, error)
}

type Redis interface {
	CreateUser(user models.User) error
	GetUser(name string) (models.User, error)
}

type Service struct {
	store Postgres
	redis Redis
}

func NewService(s Postgres, r Redis) *Service {
	return &Service{store: s, redis: r}
}

func (s *Service) CreateUser(user models.User) error {
	s.store.CreateUser(user)
	return s.redis.CreateUser(user)
}

func (s *Service) GetUsers() ([]models.User, error) {
	return s.store.GetUsers()
}

func (s *Service) GetUser(name string) (models.User, error) {
	user, err := s.redis.GetUser(name)

	if err == nil {
		fmt.Println("user from redis")
		return user, nil
	}

	user, err = s.store.GetUser(name)
	if err != nil {
		return user, err
	}

	if err := s.redis.CreateUser(user); err != nil {
		return user, err
	}

	fmt.Println("user from postgres")

	return user, nil

}
