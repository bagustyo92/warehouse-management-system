package config

import (
	"fmt"
	"os"

	log "github.com/bagustyo92/wms/middleware/logger"
	"github.com/joho/godotenv"
)

var (
	AppPort      string
	DBName       string
	DBUsername   string
	DBPassword   string
	DBPort       string
	DBUrl        string
	DBConnection string
)

func getEnv(keyEnv string, fileEnv map[string]string) string {
	envVal, ok := os.LookupEnv(keyEnv)
	if !ok {
		return fileEnv[keyEnv]
	}
	return envVal
}

// InitApp will get all the important env from env file
// Do not include env file if you want get the env from host
func InitApp(envPath string) {
	envFile, err := godotenv.Read(envPath)
	if err != nil {
		log.MakeLogEntry(nil).Panic(err)
	}

	AppPort = getEnv("APP_PORT", envFile)
	// fmt.Println(APP_PORT)

	// DB Env
	DBName = getEnv("DB_NAME", envFile)
	DBUsername = getEnv("DB_USERNAME", envFile)
	DBPassword = getEnv("DB_PASSWORD", envFile)
	DBPort = getEnv("DB_PORT", envFile)
	DBUrl = getEnv("URL_DB", envFile)

	if DBUsername == "" || DBPassword == "" {
		DBConnection = fmt.Sprintf(":@tcp(%s:%s)/%s", DBUrl, DBPort, DBName)
	} else {
		DBConnection = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", DBUsername, DBPassword, DBUrl, DBPort, DBName)
	}

}
