package database

import (
	"fmt"
	"strconv"

	. "data-collector/utils"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	log "github.com/sirupsen/logrus"
)

var Conn *gorm.DB

func init() {
	//透過 Getenv 來讀取 .env
	//username := os.Getenv("")
	//password := os.Getenv("")
	//dbName := os.Getenv("")
	//dbHost := os.Getenv("")

	// init DB connection variable

	config, nil := ParserParseEnvCfg("env\\config.yaml")
	username := config.Username
	password := config.Password
	dbName := config.DBName
	dbHost := config.DBHost
	dbPort := strconv.Itoa(config.DBPort)

	//連結 db
	dbUri := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, dbPort, username, dbName, password)
	log.Info(fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, dbPort, username, dbName, MD5_SALT(password, "willis")))

	//錯誤攔截與建立連接
	var err error
	Conn, err = gorm.Open("postgres", dbUri)
	if err != nil {
		log.Fatal(err)
	}

	if Conn.Error != nil {
		log.Fatal("database error %v", Conn.Error)
	}
}
