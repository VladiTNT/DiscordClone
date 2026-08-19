package models

import _ "embed"

//go:embed user.sql
var UserTable string

type User struct {
	Name     string
	Email    string
	PassHash string
	PhotoId  int64
}
