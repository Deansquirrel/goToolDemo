package main

import (
	"context"
	"fmt"
	"github.com/Deansquirrel/goServiceSupportHelper"
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

	ctx, cancel := context.WithCancel(context.Background())
	goServiceSupportHelper.InitParam(&goServiceSupportHelper.Params{
		HttpAddress:   "http://127.0.0.1:8000",
		ClientType:    "DemoType",
		ClientVersion: "0.0.0 Build20000101",
		Ctx:           ctx,
		Cancel:        cancel,
	})

	rFunc := goServiceSupportHelper.NewJob().FormatSSJob("Demo", getWorkerFunc())
	rFunc()

	//err := goServiceSupportHelper.NewJob().JobErrRecord("Demo","abcdefghijklmnopqrstuvwxyz")
	//if err != nil {
	//	log.Error(err.Error())
	//} else {
	//	log.Debug("success")
	//}
}

func getWorkerFunc() func(id string) {
	return func(id string) {
		log.Debug("worker func")
		_ = goServiceSupportHelper.JobErrRecord(id, "abcdefghijklmnopqrstuvwxyz")
	}
}

////func init(){
////	//goServiceSupportHelper.HttpAddress = "http://192.168.8.148:8000"
////	goServiceSupportHelper.InitParam(&goServiceSupportHelper.Params{
////		HttpAddress:"http://192.168.8.148:8000",
////		ClientType:global.Type,
////		ClientVersion:global.Version,
////		DbConfig:&goToolMSSql.MSSqlConfig{
////			Server:"192.168.5.1",
////			Port:2003,
////			User:"sa",
////			Pwd:"",
////			DbName:"Z9门店",
////		},
////		//数据库类型，0-非2000,1-2000
////		DbType:1,
////		IsSvrV3:true,
////		SvrV3AppType:"83",
////		SvrV3ClientType:"8301",
////	})
////	go goServiceSupportHelper.Start()
////}
//
