package router

import(
	"net/http"
	"ExpenseTracker/Handlers"
	"ExpenseTracker/middleware"
)

func SetupRouter() http.Handler{
mux:=http.NewServeMux()


//authorization routes are supposed to be made public
mux.HandleFunc(("/register",handlers.Registers))
mux.HandleFunc("/login",handlers.login)


//expenses routes are supposed to be protected brochacho//
mux.Handle("/Expenses",middleware.AuthMiddleware(http.HandlerFunc("handl;ers.ExpenseHandler")))

mux.handle("/Expenses/",middleware.AuthMiddleware(http.HandlerFunc(handlers.Expenses)))
}