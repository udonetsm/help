package helper

import (
	"crypto/sha512"
	"encoding/base64"
	"log"
	"os"

	"golang.org/x/crypto/bcrypt"
)

func Errors(err error, funcname string) {
	log.Println(err, "<--", funcname)
}

func PanicCapture(funcname string) {
	if err := recover(); err != nil {
		log.Println("Panic <--", funcname)
		return
	}
}

func Home() string {
	home, err := os.UserHomeDir()
	Errors(err, "oshomedir")
	return home
}

func Sha512Hasher(for_hsh string) string {
	hasher := sha512.New()
	hasher.Write([]byte(for_hsh))
	return base64.URLEncoding.EncodeToString(hasher.Sum(nil))
}

func BcryptHasher(str string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(str), 10)
	if err != nil {
		log.Fatal(err)
	}
	return string(hash)
}
