package internal

import (
	"context"
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
	_, err := conn.ExecContext(context.Background(), "insert into roleAccesses (roleID,filePath) values (?,?)", roleID, filePath)
	if err != nil {
		return err
	}
	return nil
}
