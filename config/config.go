package config

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

var RDB *redis.Client

func Init() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai",
		"127.0.0.1",
		"postgres",
		"sqltest",
		"novaro",
		5432,
	)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Database connection test failed:", err)
	}

	RDB = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
}
