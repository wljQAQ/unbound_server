package models

type Auth struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func CheckAuth(username, password string) bool {
	var auth Auth
	// 查询数据库，获取匹配的用户名和密码的Auth对象
	db.Select("id").Where(Auth{Username: username, Password: password}).First(&auth)

	// 如果查询到的Auth对象的ID大于0，表示存在该用户，返回true
	if auth.Id > 0 {
		return true
	}

	// 如果查询不到该用户，返回false
	return false
}
