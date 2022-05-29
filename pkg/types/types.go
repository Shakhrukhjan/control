package types

type Money int

type Currency string

type Card struct {
	Balance  Money
	Currency Currency 
	Activity bool
}
