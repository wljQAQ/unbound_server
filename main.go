package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func main() {

	// conn, err := pgx.Connect(context.Background(), os.Getenv("user=postgres password=123456 dbname=postgres  sslmode=disable"))
	// if err != nil {
	// 	fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
	// 	os.Exit(1)
	// }
	// defer conn.Close(context.Background())

	// var greeting string
	// err = conn.QueryRow(context.Background(), "select 'Hello, world!'").Scan(&greeting)
	// if err != nil {
	// 	fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
	// 	os.Exit(1)
	// }

	// fmt.Println(greeting)

	//连接数据库
	db, err := sql.Open("postgres", "user=postgres password=123456 dbname=test  sslmode=disable")
	if err != nil {
		fmt.Println("链接pg数据失败！ err：", err.Error())
		return
	}
	defer db.Close()
	fmt.Println("连接pg成功")
	err = db.Ping()

	fmt.Println(db.)
	if err != nil {
		fmt.Println("ping 数据出现错误！ er：", err.Error())
		fmt.Println("ping 数据出现错误！ er：", err.Error())
		return
	}

	// r := gin.Default()
	// r.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "2222",
	// 	})
	// })
	// r.Run(":3000") // listen and serve on 0.0.0.0:8080
}
