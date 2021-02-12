package logging

import (
	"os"
)

type LogLevel uint8

const (
	DebugLevel = iota
	InfoLevel
	WarnLevel
	CritLevel
)

const LogLevelEnvVar = "LOG_LEVEL"

var globalLogLevel LogLevel

func init() {
	if val, ok := os.LookupEnv(LogLevelEnvVar); ok {
		switch val {
		case "DEBUG":
			globalLogLevel = DebugLevel
		case "INFO":
			globalLogLevel = InfoLevel
		case "WARN":
			globalLogLevel = WarnLevel
		case "CRIT":
			globalLogLevel = CritLevel
		default:
			globalLogLevel = InfoLevel
		}
	}
}
