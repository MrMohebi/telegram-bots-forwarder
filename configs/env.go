package configs

import (
	"github.com/joho/godotenv"
	"os"
	"telegram-bots-forwarder/common"
)

var IsInitOnce = false

func EnvSetup() {
	if !IsInitOnce {
		err := godotenv.Load()
		common.IsErr(err, "Error loading .env file")
		IsInitOnce = true
	} else {
		println("initialized envs once")
	}
}

func EnvMongoURI() string {
	return os.Getenv("MONGO_FINAL_URI")
}

func EvnMongoDB() string {
	return os.Getenv("MONGO_DB")
}
