package main

import (
	"fmt"
	"github.com/Deansquirrel/goToolCommon"
	"github.com/Deansquirrel/goToolEnvironment"
	log "github.com/Deansquirrel/goToolLog"
	"github.com/Deansquirrel/goToolMSSql"
	"github.com/Deansquirrel/goToolRedis"
	"github.com/Deansquirrel/goToolSecret"
	"time"
)

func main() {
	log.Level = log.LevelDebug
	log.StdOut = true
	test()
}

func test() {
	currPath, err := goToolCommon.GetCurrPath()
	if err != nil {
		log.Debug(err.Error())
		return
	}
	log.Debug(fmt.Sprintf("currPath：%s", currPath))
	fullPath := currPath + "\\aa\\bb\\cc\\cc\\ee"
	err = goToolCommon.CheckAndCreateFolder(fullPath)
	if err != nil {
		log.Debug(err.Error())
	}
}

func secretTest() {
	//goToolSecret.SetCode("")
	str := "yuansong"
	s, err := goToolSecret.EncryptStr(str)
	if err != nil {
		log.Debug(err.Error())
		return
	}
	log.Debug(s)

	//s = "VEh7RWN2dwFoYngHeF9FUnhcAGdSA2dIYmZjUnhbXkZXYnhCeFwAdVMDY114W15FfGJGWFNkY11SA3xYflh9X116W1xYfVddWngCXFh7U11DBVNfXnh5XF0HdVxdaFZYfXJ7RWd2CENrZgVYflh7aVN0BF9oZ1ZDfwEAZX90ZwdmdgBlf3ZZSGRcfwJpeWRGfHV/Wml5ZEZ8dX9aeF8BDA=="

	r, err := goToolSecret.DecryptStr(s)
	if err != nil {
		log.Debug(err.Error())
		return
	}
	log.Debug(r)
}

func netTest() {
	addr, err := goToolEnvironment.GetIntranetAddr()
	if err != nil {
		log.Debug(err.Error())
	} else {
		log.Debug(addr)
	}
	iAddr, err := goToolEnvironment.GetInternetAddr()
	if err != nil {
		log.Debug(err.Error())
	} else {
		log.Debug(iAddr)
	}

}

func environmentTest() {
	log.Debug(goToolEnvironment.GetOsName())
	ver, err := goToolEnvironment.GetOsVer()
	if err != nil {
		log.Debug(err.Error())
	} else {
		log.Debug(ver)
	}
	hostName, err := goToolEnvironment.GetHostName()
	if err != nil {
		log.Debug(err.Error())
	} else {
		log.Debug(hostName)
	}
}

func sqlTest() {
	config := &goToolMSSql.MSSqlConfig{
		Server: "192.168.5.1",
		Port:   2006,
		DbName: "master",
		User:   "sa",
		Pwd:    "",
	}
	conn, err := goToolMSSql.GetConn(config)
	if err != nil {
		log.Debug(err.Error())
		return
	}
	rows, err := conn.Query("" +
		"SELECT * FROM SysDatabases")
	if err != nil {
		log.Debug(err.Error())
		return
	}
	defer func() {
		_ = rows.Close()
	}()
	columns, err := rows.Columns()
	if err != nil {
		log.Debug(err.Error())
		return
	}
	for _, val := range columns {
		fmt.Println(val)
	}

	//for rows.Next(){
	//	colTypes,err := rows.ColumnTypes()
	//	if err != nil {
	//		log.Debug(err.Error())
	//	} else {
	//		fmt.Println(colTypes)
	//	}
	//
	//}
}

