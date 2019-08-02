package main

import (
	"fmt"
	"github.com/Deansquirrel/goToolCommon"
	log "github.com/Deansquirrel/goToolLog"
	"time"
)

func main() {
	log.Level = log.LevelDebug
	log.StdOut = true

	st := time.Now()
	log.Debug(fmt.Sprintf("start time: %s", goToolCommon.GetDateTimeStr(st)))
	defer func() {
		et := time.Now()
		log.Debug(fmt.Sprintf("end time: %s", goToolCommon.GetDateTimeStr(et)))
		log.Debug(fmt.Sprintf("use %fs", et.Sub(st).Seconds()))
	}()
}
