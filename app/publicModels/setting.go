package publicModels

type Setting struct {
	Key        		string 		`gorm:"primarykey"`
	Value			string
}



func GetSetting(key string)(err error, value string){
	var data Setting
	result := db.DB.Model(Setting{}).Where("key = ?", key).First(&data)
	if result.Error != nil{
		return result.Error, ""
	}else {
		return nil,data.Value
	}
}

func SetSetting(key, value string)(err error){
	var data Setting
	result := db.DB.Model(Setting{}).Where("key = ?", key).First(&data)
	if result.Error != nil {
		result = db.DB.Model(Setting{}).Create(Setting{
			Key:   key,
			Value: value,
		})
	}else{
		data.Value = value
		result = db.DB.Model(&Setting{}).Updates(data)
	}
	return result.Error
}