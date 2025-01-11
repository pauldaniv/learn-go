package model

import "time"

// Currency type (Enum in PostgreSQL)
type Currency string

const (
	CurrencyUAH Currency = "UAH"
	CurrencyUSD Currency = "USD"
	CurrencyEUR Currency = "EUR"
)

// Transaction model
type Transaction struct {
	ID           string    `json:"id"`
	Amount       int       `json:"amount"`
	CurrencyType Currency  `json:"currency_type"`
	CreatedAt    time.Time `json:"created_at"`
}

// Broker model
type Broker struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

// BrokerAccount model
type BrokerAccount struct {
	ID        string    `json:"id"`
	BrokerID  string    `json:"broker_id"`
	CreatedAt time.Time `json:"created_at"`
}

// Bank model
type Bank struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

// BankAccount model
type BankAccount struct {
	ID           string    `json:"id"`
	Name         string    `json:"name"`
	Amount       int       `json:"amount"`
	CurrencyType Currency  `json:"currency_type"`
	BankID       string    `json:"bank_id"`
	CreatedAt    time.Time `json:"created_at"`
}

// Deposit model
type Deposit struct {
	ID           string    `json:"id"`
	Name         string    `json:"name"`
	Amount       int       `json:"amount"`
	CurrencyType Currency  `json:"currency_type"`
	StartDate    time.Time `json:"start_date"`
	EndDate      time.Time `json:"end_date"`
	CreatedAt    time.Time `json:"created_at"`
}

// Bond model
type Bond struct {
	ID           string    `json:"id"`
	Name         string    `json:"name"`
	Count        int       `json:"count"`
	BuyPrice     float64   `json:"buy_price"`
	SellPrice    float64   `json:"sell_price"`
	CurrencyType Currency  `json:"currency_type"`
	StartDate    time.Time `json:"start_date"`
	EndDate      time.Time `json:"end_date"`
	CreatedAt    time.Time `json:"created_at"`
}

// Coupon model
type Coupon struct {
	ID          string    `json:"id"`
	Amount      float64   `json:"amount"`
	PaymentDate time.Time `json:"payment_date"`
	BondID      string    `json:"bond_id"`
	CreatedAt   time.Time `json:"created_at"`
}
