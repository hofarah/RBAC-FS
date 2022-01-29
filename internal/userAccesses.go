package internal

import (
	"context"
	"errors"
)

type userAccess struct {
	userID   int
	filePath string
}

func UserAccess(userID int, filePath string) bool {
	var count int
	err := conn.QueryRowContext(context.Background(), "select count(*) from userAccesses where userID=? AND filePath=?", userID, filePath).Scan(&count)
	if err != nil {
		return false
	}

	return count >= 1
}

func NewUserAccess(userID int, filePath string) error {
	d, err := conn.ExecContext(context.Background(), "insert into userAccesses (userID,filePath) values (?,?)", userID, filePath)
	if err != nil {
		return err
	}
	if count, _ := d.RowsAffected(); count == 0 {
		return errors.New("not deleted")
	}
	return nil
}

func RemoveUserAccess(userID int, filePath string) error {
	_, err := conn.ExecContext(context.Background(), "delete from userAccesses where userID=? AND filePath=?", userID, filePath)
	if err != nil {
		return err
	}
	return nil
}
