package entities

import (
	"regexp"
)

type User struct {
	ID string `json:"-"`
	UserName string `json:"user_name,omitempty"`
	Password string `json:"password,omitempty"`
	Role string `json:"role,omitempty"`
}

func (u *User) IsValid() bool{
	return u.UserName != "" && len(u.UserName) > 5 && IsValidPassword(u.Password)
}

func IsValidPassword(pwd string) bool{
	if(pwd == "" || len(pwd) < 6 || len(pwd) > 15 ){ return false}
	re := regexp.MustCompile(`^(?!.*123456)(?=.*[A-Z])(?=.*\d).+$`)
	return re.MatchString(pwd)
}
