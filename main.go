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

//初始化
func init() {
	global.Args = &object.ProgramArgs{}
	global.SysConfig = &object.SystemConfig{}

	global.Ctx, global.Cancel = context.WithCancel(context.Background())
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

	lastDate, err := getLastOpenDate()
	if err != nil {
		log.Error(err.Error())
		return
	}

	fDate, lDate := getDateRange(lastDate)

	for currDate := fDate; goToolCommon.GetDateStr(currDate) <= goToolCommon.GetDateStr(lDate); currDate = currDate.AddDate(0, 0, 1) {
		idList, err := getMdIdByOpenDate(currDate.AddDate(0, 0, -7))
		if err != nil {
			log.Error(err.Error())
			return
		}
		for _, id := range idList {
			err = addOpenDate(currDate, id)
			if err != nil {
				log.Error(err.Error())
				return
			}
		}
	}
	log.Info("Complete")
}

//获取最后一个有效营业日
const (
	sqlGetLastOpenDate = "" +
		"select top 1 [date] " +
		"from opendate " +
		"order by [date] desc"
)

func getLastOpenDate() (time.Time, error) {
	var t time.Time
	var flag bool
	rows, err := goToolMSSqlHelper.GetRowsBySQL(getLocalDbConfig(), sqlGetLastOpenDate)
	if err != nil {
		return time.Now(), err
	}
	defer func() {
		_ = rows.Close()
	}()
	for rows.Next() {
		err = rows.Scan(&t)
		if err != nil {
			return time.Now(), nil
		}
		flag = true
	}
	if rows.Err() != nil {
		return time.Now(), rows.Err()
	}
	if flag {
		return t, nil
	} else {
		return time.Now().AddDate(0, -1, 0), nil
	}
}

func getDateRange(t time.Time) (time.Time, time.Time) {
	fDate := t.AddDate(0, 1, -t.Day()+1)
	lDate := fDate.AddDate(0, 1, -1)
	return fDate, lDate
}

const sqlGetMdIdByOpenDate = "" +
	"select mdid " +
	"from opendate " +
	"where [date] = ?"

func getMdIdByOpenDate(t time.Time) ([]int, error) {
	rows, err := goToolMSSqlHelper.GetRowsBySQL(getLocalDbConfig(), sqlGetMdIdByOpenDate, goToolCommon.GetDateStr(t))
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = rows.Close()
	}()
	rList := make([]int, 0)
	for rows.Next() {
		var id int
		err = rows.Scan(&id)
		if err != nil {
			return nil, err
		}
		rList = append(rList, id)
	}
	if rows.Err() != nil {
		return nil, rows.Err()
	}
	return rList, nil
}

const sqlAddOpenDate = "" +
	"INSERT INTO [opendate]([mdid],[date],[yystate],[lastupdate]) " +
	"VALUES (?,?,?,?)"

func addOpenDate(t time.Time, mdId int) error {
	log.Debug(fmt.Sprintf("%d %s", mdId, goToolCommon.GetDateStr(t)))
	return goToolMSSqlHelper.SetRowsBySQL(getLocalDbConfig(), sqlAddOpenDate, mdId, goToolCommon.GetDateStr(t), 0, time.Now())
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
