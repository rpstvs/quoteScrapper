// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package database

type Quote struct {
	ID     int32
	Quote  string
	Author string
	Book   string
}
