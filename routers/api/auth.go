package api

import (
	"fmt"
	"log"
	"net/http"
	"unbound/models"
	"unbound/pkg/e"
	"unbound/pkg/util"

	"github.com/beego/beego/v2/core/validation"
	"github.com/gin-gonic/gin"
)

type auth struct {
	Username string `valid:"Required; MaxSize(50)"`
	Password string `valid:"Required; MaxSize(50)"`
}

func GetAuth(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	vaild := validation.Validation{}
	a := auth{Username: username, Password: password}
	code := e.INVALID_PARAMS
	data := make(map[string]interface{})

	ok, _ := vaild.Valid(&a)
	if ok {
		isExist := models.CheckAuth(username, password)
		if isExist {
			token, err := util.GenerateToken(username, password)
			fmt.Println("err", err)
			if err != nil {
				code = e.ERROR_AUTH_TOKEN
			} else {
				code = e.SUCCESS
				data["token"] = token
			}
		} else {
			code = e.ERROR_AUTH
		}
	} else {
		for _, err := range vaild.Errors {
			log.Println(err.Key, err.Message)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"data": data,
		"msg":  e.GetMsg(code),
	})
}
