package main

import (
	"fmt"
	"github.com/Deansquirrel/goToolCommon"
	log "github.com/Deansquirrel/goToolLog"
	"github.com/Deansquirrel/goToolMSSql"
	"time"
)

const (
	sqlInsert = "" +
		"INSERT INTO [BeatHeart]([task%d]) " +
		"VALUES (getDate())"
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

	//s, err := goToolSecret.EncryptToBase64Format("dsfsef是的法规设法efes是dsfe的法sfghrter规色鬼", "abc")
	//if err != nil {
	//	log.Debug(err.Error())
	//} else {
	//	log.Debug(s)
	//}
	//
	//r, err := goToolSecret.DecryptFromBase64Format("HDkzKn8uLZjGHEgu/WnfBZ0/hvMfODEuHi80LxgoNy4M7OHwtu+K8IDirfKD46+B1OGsOhk9NSwbPTc=", "abc")
	//if err != nil {
	//	log.Debug(err.Error())
	//} else {
	//	log.Debug(r)
	//}

	dbConfig := goToolMSSql.MSSqlConfig{
		Server: "192.168.10.166",
		Port:   2433,
		DbName: "SmpTest",
		User:   "sa",
		Pwd:    "",
	}

	//closeChan := make(chan error)

	//0-19 * * * * ?
	//10-29 * * * * ?
	//20-39 * * * * ?
	//30-49 * * * * ?
	//40-59 * * * * ?
	//0-9,50-59 * * * * ?
	taskIndex := 1
	//cronStr := "0-19 * * * * ?"

	log.Debug("TEST")
	conn, err := goToolMSSql.GetConn(&dbConfig)
	if err != nil {
		log.Error(err.Error())
		return
	}
	_, err = conn.Exec(fmt.Sprintf(sqlInsert, taskIndex))
	if err != nil {
		log.Error(err.Error())
		return
	}

	//c := cron.New()
	//err := c.AddFunc(cronStr, func() {
	//	log.Debug("TEST")
	//	conn, err := goToolMSSql.GetConn(&dbConfig)
	//	if err != nil {
	//		log.Error(err.Error())
	//		return
	//	}
	//	_, err = conn.Exec(fmt.Sprintf(sqlInsert, taskIndex))
	//	if err != nil {
	//		log.Error(err.Error())
	//		return
	//	}
	//})
	//if err != nil {
	//	closeChan <- err
	//} else {
	//	c.Start()
	//}
	//
	//select {
	//case e := <-closeChan:
	//	log.Error(e.Error())
	//	return
	//}
}
