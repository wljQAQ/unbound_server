package models

import (
	"fmt"
	"time"
	"unbound/pkg/setting"

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
	// 获取数据库DNS
	dns := getDataBaseDNS()
	println(dns)

	// 使用postgres驱动打开数据库连接
	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{})

	// 如果打开数据库连接出错，则抛出异常
	if err != nil {
		panic(fmt.Sprintf("连接数据库失败: %v", err))
	}

	// 获取底层SQL数据库连接
	sqlDB, err := db.DB()

	// 如果设置数据库连接池出错，则抛出异常
	if err != nil {
		panic(fmt.Sprintf("设置数据库连接池错误: %v", err))
	}

	// 设置连接的最大空闲时间
	sqlDB.SetConnMaxIdleTime(10)

	// 设置连接池的最大打开连接数
	sqlDB.SetMaxOpenConns(100)

	// 设置连接的最大生命周期
	sqlDB.SetConnMaxLifetime(10 * time.Second)
}

func getDataBaseDNS() string {
	db := setting.Server.Db
	// return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=require", db.Host, db.Port, db.User, db.Password, db.Name)
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s", db.Host, db.Port, db.User, db.Password, db.Name)
}
