package models

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-yaml/yaml"
	"github.com/udonetsm/help/helper"
	"golang.org/x/crypto/bcrypt"
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

type AUser struct {
	Auth Auth `json:"auth"`
	User User `json:"user"`
}

type Claims struct {
	jwt.StandardClaims
	User User
}

type ResponseAuth struct {
	Message string `json:"message,omitempty"`
	Error   string `json:"error,omitempty"`
}

type Srver_Conf struct {
	Addr   string `yaml:"addr"`
	Secret string `yaml:"secret"`
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

func (conf Postgres_conf) StoreConf(path string) Postgres_conf {
	defer helper.PanicCapture("parseYaml")
	content, err := ioutil.ReadFile(path)
	helper.Errors(err, "ioutillreadfile(parseyaml)")
	err = yaml.Unmarshal(content, &conf)
	helper.Errors(err, "yamlunmarshal(parseyaml)")
	return conf
}

func (conf Srver_Conf) ServerConf(path string) Srver_Conf {
	defer helper.PanicCapture("parseYaml")
	content, err := ioutil.ReadFile(path)
	helper.Errors(err, "ioutillreadfile(parseyaml)")
	err = yaml.Unmarshal(content, &conf)
	helper.Errors(err, "yamlunmarshal(parseyaml)")
	return conf
}

func (user *AUser) BuildUser(w http.ResponseWriter, r *http.Request) []byte {
	r.ParseForm()
	user.User.Name = r.FormValue("name")
	user.User.Dob = r.FormValue("date")
	user.User.Email = r.FormValue("email")
	user.Auth.Email = user.User.Email
	user.Auth.Password = BcryptHasher(r.FormValue("password"))
	return Encode(user)
}

func (auth *Auth) BuildAuth(w http.ResponseWriter, r *http.Request) []byte {
	r.ParseForm()
	auth.Email = r.FormValue("email")
	auth.Password = r.FormValue("password")
	return Encode(auth)
}

func BcryptHasher(str string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(str), 10)
	helper.Errors(err, "bcrypthasher")
	return string(hash)
}
