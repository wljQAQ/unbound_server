package models

import (
	"fmt"
	"log"
	"unbound/pkg/setting"

	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

type Model struct {
	ID         int
	CreatedOn  int
	ModifiedOn int
}

func init() {
	var (
		err                                error
		dbName, user, password, host, port string
	)

	sec, err := setting.Cfg.GetSection("database")

	if err != nil {
		log.Fatal(2, "Fail to get section 'database': %v", err)
	}

	// dbType = sec.Key("TYPE").String()
	dbName = sec.Key("NAME").String()
	user = sec.Key("USER").String()
	password = sec.Key("PASSWORD").String()
	host = sec.Key("HOST").String()
	port = sec.Key("PORT").String()
	// tablePrefix = sec.Key("TABLE_PREFIX").String()
	// dsn := "host=localhost port=5432 user=postgres password=123456 dbname=blog sslmode=disable"
	// dsn := "host=localhost user=postgres password=123456 dbname=blog port=5432 sslmode=disable"
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Println(err)
	}

	//如果没有就创建这个表
	db.AutoMigrate(&Tag{})

	sqlDB, err := db.DB()

	if err != nil {
		log.Println(err)
	}

	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)
	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)
	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(10 * time.Second) // 10秒钟
}
