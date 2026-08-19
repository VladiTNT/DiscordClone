package models

import _ "embed"

//go:embed photo.sql
var PhotoTable string

type Photo struct {
	Id     int64
	Format string
	Data   []byte
}
