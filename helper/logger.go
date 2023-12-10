package helper

import (
	"fmt"
	"gin_admin/define"
	"github.com/go-ini/ini"
	"os"
	"path"
	"runtime"
	"strconv"
	"time"
)

func dataFormat(showDetail bool, format string, v ...interface{}) string {
	prefix := "[" + define.FrameName + "] " + time.Now().Format(define.DateTimeLayout) + " "
	if showDetail {
		_, file, line, _ := runtime.Caller(2)
		prefix += "file: " + file + " line: " + strconv.Itoa(line) + " ==> "
	}
	return prefix + fmt.Sprintf(format, v...)
}

// Info INFO
func Info(format string, v ...interface{}) {
	file := InitFile()
	if file == nil {
		return
	}
	_, err := file.WriteString(fmt.Sprintf("\033[32m%s\033[0m\n", dataFormat(true, format, v...)))
	if err != nil {
		fmt.Println("Log Error: ", err)
		return
	}
	fmt.Printf("\033[32m%s\033[0m\n", dataFormat(false, format, v...))
}

// Error ERROR
func Error(format string, v ...interface{}) {
	file := InitFile()
	if file == nil {
		return
	}
	_, err := file.WriteString(fmt.Sprintf("\033[31m%s\033[0m\n", dataFormat(true, format, v...)))
	if err != nil {
		fmt.Println("Log Error: ", err)
		return
	}
	fmt.Printf("\033[31m%s\033[0m\n", dataFormat(true, format, v...))
}

func InitFile() *os.File {
	file, err := ini.Load("conf/config.ini")
	if err != nil {
		panic(err)
	}
	logFilePath := file.Section("log_file").Key("FilePath").String()
	logFileName := file.Section("log_file").Key("FileName").String()
	fileName := path.Join(logFilePath, logFileName)
	src, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		fmt.Println("Create LogFile Error: ", err)
		return nil
	}
	return src
}
