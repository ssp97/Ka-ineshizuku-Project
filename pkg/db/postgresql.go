package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type ConfigPostgresql struct {
	DSN         string         // orm DSN
	//Debug       bool           // orm is debug
	//Active      int            // pool
	//Idle        int            // pool
	//IdleTimeout xtime.Duration // connect max life time.
}



func connectPostgresql(cSqlite  *ConfigPostgresql)  (*gorm.DB, error){
	db, err := gorm.Open(postgres.Open(cSqlite.DSN), &gorm.Config{})
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