package elea

import (
	"fmt"
	"log"
	"os"
	"path"
	"time"
)

type LogLevel uint16

var DebugLogger *FileLogger
var TraceLogger *FileLogger
var InfoLogger *FileLogger
var WarnLogger *FileLogger
var ErrorLogger *FileLogger
var FatalLogger *FileLogger

// 日志文件输出目录，存放在进程目录log下面
func init() {
	folderPath := "log"
	if _, err := os.Stat(folderPath); os.IsNotExist(err) {
		err := os.Mkdir(folderPath, 0777)
		if err != nil {
			log.Printf("create log directory fail, err: %v\n", err)
			return
		}
	}
	DebugLogger = NewFileLogger(DEBUG, folderPath, "debug.log", 100*1024*1024)
	TraceLogger = NewFileLogger(TRACE, folderPath, "trace.log", 100*1024*1024)
	InfoLogger = NewFileLogger(INFO, folderPath, "info.log", 100*1024*1024)
	WarnLogger = NewFileLogger(WARN, folderPath, "warn.log", 100*1024*1024)
	ErrorLogger = NewFileLogger(ERROR, folderPath, "error.log", 100*1024*1024)
	FatalLogger = NewFileLogger(FATAL, folderPath, "fatal.log", 100*1024*1024)
}

// 日志级别
const (
	UNKNOWN LogLevel = iota
	TRACE
	DEBUG
	INFO
	WARN
	ERROR
	FATAL
)

// 获取日志的字符串格式
func getLevelStr(level LogLevel) string {
	switch level {
	case DEBUG:
		return "DEBUG"
	case TRACE:
		return "TRACE"
	case INFO:
		return "INFO"
	case WARN:
		return "WARN"
	case ERROR:
		return "ERROR"
	case FATAL:
		return "FATAL"
	case UNKNOWN:
		return "UNKNOWN"
	default:
		return ""
	}
}

// 定义日志的结构体
type FileLogger struct {
	Level       LogLevel
	filePath    string
	fileName    string
	maxFileSize int64
}

func NewFileLogger(level LogLevel, fp, fn string, size int64) *FileLogger {
	return &FileLogger{
		Level:       level,
		filePath:    fp,
		fileName:    fn,
		maxFileSize: size, // 暂时没有实现日志文件分割
	}
}

// 记录日志操作
func (f *FileLogger) Log(msg string) {
	nowStr := time.Now().Format("2006-01-02 15:04:05")
	levelStr := getLevelStr(f.Level)
	name := path.Join(f.filePath, f.fileName)

	file, err := os.OpenFile(name, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		log.Printf("open log file fail, err: %v\n", err)
		return
	}

	defer func() {
		err = file.Close()
		if err != nil {
			log.Printf("close log file fail, err: %v\n", err)
		}
	}()

	_, err = fmt.Fprintf(file, "[%s][%s] %s\r", nowStr, levelStr, msg)
	if err != nil {
		log.Printf("write log file fail, err: %v\n", err)
		return
	}
}
