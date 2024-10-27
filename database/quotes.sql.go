// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: quotes.sql

package database

import (
	"context"
)

const insertQuote = `-- name: InsertQuote :one
INSERT INTO Quotes (Quote, Author, Book)
VALUES ($1, $2, $3)
RETURNING id, quote, author, book
`

type InsertQuoteParams struct {
	Quote  string
	Author string
	Book   string
}

func (q *Queries) InsertQuote(ctx context.Context, arg InsertQuoteParams) (Quote, error) {
	row := q.db.QueryRowContext(ctx, insertQuote, arg.Quote, arg.Author, arg.Book)
	var i Quote
	err := row.Scan(
		&i.ID,
		&i.Quote,
		&i.Author,
		&i.Book,
	)
	return i, err
}
