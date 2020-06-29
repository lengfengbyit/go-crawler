package util

import "github.com/joho/godotenv"

func init() {
	err := godotenv.Load("./.env")
	if err != nil {
		panic(err)
	}
}
