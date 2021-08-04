package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type ConfigSqlite struct {
	Path string
}



func connectSqlite(cSqlite  *ConfigSqlite)  (*gorm.DB, error){
	db, err := gorm.Open(sqlite.Open(cSqlite.Path), &gorm.Config{})
	if err != nil {
		//panic("failed to connect database")
		return nil,err
	}
	return db, nil
}

func NewSqlite(c *ConfigSqlite) *ORM {
	d, err := connectSqlite(c)
	if err != nil {
		panic(err)
	}
	return &ORM{DB: d, cSqlite: c}
}