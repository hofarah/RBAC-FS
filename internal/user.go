package internal

import (
	"context"
	"crypto/md5"
	"database/sql"
	"errors"
	"fmt"
	"github.com/hofarah/RBAC-FS/tools/sqlite"
)

var conn *sql.Conn

func init() {
	var err error
	conn, err = sqlite.GetConnection().Conn(context.Background())
	if err != nil {
		panic(err)
	}

}

type user struct {
	id       int
	username string
}

func (u *user) isAdmin() bool {
	return u.id == 1
}
func FirstUser() bool {
	var count int
	err := conn.QueryRowContext(context.Background(), "select count(*) from users").Scan(&count)
	if err != nil {
		panic(err)
	}
	return count == 0
}

func GetUser(username, password string) *user {

	hash := fmt.Sprintf("%x", md5.Sum([]byte(password)))
	var id int
	err := conn.QueryRowContext(context.Background(), "select id from users where username=? AND password=?", username, hash).Scan(&id)
	if err != nil {
		return nil
	}
	return &user{username: username, id: id}
}

func (u *user) UpdateUserPass(newPass string) error {

	hash := fmt.Sprintf("%x", md5.Sum([]byte(newPass)))
	_, err := conn.ExecContext(context.Background(), "update users set password=? where id=?", hash, u.id)
	if err != nil {
		return err
	}
	return nil
}

func NewUser(username, password, role string) error {

	hash := fmt.Sprintf("%x", md5.Sum([]byte(password)))
	r, err := conn.ExecContext(context.Background(), "insert into users (username,password) values (?,?)", username, hash)
	if err != nil {
		return err
	}
	id, _ := r.LastInsertId()

	rl := GetRole(role)
	if rl == nil {
		TPrint(NewError("role not found"))
		return errors.New("role not found")
	}
	_, err = conn.ExecContext(context.Background(), "insert into userRoles (roleID,userID) values (?,?)", rl.id, int(id))
	if err != nil {
		return err
	}

	return nil
}
