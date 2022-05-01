package models

import (
	"encoding/json"
	"io/ioutil"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-yaml/yaml"
	"github.com/udonetsm/help/helper"
)

type Postgres_conf struct {
	SslMode    string `yaml:"sslmode"`
	Dbname     string `yaml:"dbname"`
	Dbpassword string `yaml:"key"`
	Dbuser     string `yaml:"user"`
	Dbhost     string `yaml:"host"`
	Dbport     string `yaml:"port"`
}

type Auth struct {
	Uid      string `json:"uid,omitempty"`
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}

type User struct {
	Uid   string `json:"uid,omitempty"`
	Email string `json:"email,omitempty"`
	Name  string `json:"name,omitempty"`
	Dob   string `json:"dob,omitempty"`
}

type Claims struct {
	jwt.StandardClaims
	User User
}

type ResponseAuth struct {
	Message string `json:"message,omitempty"`
	Error   string `json:"error,omitempty"`
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

func PgConfStirng(filename string) string {
	defer helper.PanicCapture("parseYaml")
	pgconf := Postgres_conf{}
	content, err := ioutil.ReadFile(helper.Home() + "/" + filename)
	helper.Errors(err, "ioutillreadfile(parseyaml)")
	err = yaml.Unmarshal(content, &pgconf)
	helper.Errors(err, "yamlunmarshal(parseyaml)")
	return pgconf.Dbname + pgconf.Dbpassword + pgconf.SslMode +
		pgconf.Dbhost + pgconf.Dbport + pgconf.Dbuser
}
