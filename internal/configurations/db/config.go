package db

import (
	"fmt"
	"github.com/GenesisBlock3301/role_based_access_boilerplate_go/internal/configurations"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectionWithDB() {
	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable TimeZone=Asia/Shanghai",
		configurations.DBHost, configurations.DBUser, configurations.DBPassword, configurations.DBName, configurations.DBPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect configurations!")
	}
	// set default schema to `public`
	db.Migrator().CurrentDatabase()
	db.Set("gorm:schema_option", "public")
	db.AutoMigrate(&Product{})
	DB = db
}
