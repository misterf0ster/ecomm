package handler

import (
	psql "ecomm/pkg/db"
)

type UserHandler struct {
	DB *psql.DB
}

func CreateUserHandler(db *psql.DB) *UserHandler {
	return &UserHandler{DB: db}
}
