package main

import (
	"math/rand"
	"time"
)

type Account struct {
	ID         int       `json:"id"`
	FirstName  string    `json:"firstName"`
	LastName   string    `json:"lastName"`
	Number     int       `json:"number"`
	Balance    int64     `json:"balance"`
	CreatedAt time.Time `json:"createdAt`
}

func NewAccount(firstN string, lastN string) *Account {
	return &Account{
		ID:         rand.Intn(100000),
		FirstName:  firstN,
		LastName:   lastN,
		Number:     rand.Intn(100000),
		Balance:    0,
		CreatedAt: time.Now().UTC(), //best practice to use utc
	}
}

type CreateAccountRequest struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}
