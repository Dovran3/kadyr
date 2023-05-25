package src

import (
	//"fmt"
	//"database/sql"
    //"net/http"
	"encoding/json"
	//"log"
	db "admin/src/database"
	"context"
	//"github.com/jackc/pgx/v5"
    bod "admin/src/structs"
	"github.com/gin-gonic/gin"
)



func GetUsers(ctx *gin.Context) {
	var users []bod.Users

	rows, err := db.DB.Query(context.Background(), `select * from test1`)

	if err != nil {
		ctx.JSON(500, err.Error())
		return
	}

	defer rows.Close()

	for rows.Next() {
		var user bod.Users
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

	ctx.IndentedJSON(200, users)
}



func PostData(ctx *gin.Context){
 
	var user bod.Users
	data := ctx.MustGet("data")
	byte_data, _ := json.Marshal(data)
	json.Unmarshal(byte_data, &user)

	db.DB.Exec(context.Background(),
			`INSERT INRO test1 ("id", "first_name", "last_name", "email")
			VALUE ($1, $2, $3, $4)`,user.ID, user.First_name, user.Last_name, user.Email)

	ctx.JSON(201, "ALL ok")
    
}





