package entities

import (
	"regexp"
)

type User struct {
	ID string `json:"-"`
	UserName string `json:"userName,omitempty"`
	Password string `json:"password,omitempty"`
	Role string `json:"role,omitempty"`
}

type LoginUser struct{
	UserName string `json:"userName"`
	Password string `json:"password"`
}

func (u *User) IsValid() bool{
	return u.UserName != "" && len(u.UserName) > 5 && IsValidPassword(u.Password)
}

func (lu *LoginUser) IsValidAttempt() bool{
	return lu.UserName != "" && lu.Password != ""
}

func IsValidPassword(pwd string) bool{
	if(pwd == "" || len(pwd) < 6 || len(pwd) > 15 ){ return false}
	hasUpper := regexp.MustCompile(`[A-Z]`).MatchString(pwd)
    hasDigit := regexp.MustCompile(`\d`).MatchString(pwd)
    noSequence := !regexp.MustCompile(`123456`).MatchString(pwd)

	return hasUpper && hasDigit && noSequence
}
