package auth

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func IsLogin(loginFlg LoginFlg, c *gin.Context) bool {

	var isLogin bool
	session := sessions.Default(c)

	var key string
	if loginFlg == AdminLoginFlg {
		key = "AdminId"
	} else if loginFlg == UserLoginFlg {
		key = "UserId"
	}

	value := session.Get(key)

	if value != nil {
		isLogin = true
	}

	return isLogin
}

func SetLogin(value string, loginFlg LoginFlg, c *gin.Context) error {

	session := sessions.Default(c)

	var key string
	if loginFlg == AdminLoginFlg {
		key = "AdminId"
	} else if loginFlg == UserLoginFlg {
		key = "UserId"
	}

	session.Set(key, value)

	err := session.Save()
	if err != nil {
		return err
	}

	return nil
}

func TruncateLogin(loginFlg LoginFlg, c *gin.Context) error {

	session := sessions.Default(c)

	var key string
	if loginFlg == AdminLoginFlg {
		key = "AdminId"
	} else if loginFlg == UserLoginFlg {
		key = "UserId"
	}

	session.Delete(key)

	err := session.Save()
	if err != nil {
		return err
	}

	return nil
}
