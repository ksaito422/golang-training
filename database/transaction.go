package main

import (
	"context"
	"database/sql"
	"fmt"
)

type txAdmin struct {
	*sql.DB
}

type Service struct {
	tx txAdmin
}

// シンプルなトランザクション実装
func UpdateUser(db *sql.DB, userName, userID string) {
	tx, _ := db.Begin()

	if _, err := tx.Exec("UPDATE users SET user_name = $1 WHERE user_id = $2", userName, userID); err != nil {
		tx.Rollback()
	}
	tx.Commit()
}
