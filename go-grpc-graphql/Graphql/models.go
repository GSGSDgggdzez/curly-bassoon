package main

type Account struct {
	ID     string  `json:"id"` // Fixed json tag
	Name   string  `json:"name"`
	Orders []Order `json:"orders"` // Changed to Order to match schema
}
