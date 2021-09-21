package services

import (
	"os"

	"github.com/gohade/hade/framework"
	"github.com/gohade/hade/framework/contract"
)

type HadeConsoleLog struct {
	HadeLog
}

func NewHadeConsoleLog(params ...interface{}) (interface{}, error) {
	level := params[0].(contract.LogLevel)
	ctxFielder := params[1].(contract.CtxFielder)
	formatter := params[2].(contract.Formatter)
	c := params[3].(framework.Container)

	log := &HadeConsoleLog{}

	log.SetLevel(level)
	log.SetCxtFielder(ctxFielder)
	log.SetFormatter(formatter)

	log.SetOutput(os.Stdout)
	log.c = c
	return log, nil
}
