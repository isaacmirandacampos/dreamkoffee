package entity

import (
	"database/sql"
	"errors"
	"time"

	"github.com/shopspring/decimal"
)

type Expense struct {
	ID          int32           `db:"id" json:"id"`
	Value       decimal.Decimal `db:"value" json:"value"`
	Description string          `db:"description" json:"description"`
	CreatedAt   time.Time       `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time       `db:"updated_at" json:"updated_at"`
	DeletedAt   sql.NullTime    `db:"deleted_at" json:"deleted_at"`
}

func New(expense *Expense) Expense {
	return Expense{
		ID:          expense.ID,
		Value:       expense.Value,
		Description: expense.Description,
		CreatedAt:   expense.CreatedAt,
		UpdatedAt:   expense.UpdatedAt,
		DeletedAt:   expense.DeletedAt,
	}
}

func (e *Expense) ValueIsValid() error {
	if e.Value.IsNegative() {
		return errors.New("value cannot be negative")
	}
	if e.Value.IsZero() {
		return errors.New("value cannot be zero")
	}
	return nil
}
