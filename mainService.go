package main

import (
	"fmt"
	"github.com/Deansquirrel/goToolCommon"
	"github.com/Deansquirrel/goToolDemo/object"
	"github.com/kardianos/service"
	"os/exec"
)

import log "github.com/Deansquirrel/goToolLog"

//初始化
func init() {

}

func main() {
	log.Level = log.LevelDebug
	log.StdOut = true
	//安装、卸载或运行程序
	{
		svcConfig := &service.Config{
			Name:        "DemS",
			DisplayName: "DemS",
			Description: "DemS",
		}
		prg := &program{}
		s, err := service.New(prg, svcConfig)
		if err != nil {
			log.Error("定义服务配置时遇到错误：" + err.Error())
			return
		}

		pArgs := object.ProgramArgs{}
		pArgs.Definition()
		pArgs.Parse()
		if err := pArgs.Check(); err != nil {
			log.Error(err.Error())
			return
		}

		if pArgs.IsInstall {
			err = s.Install()
			if err != nil {
				log.Error("安装为服务时遇到错误：" + err.Error())
			} else {
				fmt.Println(fmt.Sprintf("服务安装成功"))
			}
			return
		}
		if pArgs.IsUninstall {
			err = s.Uninstall()
			if err != nil {
				log.Error("卸载服务时遇到错误：" + err.Error())
			} else {
				fmt.Println(fmt.Sprintf("服务卸载成功"))
			}
			return
		}
		_ = s.Run()
	}
}

type program struct{}

func (p *program) Start(s service.Service) error {
	err := p.run()
	if err != nil {
		log.Error(fmt.Sprintf("服务启动时遇到错误：%s", err.Error()))
	}
	return err
}

func (p *program) run() error {
	//服务所执行的代码
	log.Warn("Service Starting")
	defer log.Warn("Service Started")
	{
		//go func(){
		//	c := cron.New()
		//	err := c.AddFunc("* * * * * ?",func(){
		//		log.Debug(goToolCommon.GetDateTimeStr(time.Now()))
		//	})
		//	if err != nil {
		//		log.Error(err.Error())
		//		return
		//	}
		//	c.Start()
		//}()
		p, err := goToolCommon.GetCurrPath()
		if err != nil {
			log.Error(err.Error())
			return nil
		}

		cmd := exec.Command(p + goToolCommon.GetFolderSplitStr() + "start.exe")
		_, err = cmd.CombinedOutput()
		if err != nil {
			log.Error(fmt.Sprintf("call start err: %s", err.Error()))
		} else {
			log.Debug("call start success")
		}
	}
	return nil
}

func (p *program) Stop(s service.Service) error {
	log.Warn("Service Stopping")
	defer log.Warn("Service Stopped")
	{
	}
	return nil
}
