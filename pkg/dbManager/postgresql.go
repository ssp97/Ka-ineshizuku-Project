package dbManager

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

type ConfigPostgresql struct {
	DSN         string         // orm DSN
	//Debug       bool           // orm is debug
	//Active      int            // pool
	//Idle        int            // pool
	//IdleTimeout xtime.Duration // connect max life time.
}



func connectPostgresql(cSqlite  *ConfigPostgresql)  (*gorm.DB, error){
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second,   // 慢 SQL 阈值
			LogLevel:      logger.Warn, // Log level
			Colorful:      true,         // 禁用彩色打印
		},
	)

	db, err := gorm.Open(postgres.Open(cSqlite.DSN), &gorm.Config{
		PrepareStmt: false,
		Logger: newLogger,
	})
	if err != nil {
		//panic("failed to connect database")
		return nil,err
	}
	return db, nil
}

func NewPostgresql(c *ConfigPostgresql) *ORM {
	d, err := connectPostgresql(c)
	if err != nil {
		panic(err)
	}
	return &ORM{DB: d, cPostgresql: c}
}