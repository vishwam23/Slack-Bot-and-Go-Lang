// package main

// import (
// 	"fmt"
// 	"time"
// )

// type LogLevel int

// const (
// 	Log LogLevel = iota
// 	Warning
// 	Error
// )

// type CustomLogger struct{}

// func (l *CustomLogger) Log(logType LogLevel, logMessage string) {
// 	logTypes := [...]string{"Log", "Warning", "Error"}
// 	logTime := time.Now().Format("2006/01/02 15:04:05")

// 	if logType >= Log && logType <= Error {
// 		fmt.Printf("%s %s :: %s\n", logTime, logTypes[logType], logMessage)
// 	} else {
// 		fmt.Printf("%s < UNSUPPORTED LOG TYPE DETECTED >\n", logTime)
// 	}
// }

// func main() {
// 	logger := CustomLogger{}

//		logger.Log(Log, "This is a normal log")
//		logger.Log(Warning, "This is a warning log")
//		logger.Log(Error, "This is an error log")
//		logger.Log(LogLevel(100), "This is an unsupported log type")
//	}
package main

import (
	"log"
)

var LogType = map[string]int{
	"NormalLog":  1,
	"WarningLog": 2,
	"ErrorLog":   3,
}

func CustomLogger(logMsg string, logType int) {
	if logType == LogType["NormalLog"] {
		log.Printf("Log :: %v", logMsg)
	} else if logType == LogType["WarningLog"] {
		log.Printf("Warning :: %v", logMsg)
	} else if logType == LogType["ErrorLog"] {
		log.Printf("Error :: %v", logMsg)
	} else {
		log.Printf("< UNSUPPORTED LOG TYPE DETECTED >")
	}
}

func main() {
	CustomLogger("this is a normal log", LogType["NormalLog"])
	CustomLogger("this is a Warning log", LogType["WarningLog"])
	CustomLogger("this is a Error log", LogType["ErrorLog"])
	CustomLogger("this is a default log", LogType["Nor"])
}
