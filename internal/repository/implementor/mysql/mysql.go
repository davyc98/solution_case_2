package mysql

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//Connect connection DB
func Connect(host, database string, port int, username, password string) (db *gorm.DB, err error) {

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, port, database)
	db, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		return db, err
	}

	log.Println("Mysql Database Connected")

	return db, err
}
