package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"
)

type User struct {
	UserID    string
	UserName  string
	CreatedAt time.Time
}

func SelectAll(db *sql.DB) {
	ctx := context.TODO()
	rows, err := db.QueryContext(ctx, `SELECT user_id, user_name, created_at FROM users ORDER BY user_id;`)
	if err != nil {
		log.Fatalf("query all users: %v", err)
	}
	defer rows.Close()

	var users []*User
	for rows.Next() {
		var (
			userID, userName string
			createdAt        time.Time
		)
		if err := rows.Scan(&userID, &userName, &createdAt); err != nil {
			log.Fatalf("scan the user: %v", err)
		}
		users = append(users, &User{
			UserID:    userID,
			UserName:  userName,
			CreatedAt: createdAt,
		})
	}
	if err := rows.Close(); err != nil {
		log.Fatalf("rows close: %v", err)
	}

	if err := rows.Err(); err != nil {
		log.Fatalf("scan users: %v", err)
	}

	fmt.Println(users[0].UserID, users[0].UserName, users[0].CreatedAt)
	fmt.Println(users[1].UserID, users[1].UserName, users[1].CreatedAt)
}

func Select(db *sql.DB, userID string) {
	var (
		userName  string
		createdAt time.Time
	)

	ctx := context.TODO()
	row := db.QueryRowContext(ctx, `SELECT user_name, created_at FROM users WHERE user_id = $1;`, userID)
	err := row.Scan(&userName, &createdAt)
	if err != nil {
		log.Fatalf("query row(user_id=%s): %v", userID, err)
	}
	u := User{
		UserID:    userID,
		UserName:  userName,
		CreatedAt: createdAt,
	}

	fmt.Println(u)
}
