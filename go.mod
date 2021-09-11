module github.com/ssp97/Ka-ineshizuku-Project

go 1.16

require (
	github.com/BurntSushi/toml v0.3.1
	github.com/FloatTech/AnimeAPI v0.0.0-20210713044920-63367fe18ccd
	github.com/FloatTech/ZeroBot-Plugin-Timer v1.2.4
	github.com/StackExchange/wmi v0.0.0-20210224194228-fe8f1750fd46 // indirect
	github.com/adamzy/cedar-go v0.0.0-20170805034717-80a9c64b256d // indirect
	github.com/gin-gonic/gin v1.3.0 // indirect
	github.com/go-gl/gl v0.0.0-20210426225639-a3bfa832c8aa // indirect
	github.com/go-ole/go-ole v1.2.5 // indirect
	github.com/golang/freetype v0.0.0-20170609003504-e2365dfdc4a0
	github.com/gorilla/websocket v1.4.2
	github.com/json-iterator/go v1.1.11
	github.com/lestrrat-go/file-rotatelogs v2.4.0+incompatible
	github.com/lestrrat-go/strftime v1.0.5 // indirect
	github.com/liuzl/cedar-go v0.0.0-20170805034717-80a9c64b256d // indirect
	github.com/liuzl/da v0.0.0-20180704015230-14771aad5b1d // indirect
	github.com/liuzl/gocc v0.0.0-20200216023908-f8cb162baf44
	github.com/nfnt/resize v0.0.0-20180221191011-83c6a9932646
	github.com/orcaman/writerseeker v0.0.0-20200621085525-1d3f536ff85e
	github.com/pkg/errors v0.9.1 // indirect
	github.com/shirou/gopsutil v3.21.6+incompatible
	github.com/sirupsen/logrus v1.8.1
	github.com/sosedoff/pgweb v0.11.8
	github.com/tfriedel6/canvas v0.12.1
	github.com/tidwall/gjson v1.8.1
	github.com/tidwall/pretty v1.2.0 // indirect
	github.com/tklauser/go-sysconf v0.3.6 // indirect
	github.com/wdvxdr1123/ZeroBot v1.2.2
	github.com/yanyiwu/gojieba v1.1.2
	golang.org/x/image v0.0.0-20210628002857-a66eb6448b8d
	golang.org/x/net v0.0.0-20210510120150-4163338589ed // indirect
	golang.org/x/sys v0.0.0-20210809222454-d867a43fc93e // indirect
	golang.org/x/text v0.3.7 // indirect
	golang.org/x/tools v0.1.3-0.20210608163600-9ed039809d4c // indirect
	gopkg.in/yaml.v2 v2.2.8 // indirect
	gorm.io/driver/postgres v1.1.0
	gorm.io/driver/sqlite v1.1.4
	gorm.io/gorm v1.21.12
	modernc.org/sqlite v1.11.1
	modernc.org/strutil v1.1.1
)

replace (
	github.com/wdvxdr1123/ZeroBot v1.2.2 => github.com/ssp97/ZeroBot v1.2.4
	github.com/yanyiwu/gojieba v1.1.2 => github.com/ttys3/gojieba v1.1.3
)
