package service

import (
	"github.com/chorshik/testToimi"
	"github.com/jmoiron/sqlx"
)

type Ad interface {
	Create(list testToimi.Ad) (int, error)
	GetAll()([]testToimi.Ad, error)
	GetById(adId int) (testToimi.Ad, error)
}

type Repository struct {
	Ad
}

func NewRepository(db *sqlx.DB) *Repository{
	return &Repository{
		Ad: NewAdPostgres(db),
	}
}
