package model

import "time"

type Transaction struct {
	ID              int    ` gorm:"primarykey" ` //account_id
	AccountID       string ` gorm:"foreignkey" `
	BankID          string ` gorm:"foreignkey" ` //mempertegas kolom
	Amount          int    ` gorm:"column:amount" `
	TransactionDate *time.Time
}

func (c *Transaction) TableName() string {
	return ("transaction")
}
