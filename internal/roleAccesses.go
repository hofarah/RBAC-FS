package internal

import (
	"context"
	"errors"
)

type roleAccess struct {
	roleID   int
	filePath string
}

func GetRoleAccess(filePath string) *roleAccess {
	var roleID int
	err := conn.QueryRowContext(context.Background(), "select roleID from roleAccesses where filePath=?", filePath).Scan(&roleID)
	if err != nil {
		return nil
	}
	return &roleAccess{roleID: roleID, filePath: filePath}
}

func NewRoleAccess(roleID int, filePath string) error {
	d, err := conn.ExecContext(context.Background(), "insert into roleAccesses (roleID,filePath) values (?,?)", roleID, filePath)
	if err != nil {
		return err
	}
	if count, _ := d.RowsAffected(); count == 0 {
		return errors.New("not deleted")
	}
	return nil
}

func RemoveRoleAccess(roleID int, filePath string) error {
	_, err := conn.ExecContext(context.Background(), "delete from roleAccesses where roleID=? AND filePath=?", roleID, filePath)
	if err != nil {
		return err
	}
	return nil
}
func Access(userID int, path string) bool {
	var count int
	err := conn.QueryRowContext(context.Background(), "select count(*) from roleAccesses inner join userRoles on (roleAccesses.roleID=userRoles.roleID) where userID=? AND filePath=?",
		userID, path).Scan(&count)
	if err != nil {
		return false
	}
	return count >= 1
}
