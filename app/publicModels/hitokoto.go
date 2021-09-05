package publicModels

type Hitokoto struct {
	Id			int
	Hitokoto	string		`gorm:"uniqueIndex"`
	Source 		string
	Catname		string
}


