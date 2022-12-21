package Config

import (
	"fmt"

	"gorm.io/gorm"
)

var DB *gorm.DB

type DbConfig struct {
	Host     string
	Port     int
	User     string
	DBName   string
	Password string
}

func BuildDBConfig() *DbConfig {
	dbConfig := DbConfig{
		Host:     "0.0.0.0",
		Port:     3232,
		User:     "root",
		DBName:   "todos",
		Password: "mypassword",
	}

	return &dbConfig
}

func DbURL(dbConfig *DbConfig) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DBName,
	)
}
