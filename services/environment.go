package services

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func InitEnv() {
	e := godotenv.Load()
	if e != nil {
		fmt.Print(e)
	}
}

func GetEnvVar(name string) string {
	value := os.Getenv(name)
	if value == "" {
		log.Fatal("There is no `" + name + "` env variable!")
	}
	return value
}
