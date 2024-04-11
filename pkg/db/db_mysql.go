package db

import (
	"github.com/jinzhu/gorm"
	"github.com/sharch/sim/config"
	"github.com/sharch/sim/pkg/logger"

	_ "github.com/go-sql-driver/mysql"
)

var (
	DB *gorm.DB
)

func init() {
	InitMysql(config.Config.MySQL)
}

// InitMysql 初始化MySQL
func InitMysql(dataSource string) {
	logger.Logger.Info("init mysql")
	var err error
	DB, err = gorm.Open("mysql", dataSource)
	if err != nil {
		panic(err)
	}
	DB.SingularTable(true)
	DB.LogMode(true)
	logger.Logger.Info("init mysql ok")
}
