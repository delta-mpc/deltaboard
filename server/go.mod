module deltaboard-server

go 1.14

require (
	github.com/ethereum/go-ethereum v1.9.25
	github.com/gin-contrib/static v0.0.0-20191128031702-f81c604d8ac2
	github.com/gin-gonic/gin v1.6.3
	github.com/go-gormigrate/gormigrate/v2 v2.0.0
	github.com/go-sql-driver/mysql v1.5.0
	github.com/google/uuid v1.1.2 // indirect
	github.com/gorilla/sessions v1.2.1
	github.com/loopfz/gadgeto v0.9.0
	github.com/quasoft/memstore v0.0.0-20191010062613-2bce066d2b0b
	github.com/sirupsen/logrus v1.4.2
	github.com/spf13/cobra v0.0.6
	github.com/spf13/viper v1.4.0
	github.com/wI2L/fizz v0.13.4
	github.com/wader/gormstore/v2 v2.0.0
	google.golang.org/protobuf v1.25.0 // indirect
	gorm.io/driver/mysql v1.0.4
	gorm.io/driver/sqlite v1.1.4
	gorm.io/gorm v1.21.10
)

replace github.com/loopfz/gadgeto v0.9.0 => github.com/we-miks/gadgeto v0.10.2-0.20200623025716-393d1a68186b
