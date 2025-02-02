package models

import (
	"time"

	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	Name         string        `json:"name" gorm:"not null"`
	Balance      float64       `json:"balance" gorm:"default:0"`
	Transactions []Transaction `json:"transactions" gorm:"foreignKey:account"`
}

type Transaction struct {
	gorm.Model
	Amount   float64   `json:"amount" gorm:"not null"`
	Date     time.Time `json:"date" gorm:"not null"`
	Category Category  `json:"category" gorm:"foreignKey:Id;not null"`
	Account  Account   `json:"account" gorm:"foreignKey:Id;not null"`
}

type Category struct {
	gorm.Model
	Name string `json:"name" gorm:"not null"`
}

type Budget struct {
	gorm.Model
	Category Category        `json:"category" gorm:"not null"`
	Amount   float64         `json:"amount" gorm:"not null"`
	Account  Account         `json:"account" gorm:"not null"`
	Changes  []BudgetChanges `json:"changes"`
	Date     time.Time       `json:"date" gorm:"not null"`
}

type BudgetChanges struct {
	gorm.Model
	Budget Budget    `json:"budget" gorm:"not null"`
	Amount float64   `json:"amount" gorm:"not null"`
	Reason string    `json:"reason" gorm:"not null"`
	Date   time.Time `json:"date" gorm:"not null"`
}
