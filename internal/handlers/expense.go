package handlers

import (
	"database/sql"
	"net/http"
)

type ExpenseHandler struct {
	DB *sql.DB
}

func NewExpenseHandler(db *sql.DB) *ExpenseHandler {
	return &ExpenseHandler{DB: db}
}

func (h *ExpenseHandler) CreateExpense(w http.ResponseWriter, r *http.Request) {
	//to be filled , lets not do this for now
}

func (h *ExpenseHandler) GetExpense(w http.ResponseWriter, r *http.Request) {
	//to be filled , lets not do this for now
}

func (h *ExpenseHandler) GetExpenseById(w http.ResponseWriter, r *http.Request) {
	//to be filled , lets not do this for now
}

func (h *ExpenseHandler) UpdateExpense(w http.ResponseWriter, r *http.Request) {
	//to be filled , lets not do this for now
}

func (h *ExpenseHandler) DeleteExpense(w http.ResponseWriter, r *http.Request) {
	//to be filled , lets not do this for now
}
