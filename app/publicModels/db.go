package publicModels

import "github.com/ssp97/Ka-ineshizuku-Project/pkg/dbManager"

func Init() {
	db := dbManager.GetDb(dbManager.DEFAULT_DB_NAME)
	db.DB.AutoMigrate(Hitokoto{})
}
