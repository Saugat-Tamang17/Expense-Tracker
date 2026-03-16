# Models Layer (`models` Directory)

The `models` directory contains all **data structures and database access logic** for the Expense Tracker application.

This layer is responsible for:
- Defining how data is represented in the application
- Executing SQL queries on the database
- Performing CRUD (Create, Read, Update, Delete) operations
- Keeping database logic separate from HTTP handlers

---

## Components Used

### Packages

| Package | Purpose |
|------|------|
| `config` | Provides the shared database connection (`config.DB`) |
| `time` | Used for date and timestamp fields |
| `database/sql` | Used internally by `config.DB` for executing SQL queries |

---

## Files Overview
