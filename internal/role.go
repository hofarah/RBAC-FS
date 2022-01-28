package internal

import (
	"context"
	"errors"
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
func RemoveRole(roleName string) error {
	d, err := conn.ExecContext(context.Background(), "delete from roles where name=?", roleName)
	if err != nil {
		return err
	}
	if count, _ := d.RowsAffected(); count == 0 {
		return errors.New("not deleted")
	}
	return nil
}
