package entities

import (
	"regexp"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	UserName string `bson:"userName" json:"user_name"`
	Password string `bson:"password" json:"-"`
}

func (u *User) IsValid() bool{
	return u.UserName != "" && len(u.UserName) > 5 && IsValidPassword(u.Password)
}

func IsValidPassword(pwd string) bool{
	if(pwd == "" || len(pwd) < 6 || len(pwd) > 15 ){ return false}
	re := regexp.MustCompile(`^(?!.*123456)(?=.*[A-Z])(?=.*\d).+$`)
	return re.MatchString(pwd)
}
