-- name: InsertQuote :one
INSERT INTO Quotes (Quote, Author, Book)
VALUES ($1, $2, $3)
RETURNING *;