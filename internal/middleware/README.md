# middleware/auth.go

## Overview

This file contains the JWT authentication middleware for the Expense Tracker API. It intercepts all protected routes and verifies the client's token before allowing the request to reach the actual handler.

---

## How It Works

**1. Extract the token**

The middleware reads the `Authorization` header from the incoming HTTP request. If the header is missing, the request is rejected immediately with a `401 Unauthorized` response.

**2. Trim the Bearer prefix**

Tokens arrive in the format `Bearer <token>`. The `Bearer` prefix is a standard HTTP convention and not part of the actual JWT, so it is stripped before validation.

**3. Validate the token**

The raw token string is parsed and verified against the `JWT_SECRET` from the environment variables. If the token has been tampered with, has expired, or is invalid for any reason, the request is rejected.

**4. Extract the user ID**

Once the token is confirmed valid, the `user_id` claim is extracted from the token's payload (claims).

**5. Pass via context**

Since Go's HTTP handler functions have a fixed signature `(w http.ResponseWriter, r *http.Request)`, extra data cannot be passed as parameters. Instead, the `user_id` is stored inside the request context using `context.WithValue`, acting as a bag that travels with the request to the next handler.

**6. Call the next handler**

The request is forwarded to the intended handler with the updated context attached via `r.WithContext(ctx)`.

---

## Request Flow

```
Incoming Request
    -> AuthMiddleware
        -> Check Authorization header
        -> Strip "Bearer" prefix
        -> Validate JWT against JWT_SECRET
        -> Extract user_id from claims
        -> Store user_id in context
        -> Forward to ExpenseHandler
```

---

## Usage

In `router/router.go`, protected routes are wrapped with this middleware:

```go
mux.Handle("/expenses", middleware.AuthMiddleware(http.HandlerFunc(handlers.ExpenseHandler)))
```

In `handlers/expense.go`, the user_id is retrieved from context:

```go
userID := int(r.Context().Value("userID").(float64))
```

---

## Notes

- The JWT validation block is standard boilerplate from the `golang-jwt/jwt` library and remains consistent across all token verification calls.
- The `JWT_SECRET` is loaded from the `.env` file via `os.Getenv` and never hardcoded.
- Context should only be used for request-scoped data such as user identity. It is not a substitute for proper function parameters in business logic.
