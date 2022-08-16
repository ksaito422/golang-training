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

// トランザクションを制御するメソッド
// アプリ開発者が本メソッドを使って、DMLのクエリを発行する
func (t *txAdmin) Transaction(ctx context.Context, f func(ctx context.Context) (err error)) error {
	tx, err := t.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()
	if err := f(ctx); err != nil {
		return fmt.Errorf("transaction query failed: %w", err)
	}
	return tx.Commit()
}

func (s *Service) UpdateUser(ctx context.Context, userID, userName string) error {
	updateFunc := func(ctx context.Context) error {
		if _, err := s.tx.ExecContext(ctx, "UPDATE users SET user_name = $1 WHERE user_id = $2", userName, userID); err != nil {
			return err
		}
		return nil
	}
	return s.tx.Transaction(ctx, updateFunc)
}
