package database

import (
	"fmt"

	"github.com/StackItHQ/pes-devmegablaster/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DatabaseSvc struct {
	DB *gorm.DB
}

func New() *DatabaseSvc {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		config.Cfg.DbConfig.HOST,
		config.Cfg.DbConfig.USER,
		config.Cfg.DbConfig.PASS,
		config.Cfg.DbConfig.DBNAME,
		config.Cfg.DbConfig.PORT,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	return &DatabaseSvc{
		DB: db,
	}
}
