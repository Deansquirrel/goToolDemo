package main

import (
	"fmt"
	"github.com/Deansquirrel/goToolCommon"
	"github.com/Deansquirrel/goToolEnvironment"
	log "github.com/Deansquirrel/goToolLog"
	"github.com/Deansquirrel/goToolMSSql"
	"strings"
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
	printIP()

	sqlTest()

	//test()
	//getData()

}

func printIP() {
	ip, err := goToolEnvironment.GetInternetAddr()
	if err != nil {
		log.Error(fmt.Sprintf("get ip error: %s", err.Error()))
		return
	}
	log.Debug(fmt.Sprintf("IP: %s", ip))
}

func sqlTest() {
	config := goToolMSSql.MSSqlConfig{
		"192.168.10.166",
		2433,
		"master",
		"sa",
		"",
	}

	conn, err := goToolMSSql.GetConn(&config)
	if err != nil {
		log.Error(fmt.Sprintf("get conn error: %s", err.Error()))
		return
	}
	tx, err := conn.Begin()
	if err != nil {
		log.Error(fmt.Sprintf("get tx error: %s", err.Error()))
		return
	}

	var txErr error
	defer func() {
		if txErr != nil {
			err = tx.Rollback()
			if err != nil {
				log.Error(fmt.Sprintf("rollback error: %s", err.Error()))
			}
		} else {
			err = tx.Commit()
			if err != nil {
				log.Error(fmt.Sprintf("commit error: %s", err.Error()))
			}
		}
	}()

	_, txErr = tx.Exec(getCreateSql())
	if txErr != nil {
		log.Error(fmt.Sprintf("create error: %s", txErr.Error()))
		return
	}

	_, txErr = tx.Exec(getInsertSql(), "001", "A", "M", "1900-01-01", "11")
	if txErr != nil {
		log.Error(fmt.Sprintf("insert error: %s", txErr.Error()))
		return
	}

	rows, txErr := tx.Query("select * from #Student")
	if txErr != nil {
		log.Error(fmt.Sprintf("search error: %s", txErr.Error()))
		return
	}
	defer func() {
		_ = rows.Close()
	}()

	var sno, sname, ssex, sclass string
	var b time.Time
	for rows.Next() {
		err = rows.Scan(&sno, &sname, &ssex, &b, &sclass)
		if err != nil {
			log.Error(fmt.Sprintf("scan error: %s", err.Error()))
			return
		}
		log.Debug(fmt.Sprintf("%s %s %s %s %s", sno, sname, ssex, goToolCommon.GetDateTimeStr(b), sclass))
	}
}

func getCreateSql() string {
	str := strings.Builder{}
	str.WriteString("if object_id('tempdb..#Student') is not null ")
	str.WriteString("begin ")
	str.WriteString("	drop table #Student ")
	str.WriteString("end ")

	str.WriteString("CREATE TABLE #Student( ")
	str.WriteString("	[Sno] [varchar](10) NOT NULL, ")
	str.WriteString("	[Sname] [varchar](20) NULL, ")
	str.WriteString("	[Ssex] [varchar](2) NULL, ")
	str.WriteString("	[Sbirthday] [datetime] NULL, ")
	str.WriteString("	[class] [varchar](20) NULL) ")
	return str.String()
}

func getInsertSql() string {
	str := strings.Builder{}
	str.WriteString("INSERT INTO #Student([Sno],[Sname],[Ssex],[Sbirthday],[class]) ")
	str.WriteString("VALUES (?,?,?,?,?)")
	return str.String()
}

