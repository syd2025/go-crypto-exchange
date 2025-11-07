package database

import (
	"time"

	"mscoin-common/msdb"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type MysqlConfig struct {
	Database string
}

func ConnMySql(dsn string) *msdb.MsDB {
	var err error
	_db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		panic("连接数据库失败, error =" + err.Error())
	}

	db, _ := _db.DB()
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxLifetime(time.Hour)
	return &msdb.MsDB{
		Conn: _db,
	}
}
