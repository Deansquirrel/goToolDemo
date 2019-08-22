package main

import (
	"context"
	"fmt"
	"github.com/Deansquirrel/goToolCommon"
	"github.com/Deansquirrel/goToolDemo/common"
	"github.com/Deansquirrel/goToolDemo/global"
	"github.com/Deansquirrel/goToolDemo/object"
	log "github.com/Deansquirrel/goToolLog"
	"github.com/Deansquirrel/goToolMSSql"
	"github.com/Deansquirrel/goToolMSSqlHelper"
	"time"
)

const (
	sqlCrateLinkServer = "" +
		"declare @rmtsvr sysname " +
		"declare @rmtsvrsapassword sysname " +
		"set @rmtsvr = ? " +
		"set @rmtsvrsapassword  = ? " +
		"exec sp_addlinkedserver '%s','','sqloledb',@rmtsvr " +
		"exec sp_addlinkedsrvlogin '%s','false',null,'sa',@rmtsvrsapassword "
	sqlDropLinkServer = "" +
		"exec sp_dropserver '%s','droplogins'"
)

var prList []string

//初始化
func init() {
	global.Args = &object.ProgramArgs{}
	global.SysConfig = &object.SystemConfig{}

	global.Ctx, global.Cancel = context.WithCancel(context.Background())

	prList = make([]string, 0)
	prList = append(prList, "pr_imp_repcgrkdt")
	prList = append(prList, "pr_imp_rephpckdjdt")
	prList = append(prList, "pr_imp_rephprkdjdt")
	prList = append(prList, "pr_imp_repmdpsckdt")
	prList = append(prList, "pr_imp_repmdxshprhz")
	prList = append(prList, "pr_imp_reppsshdt")
	prList = append(prList, "pr_imp_reppstzdt")
	prList = append(prList, "pr_imp_reppstzqrdt")
	prList = append(prList, "pr_imp_reppsxzdt")
	prList = append(prList, "pr_imp_reppsxzshdt")
	prList = append(prList, "pr_imp_repthdjdt")
	prList = append(prList, "pr_imp_repthshqrdt")
	prList = append(prList, "pr_imp_repthxzdt")
	prList = append(prList, "pr_imp_repthxzshdt")
	prList = append(prList, "pr_imp_repykdt")
	prList = append(prList, "pr_imp_repzbwgrkdt")
	prList = append(prList, "pr_imp_repzbwgrkylt")
	prList = append(prList, "pr_imp_repmddbckdt")
	prList = append(prList, "pr_imp_repdrshdt")
	prList = append(prList, "pr_imp_repdbtzdt")
	prList = append(prList, "pr_imp_repdbtzshdt")
	prList = append(prList, "pr_imp_repddt")
	prList = append(prList, "pr_imp_repdddt")
	prList = append(prList, "pr_imp_repxsddthdt")
	prList = append(prList, "pr_imp_reptcachz")
	prList = append(prList, "pr_imp_repmdxssrhzt")
	prList = append(prList, "pr_imp_repmdxssrzzdt")
	prList = append(prList, "pr_imp_repzzhz")
	prList = append(prList, "pr_imp_repslqhz")
	prList = append(prList, "pr_imp_reppsdbhz")
	prList = append(prList, "pr_imp_repstockchange")
	prList = append(prList, "pr_imp_repicxkcz")
	prList = append(prList, "pr_imp_repzshphz")
	prList = append(prList, "pr_imp_repwgrkylhz")
	prList = append(prList, "pr_imp_repwgrkrhz")
	prList = append(prList, "pr_imp_repszxsckdt")
	prList = append(prList, "pr_imp_repsczhdj")
	prList = append(prList, "pr_imp_reppkdjdt")
	prList = append(prList, "pr_imp_repkctzdt")
	prList = append(prList, "pr_imp_repiccz")
	prList = append(prList, "pr_imp_repddtdt")
	prList = append(prList, "pr_imp_repddtdhpmxt")
	prList = append(prList, "pr_imp_repiczk")
	prList = append(prList, "pr_imp_repick")
	prList = append(prList, "pr_imp_repwgwhhz")
	prList = append(prList, "pr_imp_repsctl")
	prList = append(prList, "pr_imp_repscrkmx")
}

