package main

import (
	"fmt"
	"github.com/Deansquirrel/goServiceSupportHelper"
	"github.com/Deansquirrel/goServiceSupportHelper/global"
	"github.com/Deansquirrel/goToolCommon"
	log "github.com/Deansquirrel/goToolLog"
	"github.com/Deansquirrel/goToolMSSqlHelper"
	"github.com/Deansquirrel/goToolSVRV3"
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

	goServiceSupportHelper.InitParam(&goServiceSupportHelper.Params{
		HttpAddress:   "http://127.0.0.1:8000",
		ClientType:    "demo",
		ClientVersion: global.Version,
		Ctx:           global.Ctx,
		Cancel:        global.Cancel,
	})

	dbConfig, err := goToolSVRV3.GetSQLConfig(
		"118.31.224.35",
		7083,
		"83",
		"8301")
	if err != nil {
		log.Error(err.Error())
		return
	}
	accList, err := goToolSVRV3.GetAccountList(goToolMSSqlHelper.ConvertDbConfigTo2000(dbConfig), "83")
	if accList != nil && len(accList) > 0 {
		dbConfig.DbName = accList[0]
	}
	goServiceSupportHelper.SetOtherInfo(dbConfig, 1, true)
	time.Sleep(time.Minute)
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
