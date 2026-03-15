package handlers

import (
	"database/sql"
)

type ExpenseHandler struct {
	DB *sql.DB
}

func NewExpenseHandler(db *sql) *ExpenseHandler {
	return &ExpenseHandler{DB: db}
}
