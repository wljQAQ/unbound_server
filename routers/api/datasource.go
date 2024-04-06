package api

import (
	"fmt"
	"unbound/models"
	"unbound/routers/response"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// DatabaseCredentials 存储用来连接数据库的凭证
type DatabaseCredentials struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Host     string `json:"host" binding:"required"`
	Port     int    `json:"port" binding:"required"`
	Database string `json:"database" binding:"required"`
	Ssl      bool   `json:"ssl"`
}

func ConnectDB(c *gin.Context) {
	var credentials DatabaseCredentials
	if err := c.ShouldBindJSON(&credentials); err != nil {
		response.ErrorValidation(c, err)
		return
	}

	_, err := connectDataBase(&credentials)

	if err != nil {
		response.Error(c, "连接数据库失败:"+err.Error())
		return
	}

	response.Ok(c, nil)
}

func GetTableSchema(c *gin.Context) {
	schema, err := models.GetTableSchema(DB)

	if err != nil {
		response.Error(c, err.Error())
		return
	}

	response.Ok(c, schema)
}

func connectDataBase(creds *DatabaseCredentials) (*gorm.DB, error) {
	sslMode := "disable"
	if creds.Ssl {
		sslMode = "require"
	}

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s", creds.Host, creds.Port, creds.Username, creds.Password, creds.Database, sslMode)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	DB = db
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
