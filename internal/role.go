package internal

import (
	"context"
)

type role struct {
	id   int
	name string
}

func GetRole(r string) *role {
	var id int
	err := conn.QueryRowContext(context.Background(), "select id from roles where name=?", r).Scan(&id)
	if err != nil {
		return nil
	}
	return &role{id: id, name: r}
}

func NewRole(role string) error {
	_, err := conn.ExecContext(context.Background(), "insert into roles (`name`) values (?)", role)
	if err != nil {
		return err
	}
	return nil
}
