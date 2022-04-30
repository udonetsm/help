package helper

import (
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
