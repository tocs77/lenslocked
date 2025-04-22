package models

import (
	"embed"
)

//go:embed sql
var FSsql embed.FS

func GetQuery(name string) (string, error) {
	query, err := FSsql.ReadFile("sql/" + name + ".sql")
	if err != nil {
		return "", err
	}
	return string(query), nil
}
