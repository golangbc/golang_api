package auth

type LoginFlg = uint

const (
	AdminLoginFlg LoginFlg = 1 << iota
	UserLoginFlg
)
