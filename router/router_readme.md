# Router Setup Documentation

## Function: `SetupRouter()`

`SetupRouter()` is responsible for creating and configuring the HTTP
routing layer for the ExpenseTracker application.\
It defines which URL paths exist in the application and which handler
functions will process requests for those paths.

The function returns an `http.Handler`, which can then be used by the
HTTP server to handle incoming requests.

------------------------------------------------------------------------

## Purpose

The main responsibilities of this function are:

-   Create an HTTP request multiplexer (`ServeMux`)
-   Register all application routes
-   Apply authentication middleware to protected routes
-   Return the configured router to the main server

This separates routing logic from the rest of the application and keeps
the server initialization clean and modular.

------------------------------------------------------------------------

## Code Overview

``` go
func SetupRouter() http.Handler {
    mux := http.NewServeMux()

    // Auth routes (public)
    mux.HandleFunc("/register", handlers.Register)
    mux.HandleFunc("/login", handlers.Login)

    // Expense routes (protected)
    mux.Handle("/expenses", middleware.AuthMiddleware(http.HandlerFunc(handlers.ExpenseHandler)))
    mux.Handle("/expenses/", middleware.AuthMiddleware(http.HandlerFunc(handlers.ExpenseHandler)))

    return mux
}
```

------------------------------------------------------------------------

## How It Works

### 1. Creating the Router

``` go
mux := http.NewServeMux()
```

`ServeMux` is the standard HTTP request multiplexer provided by Go's
`net/http` package.\
It matches incoming request paths to registered handlers.

------------------------------------------------------------------------

### 2. Public Authentication Routes

``` go
mux.HandleFunc("/register", handlers.Register)
mux.HandleFunc("/login", handlers.Login)
```

These endpoints allow users to:

-   **Register** a new account
-   **Log in** to obtain authentication credentials

These routes are **public**, meaning they do not require authentication.

------------------------------------------------------------------------

### 3. Protected Expense Routes

``` go
mux.Handle("/expenses", middleware.AuthMiddleware(http.HandlerFunc(handlers.ExpenseHandler)))
mux.Handle("/expenses/", middleware.AuthMiddleware(http.HandlerFunc(handlers.ExpenseHandler)))
```

These routes handle expense-related operations such as:

-   Creating expenses
-   Retrieving expenses
-   Updating expenses
-   Deleting expenses

They are wrapped with:

    middleware.AuthMiddleware

This middleware ensures that:

-   The request includes a valid authentication token
-   Only authenticated users can access expense data

If authentication fails, the middleware blocks the request before it
reaches the handler.

------------------------------------------------------------------------

### 4. Returning the Router

``` go
return mux
```

The configured router is returned as an `http.Handler`.\
This allows it to be passed directly to the HTTP server.

------------------------------------------------------------------------

## Example Implementation

In your `main.go`, the router can be used like this:

``` go
package main

import (
    "fmt"
    "net/http"
    "ExpenseTracker/router"
)

func main() {
    r := router.SetupRouter()

    fmt.Println("Server running on port 8080")
    http.ListenAndServe(":8080", r)
}
```

------------------------------------------------------------------------

## Request Flow

Example request flow:

    Client Request
          |
          v
    HTTP Server
          |
          v
    ServeMux (router)
          |
          +---- /register  --> handlers.Register
          |
          +---- /login     --> handlers.Login
          |
          +---- /expenses  --> AuthMiddleware --> handlers.ExpenseHandler

------------------------------------------------------------------------

## Key Design Advantages

1.  **Separation of concerns**\
    Routing logic is isolated from business logic.

2.  **Middleware integration**\
    Authentication can be added easily without modifying handlers.

3.  **Scalability**\
    New routes can be added to the router without changing the server
    setup.

4.  **Maintainability**\
    All endpoints are clearly defined in one place.

------------------------------------------------------------------------

## Summary

`SetupRouter()` acts as the central routing configuration for the
ExpenseTracker API.\
It registers all endpoints, attaches authentication middleware to
protected routes, and returns a fully configured HTTP handler that can
be used by the Go HTTP server.
