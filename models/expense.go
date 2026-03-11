package models

import (
	"ExpenseTracker/config"
	"time"
)

type Expense struct {
	Id        int       `json:"id"`
	UserId    int       `json:"user_id"`
	Title     string    `json:"title"`
	Amount    string    `json:"amount"`
	Category  string    `json:"category"`
	Note      string    `json:"note"`
	Date      time.Time `json:"date"`
	CreatedAt time.Time `json:"created_at"`
}

func CreateExpense(e Expense) error {
	_, err := config.DB.Exec(
		"INSERT INTO expense( user_id,title,amount,category,note,date) VALUES ($1, $2, $3, $4, $5, $6)", e.UserId, e.Title, e.Amount, e.Category, e.Note, e.Date)
	return err
}

func GetExpenseByUser(userId int) ([]Expense, error) {
	rows, err := config.DB.Query(
		"SELECT id,user_id,title,amount,category,note,date,created_at WHERE user_id= $1 ORDER BY date DESC", userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var expenses []Expense
	for rows.Next() {
		var e Expense
		rows.Scan(&e.Id, &e.UserId, &e.Title, &e.Amount, &e.Category, &e.Note, &e.Date, &e.CreatedAt)
		expenses = append(expenses, e)
	}
	return expenses, nil
}

func UpdateExpense(id, userId int, e Expense) error {
	_, err := config.DB.Exec(
		"UPDATE expenses SET title=$1, amount=$2, category=$3, note=$4, date=$5 WHERE id=$6 and user_id=$7", e.Title, e.Amount, e.Category, e.Note, e.Date, id, userId)
	return err
}

func DeleteExpense(id, userID int) error {
	_, err := config.DB.Exec(
		"DELETE FROM expenses WHERE id=$1 AND user_id=$2", id, userID,
	)
	return err
}
