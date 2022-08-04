package service

import (
	"fmt"
	"github.com/chorshik/testToimi"
	"github.com/jmoiron/sqlx"
	"log"
)

type AdPostgres struct {
	db *sqlx.DB
}

func NewAdPostgres(db *sqlx.DB) *AdPostgres {
	return &AdPostgres{db: db}
}

func (r *AdPostgres) Create(list testToimi.Ad) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var id int
	createAdQuery := fmt.Sprintf("INSERT INTO %s (title, price, description, photo) VALUES ($1, $2, $3, $4) RETURNING id", adsTable)
	row := tx.QueryRow(createAdQuery, list.Title, list.Price, list.Description, list.Photo)
	fmt.Println(row)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		println("err")
		return 0, err
	}

	return id, tx.Commit()
}

func (r *AdPostgres) GetAll() ([]testToimi.Ad, error){
	var list []testToimi.Ad
	quwery := fmt.Sprintf("SELECT * FROM %s ", adsTable)
	err := r.db.Select(&list, quwery)
	if err != nil {
		log.Printf("%s",err)
	}

	return list, err
}

func (r *AdPostgres) GetById(adId int) (testToimi.Ad, error){
	var ad testToimi.Ad
	quwery := fmt.Sprintf("SELECT * FROM %s WHERE id=$1 ", adsTable)
	err := r.db.Get(&ad, quwery, adId)
	if err != nil {
		log.Printf("%s",err)
	}

	return ad, nil
}