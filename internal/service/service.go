package service

import (
	"fmt"
	"study/internal/storage"
	"study/models"
)

type Service struct {
	service *storage.Storage
	redis   *storage.Redis
}

func NewService(s *storage.Storage, r *storage.Redis) *Service {
	return &Service{service: s, redis: r}
}

func (s *Service) CreateUser(user models.User) error {
	s.service.CreateUser(user)
	return s.redis.CreateUser(user)
}

func (s *Service) GetUsers() ([]models.User, error) {
	return s.service.GetUsers()
}

func (s *Service) GetUser(name string) (models.User, error) {
	user, err := s.redis.GetUser(name)

	if err == nil {
		fmt.Println("user from redis")
		return user, nil
	}

	user, err = s.service.GetUser(name)
	if err != nil {
		return user, err
	}

	if err := s.redis.CreateUser(user); err != nil {
		return user, err
	}

	fmt.Println("user from postgres")

	return user, nil

}
