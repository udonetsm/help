package models

import (
	"encoding/json"

	"github.com/dgrijalva/jwt-go"
	"github.com/udonetsm/help/helper"
)

type Auth struct {
	Uid      string `json:"uid,omitempty"`
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}

type User struct {
	Uid   string `json:"uid"`
	Email string `json:"email"`
	Name  string `json:"name"`
	Dob   string `json:"dob"`
}

type Claims struct {
	jwt.StandardClaims
	User User
}

type ResponseAuth struct {
	Message string `json:"message"`
}

func DecodeUser(encode []byte) (user User) {
	defer helper.PanicCapture("decodeuser")
	err := json.Unmarshal(encode, &user)
	helper.Errors(err, "jsonunmarshall(decodeuser)")
	return
}
func DecodeAuth(encode []byte) (auth Auth) {
	defer helper.PanicCapture("decodeauth")
	err := json.Unmarshal(encode, &auth)
	helper.Errors(err, "jsonunmarshal(decodeauth)")
	return
}

func Encode(data interface{}) (encoded []byte) {
	defer helper.PanicCapture("encode")
	encoded, err := json.Marshal(data)
	helper.Errors(err, "jsonmarshall(encode)")
	return
}
