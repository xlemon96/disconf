package util

import (
	"fmt"
	"log"
	"os"
)

var G_log *myLogger

func InitLogger() error {
	file, err := os.OpenFile(G_conf.LogConfig.LogPath, os.O_RDWR|os.O_APPEND, 777)
	if err != nil {
		return err
	}
	logger := log.New(file, "[gogogo]", log.LstdFlags|log.Lshortfile)
	G_log = &myLogger{logger: logger}
	return nil
}

type myLogger struct {
	logger *log.Logger
}

func (u *myLogger) Info(format string, data ...interface{}) {
	output := fmt.Sprintf(format, data...)
	u.logger.Printf("[info] %s", output)
}

func (u *myLogger) Error(format string, data ...interface{}) {
	output := fmt.Sprintf(format, data...)
	u.logger.Printf("[error] %s", output)
}

func (u *myLogger) Warn(format string, data ...interface{}) {
	output := fmt.Sprintf(format, data...)
	u.logger.Printf("[warn] %s", output)
}