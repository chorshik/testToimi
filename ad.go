package testToimi

import "github.com/lib/pq"

type Ad struct {
	Id          int            `json:"-" db:"id"`
	Title       string         `json:"title" db:"title" binding:"required"`
	Price       float64        `json:"price" db:"price" binding:"required"`
	Description string         `json:"description" db:"description" binding:"required"`
	Photo       pq.StringArray `json:"photo" db:"photo" binding:"required"`
}
