package model

import "github.com/src/github.com/jinzhu/gorm"

type Menu struct {
	gorm.Model
	Name   string
	Price  int
	Status int
}

type Order struct {
	gorm.Model
	MenuId       Menu `gorm:"foreignkey:UserRefer"`
	Qty          int
	CustomerName string
	TableNumber  int
	Status       int
}

type Payment struct {
	gorm.Model
	Type            string
	ReferenceNumber int
	OrderId         Order `gorm:"foreignkey:UserRefer"`
	Amount          int
	Status          int
}
