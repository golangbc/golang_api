package users

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

func Hello(db *sql.DB) gin.HandlerFunc {

	return func(c *gin.Context) {
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
		err := c.BindJSON(&reqData)

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

}

// ユーザー一覧を取得
func FetchUserList(db *sql.DB) gin.HandlerFunc {

	return func(c *gin.Context) {
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
		err := c.BindJSON(&reqData)

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

}
