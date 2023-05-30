package api

import (
	"context"
	"encoding/json"
	db "project/database"
	"project/src/schemas"

	"github.com/gin-gonic/gin"
)

func get_users(ctx *gin.Context) {
	var users []schemas.Users

	rows, err := db.DB.Query(context.Background(), `SELECT * FROM users`)

	if err != nil {
		ctx.JSON(500, err.Error())
		return
	}

	defer rows.Close()

	for rows.Next() {
		var user schemas.Users

		err := rows.Scan(
			&user.ID,
			&user.First_name,
			&user.Last_name,
			&user.Email,
		)

		if err != nil {
			ctx.JSON(500, err.Error())
			return
		}

		users = append(users, user)
	}

	ctx.JSON(200, users)
}

func create_user(ctx *gin.Context) {
	var user schemas.Create

	data, _ := ctx.GetRawData()
	json.Unmarshal(data, &user)

	_, err := db.DB.Exec(
		context.Background(),
		`
		INSERT INTO users (first_name, last_name, email)
		VALUES ($1, $2, $3)
		`,
		user.First_name, user.Last_name, user.Email,
	)

	if err != nil {
		ctx.JSON(500, err.Error())
		return
	}

	ctx.JSON(201, gin.H{"message": "Has been successfully created"})
}
