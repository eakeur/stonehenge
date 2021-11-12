module stonehenge

go 1.17

// This package is required to create UUID identifiers
require github.com/google/uuid v1.3.0

// This is required to enable HTTP routing
require github.com/go-chi/chi/v5 v5.0.4

require (
	github.com/go-sql-driver/mysql v1.6.0
	github.com/golang-jwt/jwt v3.2.2+incompatible
)
