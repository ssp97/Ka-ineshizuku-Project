package db

import "gorm.io/gorm"

type ORM struct {
	DB *gorm.DB
	cSqlite  *ConfigSqlite
	cPostgresql *ConfigPostgresql
}

type OrmConfig struct {
	DbType  	string
	Sqlite  	ConfigSqlite
	Postgresql ConfigPostgresql
}

func New(c *OrmConfig) *ORM{

	switch c.DbType {
	case "sqlite":
		return NewSqlite(&c.Sqlite)
	case "postgresql":
		return NewPostgresql(&c.Postgresql)
	}
	return nil
}