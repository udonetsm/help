package helper

import (
	"crypto/sha512"
	"encoding/base64"
	"log"
	"os"
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

func Hasher(for_hsh string) string {
	hasher := sha512.New()
	hasher.Write([]byte(for_hsh))
	return base64.URLEncoding.EncodeToString(hasher.Sum(nil))
}
