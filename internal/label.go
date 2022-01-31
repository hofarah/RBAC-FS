package internal

import (
	"context"
	"errors"
)

type label struct {
	id     int
	name   string
	access string
	level  int
}

func GetLabel(r string) *label {
	var id, level int
	var filePath string
	err := conn.QueryRowContext(context.Background(), "select id,access,`level` from labels where name=?", r).Scan(&id, &filePath, &level)
	if err != nil {
		return nil
	}
	return &label{id: id, name: r, access: filePath, level: level}
}

func NewLabel(label, filePath string, level, userID int) error {
	_, err := conn.ExecContext(context.Background(), "insert into labels (`name`,access,`level`) values (?,?,?)", label, filePath, level)
	if err != nil {
		return err
	}

	return nil
}
func RemoveLabel(labelName string) error {
	l := GetLabel(labelName)

	d, err := conn.ExecContext(context.Background(), "delete from labelUsers where labelID=?", l.id)
	if err != nil {
		return err
	}
	if count, _ := d.RowsAffected(); count == 0 {
		return errors.New("not deleted")
	}

	d, err = conn.ExecContext(context.Background(), "delete from labels where name=?", labelName)
	if err != nil {
		return err
	}
	if count, _ := d.RowsAffected(); count == 0 {
		return errors.New("not deleted")
	}
	return nil
}
func NewUserLabel(userID, labelID int) error {
	_, err := conn.ExecContext(context.Background(), "insert into labelUsers (userID,labelID) values (?,?)", userID, labelID)
	if err != nil {
		return err
	}
	return nil
}
