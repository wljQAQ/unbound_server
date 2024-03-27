package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DatabaseCredentials 存储用来连接数据库的凭证
type DatabaseCredentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	Database string `json:"database"`
}

func ConnectDB(c *gin.Context) {
	var credentials DatabaseCredentials
	if err := c.BindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db, err := connectDataBase(&credentials)
	tableNames, err := getTableNames(db)
	if err != nil {
		log.Fatal("failed to get table names: ", err)
	}

	tableCount, err := getTableCount(db)
	if err != nil {
		log.Fatal("failed to get table count: ", err)
	}

	c.JSON(http.StatusOK, gin.H{"tables": tableNames, "tableCount": tableCount})
}

func connectDataBase(creds *DatabaseCredentials) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", creds.Host, creds.Port, creds.Username, creds.Password, creds.Database)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	return db, err
}

// 查询所有的表名
func getTableNames(db *gorm.DB) ([]string, error) {
	var tableNames []string
	err := db.Raw("SELECT table_name FROM information_schema.tables WHERE table_schema = 'public' AND table_type = 'BASE TABLE'").Scan(&tableNames).Error
	if err != nil {
		return nil, err
	}
	return tableNames, nil
}

// 获取表的数量
func getTableCount(db *gorm.DB) (int, error) {
	var count int
	err := db.Raw("SELECT COUNT(*) FROM information_schema.tables WHERE table_schema = 'public' AND table_type = 'BASE TABLE'").Scan(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}
