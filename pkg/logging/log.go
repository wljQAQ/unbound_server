package logging

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

type Level int

var (
	F *os.File

	DefaultPrefix      = ""
	DefaultCallerDepth = 2

	logger     *log.Logger
	logPrefix  = ""
	levelFlags = []string{"DEBUG", "INFO", "WARN", "ERROR", "FATAL"}
)

const (
	DEBUG Level = iota
	INFO
	WARNING
	ERROR
	FATAL
)

func init() {
	// 获取日志文件的完整路径
	filePath := getLogFileFullPath()

	// 打开日志文件
	F = openLogFile(filePath)

	// 创建一个新的日志记录器，并将标准输出、默认前缀和标准标志传递给它
	logger = log.New(F, DefaultPrefix, log.LstdFlags)
}

func Debug(v ...interface{}) {
	setPrefix(DEBUG)
	logger.Println(v)
}

func Info(v ...interface{}) {
	setPrefix(INFO)
	logger.Println(v)
}

func Warn(v ...interface{}) {
	setPrefix(WARNING)
	logger.Println(v)
}

func Error(v ...interface{}) {
	setPrefix(ERROR)
	logger.Println(v)
}

func Fatal(v ...interface{}) {
	setPrefix(FATAL)
	logger.Fatalln(v)
}

func setPrefix(level Level) {
	// 获取调用者的文件名、行号和是否成功获取到信息
	_, file, line, ok := runtime.Caller(DefaultCallerDepth)

	// 如果成功获取到信息
	if ok {
		// 设置日志前缀为“[日志级别][文件名:行号]”
		logPrefix = fmt.Sprintf("[%s][%s:%d]", levelFlags[level], filepath.Base(file), line)
	} else {
		// 如果未成功获取到信息，则只设置日志前缀为“[日志级别]”
		logPrefix = fmt.Sprintf("[%s]", levelFlags[level])
	}

	// 设置日志记录器的日志前缀
	logger.SetPrefix(logPrefix)
}
