package handler

import (
	psql "ecomm/pkg/db"
)

type Handler struct {
	DB *psql.DB
}

func CreateHandler(db *psql.DB) *Handler {
	return &Handler{DB: db}
}