//func getData() {
//	currPath, err := goToolCommon.GetCurrPath()
//	if err != nil {
//		log.Error(fmt.Sprintf("get curr path error: %s", err.Error()))
//		return
//	}
//	fullPath := currPath + goToolCommon.GetFolderSplitStr() + "txt"
//	_, fileList, err := goToolCommon.GetFolderAndFileList(fullPath)
//	if err != nil {
//		log.Error(fmt.Sprintf("get file list error: %s", err.Error()))
//		return
//	}
//	result := make([]string, 0)
//	for _, f := range fileList {
//		fi, err := os.Open(fullPath + "\\" + f)
//		if err != nil {
//			log.Error(fmt.Sprintf("open file error: %s", err.Error()))
//			return
//		}
//
//		strPre := "备注描述："
//
//		br := bufio.NewReader(fi)
//		for {
//			data, _, err := br.ReadLine()
//			if err == io.EOF {
//				break
//			}
//			str := string(data)
//			str = strings.Trim(str, "")
//
//			if strings.HasPrefix(str, strPre) {
//				str = strings.Replace(str, strPre, "", -1)
//				result = append(result, str)
//			}
//		}
//		_ = fi.Close()
//	}
//	fileName := goToolCommon.GetDateTimeStr(time.Now()) + ".txt"
//	fileName = strings.Replace(fileName, " ", "", -1)
//	fileName = strings.Replace(fileName, "-", "", -1)
//	fileName = strings.Replace(fileName, ":", "", -1)
//	for _, s := range result {
//		err := goToolCommon.LogFile(s, fileName)
//		if err != nil {
//			log.Error(fmt.Sprintf("save data error: %s data: %s", err.Error(), s))
//		}
//	}
//}
//
//func test() {
//	currPath, err := goToolCommon.GetCurrPath()
//	if err != nil {
//		log.Debug(err.Error())
//		return
//	}
//	log.Debug(fmt.Sprintf("currPath：%s", currPath))
//	fullPath := currPath + "\\aa\\bb\\cc\\cc\\ee"
//	err = goToolCommon.CheckAndCreateFolder(fullPath)
//	if err != nil {
//		log.Debug(err.Error())
//	}
//}
//
//func secretTest() {
//	//goToolSecret.SetCode("")
//	str := "yuansong"
//	s, err := goToolSecret.EncryptStr(str)
//	if err != nil {
//		log.Debug(err.Error())
//		return
//	}
//	log.Debug(s)
//
//	//s = "VEh7RWN2dwFoYngHeF9FUnhcAGdSA2dIYmZjUnhbXkZXYnhCeFwAdVMDY114W15FfGJGWFNkY11SA3xYflh9X116W1xYfVddWngCXFh7U11DBVNfXnh5XF0HdVxdaFZYfXJ7RWd2CENrZgVYflh7aVN0BF9oZ1ZDfwEAZX90ZwdmdgBlf3ZZSGRcfwJpeWRGfHV/Wml5ZEZ8dX9aeF8BDA=="
//
//	r, err := goToolSecret.DecryptStr(s)
//	if err != nil {
//		log.Debug(err.Error())
//		return
//	}
//	log.Debug(r)
//}
//
//func netTest() {
//	addr, err := goToolEnvironment.GetIntranetAddr()
//	if err != nil {
//		log.Debug(err.Error())
//	} else {
//		log.Debug(addr)
//	}
//	iAddr, err := goToolEnvironment.GetInternetAddr()
//	if err != nil {
//		log.Debug(err.Error())
//	} else {
//		log.Debug(iAddr)
//	}
//
//}
//
//func environmentTest() {
//	log.Debug(goToolEnvironment.GetOsName())
//	ver, err := goToolEnvironment.GetOsVer()
//	if err != nil {
//		log.Debug(err.Error())
//	} else {
//		log.Debug(ver)
//	}
//	hostName, err := goToolEnvironment.GetHostName()
//	if err != nil {
//		log.Debug(err.Error())
//	} else {
//		log.Debug(hostName)
//	}
//}
//
//func sqlTest() {
//	config := &goToolMSSql.MSSqlConfig{
//		Server: "192.168.5.1",
//		Port:   2006,
//		DbName: "master",
//		User:   "sa",
//		Pwd:    "",
//	}
//	conn, err := goToolMSSql.GetConn(config)
//	if err != nil {
//		log.Debug(err.Error())
//		return
//	}
//	rows, err := conn.Query("" +
//		"SELECT * FROM SysDatabases")
//	if err != nil {
//		log.Debug(err.Error())
//		return
//	}
//	defer func() {
//		_ = rows.Close()
//	}()
//	columns, err := rows.Columns()
//	if err != nil {
//		log.Debug(err.Error())
//		return
//	}
//	for _, val := range columns {
//		fmt.Println(val)
//	}
//
//	//for rows.Next(){
//	//	colTypes,err := rows.ColumnTypes()
//	//	if err != nil {
//	//		log.Debug(err.Error())
//	//	} else {
//	//		fmt.Println(colTypes)
//	//	}
//	//
//	//}
//}
//
////
////func rabbitMQTest2() {
////	//============================================================================
////	rabbitMQConfig := &goToolRabbitMQ.RabbitMQConfig{
////		Server:      "192.168.8.39",
////		Port:        5672,
////		VirtualHost: "TestHost2",
////		User:        "sa",
////		Password:    "123456",
////	}
////	rabbitMQ, err := goToolRabbitMQ.NewRabbitMQ(rabbitMQConfig)
////	if err != nil {
////		fmt.Println(err.Error())
////		return
////	}
////	err = rabbitMQ.QueueDeclareSimple("TestQ")
////	if err != nil {
////		fmt.Println(err.Error())
////		return
////	}
////
////	errCh := make(chan *goToolRabbitMQ.RabbitMQError)
////	rabbitMQ.NotifyErr(errCh)
////	go func() {
////		for {
////			select {
////			case msg := <-errCh:
////				fmt.Println(msg.Tag)
////				fmt.Println(msg.Type)
////				fmt.Println(msg.Error.Error())
////			}
////		}
////
////	}()
////
////	err = rabbitMQ.AddProducer("")
////	if err != nil {
////		fmt.Println(err.Error())
////		return
////	}
////	go func() {
////		for {
////			msg := "TestQ test message " + goToolCommon.GetDateTimeStr(time.Now())
////			//fmt.Println(msg)
////			err = rabbitMQ.Publish("", "", "TestQ", msg)
////			if err != nil {
////				//fmt.Println(err.Error())
////			}
////			time.Sleep(time.Millisecond * 1000)
////		}
////	}()
////
////	err = rabbitMQ.AddConsumer("", "TestQ", cHandler)
////	if err != nil {
////		fmt.Println(err.Error())
////		return
////	}
////	//============================================================================
////	c := make(chan struct{})
////	<-c
////}
//
//func cHandler(msg string) {
//	fmt.Println(goToolCommon.GetDateTimeStr(time.Now()) + " " + msg)
//}
//
//func commonTest() {
//	//============================================================================
//	for i := 0; i < 10; i++ {
//		fmt.Println(goToolCommon.RandInt(10, 100))
//	}
//	guid := goToolCommon.Guid()
//	fmt.Println(guid)
//	fmt.Println(goToolCommon.Md5([]byte(guid)))
//	//============================================================================
//	fmt.Println(goToolCommon.GetDateStr(time.Now()))
//	fmt.Println(goToolCommon.GetDateTimeStr(time.Now()))
//	fmt.Println(time.Now().Unix())
//	fmt.Println(goToolCommon.GetMillisecond(time.Now()))
//	fmt.Println(goToolCommon.GetMicrosecond(time.Now()))
//	fmt.Println(time.Now().UnixNano())
//	//============================================================================
//	fmt.Println(goToolCommon.GetCurrPath())
//	_ = goToolCommon.Log("Test Message")
//	_ = goToolCommon.LogFile("Test Message", "logFile")
//	//============================================================================
//}
//
//func logTest() {
//	//============================================================================
//	msg := "test message"
//	log.Debug(msg)
//	log.Info(msg)
//	log.Warn(msg)
//	log.Error(msg)
//	//============================================================================
//}
//
//func redisTest() {
//	//============================================================================
//	redisConfig := &goToolRedis.RedisConfig{
//		Server:      "127.0.0.1",
//		Port:        6379,
//		Auth:        "12345",
//		MaxIdle:     5000,
//		MaxActive:   5000,
//		IdleTimeout: 60,
//	}
//	redis := goToolRedis.NewRedis(redisConfig)
//	fmt.Println(redis.Set(0, "testKey", "testValue"))
//	fmt.Println(redis.IsExists(0, "testKey"))
//	fmt.Println(redis.IsExists(0, "testKeyN"))
//	fmt.Println(redis.Get(0, "testKey"))
//	fmt.Println(redis.Get(0, "testKeyN"))
//	fmt.Println(redis.Del(0, "testKey"))
//	fmt.Println(redis.Del(0, "testKeyM"))
//	//============================================================================
//
//	redis.Close()
//	//============================================================================
//}
//
////
////func rabbitMQTest() {
////	//============================================================================
////	rabbitMQConfig := &goToolRabbitMQ.RabbitMQConfig{
////		Server:      "127.0.0.1",
////		Port:        5672,
////		VirtualHost: "TestHost2",
////		User:        "sa",
////		Password:    "123456",
////	}
////	rabbitMQ, err := goToolRabbitMQ.NewRabbitMQ(rabbitMQConfig)
////	if err != nil {
////		fmt.Println(err.Error())
////		return
////	}
////	err = rabbitMQ.QueueDeclareSimple("TestQ")
////	if err != nil {
////		fmt.Println(err.Error())
////		return
////	}
////	err = rabbitMQ.AddProducer("")
////	if err != nil {
////		fmt.Println(err.Error())
////		return
////	}
////	err = rabbitMQ.Publish("", "", "TestQ", "TestQ test ,essage")
////	if err != nil {
////		fmt.Println(err.Error())
////		return
////	}
////	//============================================================================
////}
