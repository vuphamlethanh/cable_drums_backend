package main

import (
	"cable_drums_backend/src/entities"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=localhost user=postgres password=postgrespw dbname=GoDriveDb port=32768 sslmode=disable TimeZone=Asia/Shanghai"
	db, _ := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	db.AutoMigrate(&entities.Planner{}, &entities.Supplier{}, &entities.Contractor{}, &entities.Contract{}, &entities.WithdrawRequest{})
}
