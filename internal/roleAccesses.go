package internal

import (
	"context"
	"errors"
	"strings"
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

func NewRoleAccess(roleID, level int, filePath string) error {
	d, err := conn.ExecContext(context.Background(), "insert into roleAccesses (roleID,filePath,access) values (?,?,?)", roleID, filePath, level)
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
func access(userID, level int, path string) bool {
	var acl int
	err := conn.QueryRowContext(context.Background(), "select access from roleAccesses inner join userRoles on (roleAccesses.roleID=userRoles.roleID) where userID=? AND filePath=?",
		userID, path).Scan(&acl)
	if err != nil {
		return false
	}
	return acl == level
}
func Access(userID, level int, path string) bool {
	if UserAccess(userID, path) {
		return true
	}
	if access(userID, level, path) {
		return true
	} else { //recursive check
		currentPath := path
		for currentPath != "" {
			if access(userID, level, currentPath) {
				return true
			}
			paths := strings.Split(currentPath, "/")
			if len(paths) == 0 {
				break
			}
			currentPath = strings.Join(paths[:len(paths)-1], "/")
		}

	}
	return false
}
