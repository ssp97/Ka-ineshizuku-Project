package setu

type setu struct {
	Id				int			`gorm:"primarykey"`
	Pid 			int			`gorm:"uniqueIndex"`
	P   			int
	Title 			string		`gorm:"index"`
	UserId 			int
	UserAccount 	string
	UserName		string
	Url 			string
	R18				int			`gorm:"index"`
	Width			int
	Height			int
	Tags  			string
	TagsTranslated	string
	Caption			string
	setuTag 		[]setuTag 	`gorm:"foreignKey:Pid;references:Pid;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type setuTag struct {
	ID        		uint 		`gorm:"primarykey"`
	Pid			  	int 		`gorm:"index;uniqueIndex:idx_tag_id"`
	Tag 			string		`gorm:"index;uniqueIndex:idx_tag_id"`
}

type setuTagTranslated struct {
	ID        		uint 		`gorm:"primarykey"`
	Zh				string		`gorm:"index;uniqueIndex:idx_zh_src"`
	Src				string		`gorm:"index;uniqueIndex:idx_zh_src"`
}


