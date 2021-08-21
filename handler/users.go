package users

import (
	"database/sql"
	"errors"

	"github.com/gin-gonic/gin"
)

// 管理者権限、バッチのみ使える処理をここにまとめます。
func Hello(c *gin.Context, db *sql.DB) {

	type ReqData struct {
		Id     int `json:"id"`
		Offset int `json:"offset"`
	}

	type User struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
	}

	type RespData struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
		Users   []User `json:"users"`
	}

	var reqData ReqData
	c.BindJSON(&reqData)

	err := errors.New("error response")
	if err != nil {

		c.JSON(200, gin.H{
			"success": false,
			"message": err.Error(),
		})
	} else {
		c.JSON(200, RespData{
			Success: true,
			Message: "ok",
			Users: []User{
				{
					FirstName: "test",
					LastName:  "name",
				},
			},
		})
	}

}
