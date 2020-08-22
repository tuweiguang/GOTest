package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
	"time"
)

const (
	LoggerType_Normal = 0
	LoggerType_Error  = 1
	LoggerType_Debug  = 2
)

// server log stuff begin
type LoggerFile struct {
	Type   int
	Prefix string
	writer *os.File
}

type GlobalWriterManager struct {
	ErrorContent []string
	LogContent   []string
	DebugContent []string
	mutex        sync.Mutex
}

func (m *GlobalWriterManager) AddDebugLog(content string) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	m.DebugContent = append(m.DebugContent, content)
}

func (m *GlobalWriterManager) AddErrorLog(content string) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	m.ErrorContent = append(m.ErrorContent, content)
}

func (m *GlobalWriterManager) AddNormalLog(content string) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	m.LogContent = append(m.LogContent, content)
}

var GlobalLogger *GlobalWriterManager

func (l *LoggerFile) Write(p []byte) (n int, err error) {
	if l.Type == LoggerType_Debug {
		return len(p), nil
	}

	//if l.Prefix == "" {
	//	var Addr string
	//	if LocalAgentConfig != nil {
	//		Addr = LocalAgentConfig.AddrHost
	//	}
	//	l.Prefix = fmt.Sprintf("%s@%s ",GetAgentName(ServerRunningMode),Addr)
	//}

	if GlobalLogger != nil {
		content := l.Prefix + string(p)
		if l.Type == LoggerType_Debug {
			GlobalLogger.AddDebugLog(content)
		} else if l.Type == LoggerType_Error {
			GlobalLogger.AddErrorLog(content)
		} else if l.Type == LoggerType_Normal {
			GlobalLogger.AddNormalLog(content)
		}
	}

	fmt.Println("========> ", string(p))
	return l.writer.Write(p)
}

var KitLogger *log.Logger
var KitLogFile *LoggerFile
var KitError *log.Logger
var KitErrFile *LoggerFile
var KitDebugFile *LoggerFile
var KitDebugger *log.Logger

const ServerBasePath = "E:\\GoTest\\src\\opensource_logger"

func InitLogger() {
	var prefix string
	path := ServerBasePath + "/var"
	_, e := os.Stat(path)
	if e != nil {
		os.Mkdir(path, os.ModePerm)
	}

	path = ServerBasePath + "/var/locallog"
	_, e = os.Stat(path)
	if e != nil {
		os.Mkdir(path, os.ModePerm)
	}

	path = ServerBasePath + "/var/locallog/log"
	_, e = os.Stat(path)
	if e != nil {
		os.Mkdir(path, os.ModePerm)
	}

	path = path + "/" + "tuweiguang"
	_, ee := os.Stat(path)
	if ee != nil {
		os.Mkdir(path, os.ModePerm)
	}

	path = path + "/" + strings.Replace(time.Now().Format("2006-01-02 15_04_05"), " ", "_", -1)
	KitLogFile = new(LoggerFile)
	KitLogFile.Type = LoggerType_Normal
	var err error
	KitLogFile.writer, err = os.OpenFile(path+fmt.Sprintf("_%v", 0)+".log", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		fmt.Printf(" create log file failed! %v,%v\n", err, path)
		os.Exit(1)
	}
	KitLogger = log.New(KitLogFile, prefix, log.Ldate|log.Ltime|log.Lshortfile)

	KitDebugFile = new(LoggerFile)
	KitDebugFile.Type = LoggerType_Debug
	KitDebugFile.writer, err = os.OpenFile(path+fmt.Sprintf("_%v", 0)+".debug", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		fmt.Printf(" create log file failed! %v,%v\n", err, path)
		os.Exit(1)
	}
	KitDebugger = log.New(KitDebugFile, prefix, log.Ldate|log.Ltime|log.Lshortfile)

	//error log
	path = ServerBasePath + "/var/locallog/errorlog"
	_, e = os.Stat(path)
	if e != nil {
		os.Mkdir(path, os.ModePerm)
	}

	path = path + "/" + "tuweiguang"
	_, ee = os.Stat(path)
	if ee != nil {
		os.Mkdir(path, os.ModePerm)
	}

	path = path + "/" + strings.Replace(time.Now().Format("2006-01-02 15_04_05")+fmt.Sprintf("_%v", 0)+".log", " ", "_", -1)
	KitErrFile = new(LoggerFile)
	KitErrFile.Type = LoggerType_Error
	KitErrFile.writer, err = os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		fmt.Printf(" create log file failed! %v,%v\n", err, path)
		os.Exit(1)
	}
	KitError = log.New(KitErrFile, prefix, log.Ldate|log.Ltime|log.Lshortfile)

	KitLogger.Printf("Use Base Path %v", ServerBasePath)

}

// server log stuff end

func main() {
	InitLogger()
	fmt.Println("init logger")

	KitLogger.Println("hello world")
}
