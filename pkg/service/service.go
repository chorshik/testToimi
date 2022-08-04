package service

import (
	"github.com/chorshik/testToimi"
	repository "github.com/chorshik/testToimi/pkg/repository"
)

type Ad interface {
	Create(ad testToimi.Ad) (int, error)
	GetAll()([]testToimi.Ad, error)
	GetById(adId int) (testToimi.Ad, error)
}

type Service struct {
	Ad
}

func NewService(repos *repository.Repository) *Service{
	return &Service{
		Ad: NewAdService(repos.Ad),
	}
}
