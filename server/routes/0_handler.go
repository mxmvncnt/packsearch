package routes

import "github.com/mxmvncnt/packsearch/database"

// Filename starts with 0 to be at the top of the list

type RoutesHandler struct {
	db *database.Queries
}

func NewRoutesHandler(pool *database.Queries) *RoutesHandler {
	return &RoutesHandler{db: pool}
}
