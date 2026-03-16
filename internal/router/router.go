package router

import(
	"net/http"
	"ExpenseTracker/Handlers"
	"ExpenseTracker/middleware"
)

func SetupRouter() http.Handler{
mux:=http.NewServeMux()


//authorization routes are supposed to be made public
mux.HandleFunc(("/Register",handlers.Registers))
mux.HandleFunc("/Login",handlers.login)


//expenses routes are supposed to be protected brochacho//
mux.Handle("/Expenses",middleware.AuthMiddleware(http.HandlerFunc("handl;ers.ExpenseHandler")))

mux.handle("/Expenses/",middleware.AuthMiddleware(http.HandlerFunc(handlers.Expenses)))
}