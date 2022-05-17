package utils

import (
	"flag"
)

// long flag
type Commands struct {
	InitMockTable    *bool
	FlushAllProgress *bool
	HTTPServer       *bool
}

func GetFlag() Commands {
	initMockTable := flag.Bool("init", false, "init mock table")
	flushAllProgress := flag.Bool("flush", false, "flush all progress")
	httpServer := flag.Bool("http-server", false, "start http server")

	flag.Parse()

	commands := Commands{
		InitMockTable:    initMockTable,
		FlushAllProgress: flushAllProgress,
		HTTPServer:       httpServer,
	}
	return commands
}
