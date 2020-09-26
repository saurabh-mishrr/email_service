package helpers

import (
	"io"
	"log"
	"os"

	// Import godotenv
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// use godot package to load/read the .env file and
// return the value of the key
func DotEnvVal(key string) string {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

func WriteLog() {
	gin.DisableConsoleColor()
	os.Mkdir(DotEnvVal("LOG_PATH"), 0755)
	f, _ := os.Create(DotEnvVal("LOG_PATH") + "/" + DotEnvVal("LOG_FILE"))
	gin.DefaultWriter = io.MultiWriter(f)

}
