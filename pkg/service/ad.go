package service

import (
	"github.com/chorshik/testToimi"
	repository "github.com/chorshik/testToimi/pkg/repository"
)

type AdService struct {
	repo repository.Ad
}

func NewAdService(repo repository.Ad) *AdService {
	return &AdService{repo: repo}
}

func (s *AdService) Create(list testToimi.Ad) (int, error) {
	return s.repo.Create(list)
}

func (s *AdService) GetAll() ([]testToimi.Ad, error) {
	return s.repo.GetAll()
}

func (s *AdService) GetById(adId int) (testToimi.Ad, error) {
	return s.repo.GetById(adId)
}