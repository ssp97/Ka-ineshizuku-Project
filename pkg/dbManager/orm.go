package dbManager

import "gorm.io/gorm"

var DEFAULT_DB_NAME = "main"

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

var DbMap = map[string]*ORM{}

func SetDb(name string, orm *ORM) {
	DbMap[name] = orm
}

func GetDb(name string)(orm *ORM) {
	orm, ok := DbMap[name]
	if !ok{
		return nil
	}
	return
}

//func GetTableName(db *ORM, m interface{})(){
//	db.DB.
//	db.DB.NewScope(model).GetModelStruct().TableName(db)
//}

func New(c *OrmConfig) *ORM{

	switch c.DbType {
	case "sqlite":
		return NewSqlite(&c.Sqlite)
	case "postgresql":
		return NewPostgresql(&c.Postgresql)
	}
	return nil
}