//
//func rabbitMQTest2() {
//	//============================================================================
//	rabbitMQConfig := &goToolRabbitMQ.RabbitMQConfig{
//		Server:      "192.168.8.39",
//		Port:        5672,
//		VirtualHost: "TestHost2",
//		User:        "sa",
//		Password:    "123456",
//	}
//	rabbitMQ, err := goToolRabbitMQ.NewRabbitMQ(rabbitMQConfig)
//	if err != nil {
//		fmt.Println(err.Error())
//		return
//	}
//	err = rabbitMQ.QueueDeclareSimple("TestQ")
//	if err != nil {
//		fmt.Println(err.Error())
//		return
//	}
//
//	errCh := make(chan *goToolRabbitMQ.RabbitMQError)
//	rabbitMQ.NotifyErr(errCh)
//	go func() {
//		for {
//			select {
//			case msg := <-errCh:
//				fmt.Println(msg.Tag)
//				fmt.Println(msg.Type)
//				fmt.Println(msg.Error.Error())
//			}
//		}
//
//	}()
//
//	err = rabbitMQ.AddProducer("")
//	if err != nil {
//		fmt.Println(err.Error())
//		return
//	}
//	go func() {
//		for {
//			msg := "TestQ test message " + goToolCommon.GetDateTimeStr(time.Now())
//			//fmt.Println(msg)
//			err = rabbitMQ.Publish("", "", "TestQ", msg)
//			if err != nil {
//				//fmt.Println(err.Error())
//			}
//			time.Sleep(time.Millisecond * 1000)
//		}
//	}()
//
//	err = rabbitMQ.AddConsumer("", "TestQ", cHandler)
//	if err != nil {
//		fmt.Println(err.Error())
//		return
//	}
//	//============================================================================
//	c := make(chan struct{})
//	<-c
//}

func cHandler(msg string) {
	fmt.Println(goToolCommon.GetDateTimeStr(time.Now()) + " " + msg)
}

func commonTest() {
	//============================================================================
	for i := 0; i < 10; i++ {
		fmt.Println(goToolCommon.RandInt(10, 100))
	}
	guid := goToolCommon.Guid()
	fmt.Println(guid)
	fmt.Println(goToolCommon.Md5(guid))
	//============================================================================
	fmt.Println(goToolCommon.GetDateStr(time.Now()))
	fmt.Println(goToolCommon.GetDateTimeStr(time.Now()))
	fmt.Println(time.Now().Unix())
	fmt.Println(goToolCommon.GetMillisecond(time.Now()))
	fmt.Println(goToolCommon.GetMicrosecond(time.Now()))
	fmt.Println(time.Now().UnixNano())
	//============================================================================
	fmt.Println(goToolCommon.GetCurrPath())
	_ = goToolCommon.Log("Test Message")
	_ = goToolCommon.LogFile("Test Message", "logFile")
	//============================================================================
}

func logTest() {
	//============================================================================
	msg := "test message"
	log.Debug(msg)
	log.Info(msg)
	log.Warn(msg)
	log.Error(msg)
	//============================================================================
}

func redisTest() {
	//============================================================================
	redisConfig := &goToolRedis.RedisConfig{
		Server:      "127.0.0.1",
		Port:        6379,
		Auth:        "12345",
		MaxIdle:     5000,
		MaxActive:   5000,
		IdleTimeout: 60,
	}
	redis := goToolRedis.NewRedis(redisConfig)
	fmt.Println(redis.Set(0, "testKey", "testValue"))
	fmt.Println(redis.IsExists(0, "testKey"))
	fmt.Println(redis.IsExists(0, "testKeyN"))
	fmt.Println(redis.Get(0, "testKey"))
	fmt.Println(redis.Get(0, "testKeyN"))
	fmt.Println(redis.Del(0, "testKey"))
	fmt.Println(redis.Del(0, "testKeyM"))
	//============================================================================

	redis.Close()
	//============================================================================
}

//
//func rabbitMQTest() {
//	//============================================================================
//	rabbitMQConfig := &goToolRabbitMQ.RabbitMQConfig{
//		Server:      "127.0.0.1",
//		Port:        5672,
//		VirtualHost: "TestHost2",
//		User:        "sa",
//		Password:    "123456",
//	}
//	rabbitMQ, err := goToolRabbitMQ.NewRabbitMQ(rabbitMQConfig)
//	if err != nil {
//		fmt.Println(err.Error())
//		return
//	}
//	err = rabbitMQ.QueueDeclareSimple("TestQ")
//	if err != nil {
//		fmt.Println(err.Error())
//		return
//	}
//	err = rabbitMQ.AddProducer("")
//	if err != nil {
//		fmt.Println(err.Error())
//		return
//	}
//	err = rabbitMQ.Publish("", "", "TestQ", "TestQ test ,essage")
//	if err != nil {
//		fmt.Println(err.Error())
//		return
//	}
//	//============================================================================
//}
