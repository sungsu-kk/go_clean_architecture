package config

import (
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB

// 데이터베이스 연결 구조체
type Database struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	Username string `json:"user"`
	Password string `json:"password"`
	Database string `json:"database"`
}

// 데이터베이스 설정정보 출력
func (db *Database) String() string {
	return "Host: " + db.Host + ", Port: " + db.Port + ", Username: " + db.Username + ", Password: " + db.Password + ", Database: " + db.Database
}

// 데이터베이스 연결 정보
func (db *Database) ConnectInfo() string {
	return db.Username + ":" + db.Password + "@tcp(" + db.Host + ":" + db.Port + ")/" + db.Database + "?charset=utf8mb4&parseTime=True&loc=Local"
}

// 데이터베이스 연결 (mariadb / mysql)
func (db *Database) Connect() {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: false,       // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,       // Disable color
		},
	)
	dsn := db.ConnectInfo()
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		panic(err)
	}
	DB = database
}
