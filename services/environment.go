package services

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func InitEnv() {
	mode := GetEnvVar("ENV")
	if mode == "local" {
		e := godotenv.Load()
		if e != nil {
			log.Fatal(e)
		}
	}
}

func GetEnvVar(name string) (value string) {
	value = os.Getenv(name)
	if value == "" {
		log.Fatal("There is no `" + name + "` env variable!")
	}
	return
}
