package base

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

const (
	tempEnvironmentPath = "/home/liuzehao/GolandProjects/pasteProject/base/"
	devPath             = ""
)

func init() {
	var err error
	err = godotenv.Load(tempEnvironmentPath + "conf/envs/db.env")
	if err != nil {
		log.Fatalf("load environment variable failed on base.go : %v", err)
	}
}

func Load() (DSN string) {
	host := os.Getenv("MYSQL_HOST")
	port := os.Getenv("MYSQL_PORT")
	user := os.Getenv("MYSQL_USERNAME")
	password := os.Getenv("MYSQL_PASSWORD")
	db := os.Getenv("MYSQL_DBNAME")
	return user + ":" + password + "@tcp(" + host + ":" + port + ")" + "/" + db + "?charset=utf8mb4&parseTime=True&loc=Local"
}
