// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package database

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Movie struct {
	ID          int32
	Title       string
	ReleaseDate pgtype.Date
}