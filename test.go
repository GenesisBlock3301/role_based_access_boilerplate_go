package main

import (
	"database/sql"
	"time"
)

func main() {
	type User struct {
		ID           uint
		Name         *string
		Email        string
		Age          uint8
		Birthday     *time.Time
		MemberNumber sql.NullString
		ActivatedAt  sql.NullTime
		CreatedAt    time.Time
		UpdatedAt    time.Time
	}
	//fmt.Println(User{}.Name)
}
