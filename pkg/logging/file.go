package logging

import (
	"fmt"
	"log"
	"os"
	"time"
)

var (
	logSavePath = "runtime/logs/"
	logSaveName = "log"
	logSaveExt  = "log"
	TimeFormat  = "20060102"
)

func getLogPath() string {
	return fmt.Sprintf("%s", logSavePath)
}

func getLogFileFullPath() string {
	prefixPath := getLogPath()
	suffixPath := fmt.Sprintf("%s%s.%s", logSaveName, time.Now().Format(TimeFormat), logSaveExt)
	return fmt.Sprintf("%s%s", prefixPath, suffixPath)
}

// openLogFile 函数用于打开或创建日志文件
// 参数filePath为文件路径
// 返回值为*os.File类型的文件句柄
func openLogFile(filePath string) *os.File {
	// 检查文件是否存在
	_, err := os.Stat(filePath)
	switch {
	// 如果文件不存在
	case os.IsNotExist(err):
		// 创建目录
		mkDir()

	// 如果文件存在但权限不足
	case os.IsPermission(err):
		// 输出错误信息并退出程序
		log.Fatalf("Permission :%v", err)

	}

	// 打开文件，如果不存在则创建，并以追加和只写模式打开，设置文件权限为0644
	handle, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		// 如果打开文件失败，输出错误信息并退出程序
		log.Fatalf("Fail to OpenFile :%v", err)
	}

	// 返回文件句柄
	return handle
}

func mkDir() {
	// 获取当前工作目录
	dir, _ := os.Getwd()

	// 在当前工作目录下创建日志文件的完整路径
	err := os.MkdirAll(dir+"/"+getLogFileFullPath(), os.ModePerm)

	// 如果创建目录失败，则抛出异常
	if err != nil {
		panic(err)
	}
}
