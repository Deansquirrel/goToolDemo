package main

import (
	"fmt"
	"github.com/Deansquirrel/goToolCommon"
	log "github.com/Deansquirrel/goToolLog"
	"github.com/Deansquirrel/goToolSecret"
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

	s, err := goToolSecret.EncryptToBase64Format("dsfsef是的法规设法efes是dsfe的法sfghrter规色鬼", "abc")
	if err != nil {
		log.Debug(err.Error())
	} else {
		log.Debug(s)
	}

	r, err := goToolSecret.DecryptFromBase64Format("HDkzKn8uLZjGHEgu/WnfBZ0/hvMfODEuHi80LxgoNy4M7OHwtu+K8IDirfKD46+B1OGsOhk9NSwbPTc=", "abc")
	if err != nil {
		log.Debug(err.Error())
	} else {
		log.Debug(r)
	}
}