func main() {
	fmt.Println(global.Version)
	log.Info(global.Version)

	//解析命令行参数
	{
		global.Args.Definition()
		global.Args.Parse()
		err := global.Args.Check()
		if err != nil {
			fmt.Print(err.Error())
			log.Error(err.Error())
			return
		}
		common.UpdateParams()
	}
	//加载系统参数
	{
		common.LoadSysConfig()
		common.RefreshSysConfig()
	}

	linkServer, err := createLinkServer(global.SysConfig.Total.RmtSvr, global.SysConfig.Total.RmtSvrSaPassWord)
	if err != nil {
		return
	}
	defer func() {
		err = dropLinkServer(linkServer)
	}()

	begDate, err := time.Parse("2006-01-02", global.SysConfig.Total.BegDate)
	if err != nil {
		log.Error(fmt.Sprintf("fmt begdate err: %s", err.Error()))
		return
	}
	endDate, err := time.Parse("2006-01-02", global.SysConfig.Total.EndDate)
	if err != nil {
		log.Error(fmt.Sprintf("fmt enddate err: %s", err.Error()))
		return
	}

	tBegDate := begDate
	d := time.Hour * 24 * 7
	addD := time.Hour * 24 * 8
	for {
		if goToolCommon.GetDateStr(tBegDate.Add(d)) < goToolCommon.GetDateStr(endDate) {
			getData(goToolCommon.GetDateStr(tBegDate), goToolCommon.GetDateStr(tBegDate.Add(d)), global.SysConfig.Total.YwDbName, linkServer)
			tBegDate = tBegDate.Add(addD)
		} else {
			getData(goToolCommon.GetDateStr(tBegDate), goToolCommon.GetDateStr(endDate), global.SysConfig.Total.YwDbName, linkServer)
			break
		}
	}
	log.Info("Complete")
}

func getData(begDate string, endDate string, ywDbName string, linkServer string) {
	for _, pr := range prList {
		err := getDataByPr(begDate, endDate, ywDbName, linkServer, pr)
		if err != nil {
			return
		}
	}
}

const (
	sqlPr = "exec %s '%s','%s','%s','%s'"
)

func getDataByPr(begDate string, endDate string, ywDbName string, linkServer string, pr string) error {
	log.Debug(fmt.Sprintf("%s %s %s", pr, begDate, endDate))
	dbConfig := goToolMSSqlHelper.ConvertDbConfigTo2000(getLocalDbConfig())
	err := goToolMSSqlHelper.SetRowsBySQL2000(dbConfig, fmt.Sprintf(sqlPr, pr, begDate, endDate, ywDbName, linkServer))
	if err != nil {
		log.Error(fmt.Sprintf("%s get data[%s][%s] err:%s", pr, begDate, endDate, err.Error()))
		return err
	}
	return nil
}

//获取本地库连接配置
func getLocalDbConfig() *goToolMSSql.MSSqlConfig {
	return &goToolMSSql.MSSqlConfig{
		Server: global.SysConfig.LocalDb.Server,
		Port:   global.SysConfig.LocalDb.Port,
		DbName: global.SysConfig.LocalDb.DbName,
		User:   global.SysConfig.LocalDb.User,
		Pwd:    global.SysConfig.LocalDb.Pwd,
	}
}

func createLinkServer(rmtSvr string, rmtSvrSaPwd string) (string, error) {
	uuid := goToolCommon.Guid()
	sqlStr := fmt.Sprintf(sqlCrateLinkServer, uuid, uuid)
	err := goToolMSSqlHelper.SetRowsBySQL2000(
		goToolMSSqlHelper.ConvertDbConfigTo2000(getLocalDbConfig()),
		sqlStr,
		rmtSvr,
		rmtSvrSaPwd)
	if err != nil {
		log.Error(fmt.Sprintf("create link server err: %s", err.Error()))
		return "", err
	}
	return uuid, err
}

func dropLinkServer(id string) error {
	sqlStr := fmt.Sprintf(sqlDropLinkServer, id)
	err := goToolMSSqlHelper.SetRowsBySQL2000(
		goToolMSSqlHelper.ConvertDbConfigTo2000(getLocalDbConfig()),
		sqlStr)
	if err != nil {
		log.Error(fmt.Sprintf("drop link server err: %s", err.Error()))
		return err
	}
	return nil
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
