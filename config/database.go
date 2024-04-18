package config

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() (*gorm.DB, error) {
	dbEnv, err := GetDBEnv()
	if err != nil {
		fmt.Println("Failed to get Env: ", err)
		return nil, err
	}
	dns := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", dbEnv.User, dbEnv.Password, dbEnv.Host, dbEnv.Port, dbEnv.DBName)
	db, err := gorm.Open(postgres.Open(dns))
	if err != nil {
		fmt.Println("Failed connect to database: ", err)
		return nil, err
	}
	fmt.Println("Success Connect to Database")
	return db, nil
}
