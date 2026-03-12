
---

## Expense Model (`expense.go`)

### Expense Struct

Represents an expense record stored in the database.

```go
type Expense struct {
 Id        int
 UserId    int
 Title     string
 Amount    string
 Category  string
 Note      string
 Date      time.Time
 CreatedAt time.Time
}