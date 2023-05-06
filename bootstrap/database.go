package bootstrap

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDB(host, user, password, database, port, driver string ) *gorm.DB {
	endpoint := fmt.Sprintf(
		"%s://%s:%s@%s:%s/%s",
		driver,
		user,
		password,
		port,
		host,
		database,
	)
	db, err := gorm.Open(mysql.Open(endpoint), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Database connected")
	return db
}
