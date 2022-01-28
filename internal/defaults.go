package internal

import (
	"context"
)

func GetDefault(action string) string {
	var defaultValue string
	err := conn.QueryRowContext(context.Background(), "select `default` from `defaults` where `action`=?", action).Scan(&defaultValue)
	if err != nil {
		return ""
	}
	return defaultValue
}
