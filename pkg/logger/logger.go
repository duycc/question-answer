//===--------------------------- pkg/logger/logger.go - [question-answer] --------------------------------*- Go -*-===//
// Brief :
//
//
// Author: YongDu
// Date  : 2022-04-24
//===--------------------------------------------------------------------------------------------------------------===//

package logger

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"question-answer/pkg/config"
	"time"
)

var (
	debugLogger *log.Logger
	infoLogger  *log.Logger
	warnLogger  *log.Logger
	errorLogger *log.Logger
)

func init() {
	logFilePath := filepath.Join(config.LogDir, fmt.Sprintf("share-%s.log", time.Now().Format("2006-01-02")))
	logFile, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0644)
	if err != nil {
		log.Fatalln("fail to create share.log")
	}
	infoLogger = log.New(logFile, "[Info]", log.LstdFlags|log.Lshortfile)
	warnLogger = log.New(logFile, "[Warn]", log.LstdFlags|log.Lshortfile)
	errorLogger = log.New(logFile, "[Error]", log.LstdFlags|log.Lshortfile)
	debugLogger = log.New(logFile, "[Debug]", log.LstdFlags|log.Lshortfile)
}

func Info(format string, v ...interface{}) {
	if v == nil {
		infoLogger.Println(format)
	} else {
		infoLogger.Printf(format+"\n", v)
	}
}

func Warn(format string, v ...interface{}) {
	if v == nil {
		warnLogger.Println(format)
	} else {
		warnLogger.Printf(format+"\n", v)
	}
}

func Error(format string, v ...interface{}) {
	if v == nil {
		errorLogger.Println(format)
	} else {
		errorLogger.Printf(format+"\n", v)
	}
}

func Debug(format string, v ...interface{}) {
	if v == nil {
		debugLogger.Println(format)
	} else {
		debugLogger.Printf(format+"\n", v)
	}
}

func Panic(v ...interface{}) {
	log.Panic(v...)
}

func Fatal(v ...interface{}) {
	log.Fatal(v...)
}
