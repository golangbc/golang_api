package auth

import (
	"database/sql"
	"net/http"

	"github.com/KatsutoshiOtogawa/batch/model/users"
	"github.com/gin-gonic/gin"
)

func CheckLogin(loginFlg LoginFlg, db *sql.DB) gin.HandlerFunc {

	return func(c *gin.Context) {
		if !IsLogin(loginFlg, c) {
			type RespData struct {
				Success bool   `json:"success"`
				Message string `json:"message"`
			}
			// ログインしていないのでエラーを返す。これを受け取ってログインを促す。
			c.AbortWithStatusJSON(
				http.StatusUnauthorized,
				RespData{
					Success: false,
					Message: "unauthorized",
				},
			)
			return
		}

		// ログインできていたら次に行く
		c.Next()
	}
}

func LoginAuth(loginFlg LoginFlg, db *sql.DB) gin.HandlerFunc {

	return func(c *gin.Context) {

		type ReqData struct {
			UserName string `json:"username" binding:"required"`
			Password string `json:"password" binding:"required"`
		}
		type RespData struct {
			Success bool   `json:"success"`
			Message string `json:"message"`
		}
		var reqData ReqData
		err := c.BindJSON(&reqData)

		if err != nil {
			// Authentication failed
			c.JSON(http.StatusBadRequest, RespData{
				Success: false,
				Message: "Parameters can't be empty",
			})
			return
		}

		// 値を送りつけていないとかだったらエラー。
		// Validate form input
		// if strings.Trim(username, " ") == "" || strings.Trim(password, " ") == "" {
		// 	c.JSON(http.StatusBadRequest, gin.H{"error": "Parameters can't be empty"})
		// 	return
		// }

		// dbにそのようなデータがあるか問い合わせ。
		isExist, err := users.PermittedLoginUser(reqData.UserName, reqData.Password, db)

		if err != nil || !isExist {
			// Authentication failed
			c.JSON(http.StatusUnauthorized, RespData{
				Success: false,
				Message: "Authentication failed",
			})
			return
		}

		// Save the username in the session
		err = SetLogin(reqData.UserName, loginFlg, c)
		if err != nil {
			c.JSON(http.StatusInternalServerError, RespData{
				Success: false,
				Message: "Failed to save session",
			})
			return
		}

		//
		c.JSON(http.StatusOK,
			RespData{
				Success: true,
				Message: "Successfully authenticated user",
			},
		)

	}
}

func Logout(loginFlg LoginFlg) gin.HandlerFunc {

	return func(c *gin.Context) {
		type RespData struct {
			Success bool   `json:"success"`
			Message string `json:"message"`
		}

		// ログインできてなかったらエラーを返す。
		if !IsLogin(loginFlg, c) {
			c.JSON(http.StatusBadRequest,
				RespData{
					Success: false,
					Message: "Invalid session token",
				})
			return
		}

		// 削除時にうまく消せなかったらエラーにする。
		err := TruncateLogin(loginFlg, c)
		if err != nil {
			c.JSON(http.StatusInternalServerError, RespData{
				Success: false,
				Message: "Failed to save session",
			})
			return
		}

		c.JSON(http.StatusOK, RespData{
			Success: true,
			Message: "Successfully logged out",
		})
	}

}
