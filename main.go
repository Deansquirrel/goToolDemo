package main

import (
	"fmt"
	"github.com/Deansquirrel/goToolCommon"
	log "github.com/Deansquirrel/goToolLog"
	"github.com/Deansquirrel/goToolRabbitMQ"
	"github.com/Deansquirrel/goToolRedis"
	"time"
)

func main() {
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
	msg := "test message"
	log.Debug(msg)
	log.Info(msg)
	log.Warn(msg)
	log.Error(msg)
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
	redis.Close()
	//============================================================================
	rabbitMQConfig := &goToolRabbitMQ.RabbitMQConfig{
		Server:      "127.0.0.1",
		Port:        5672,
		VirtualHost: "TestHost2",
		User:        "sa",
		Password:    "123456",
	}
	rabbitMQ, err := goToolRabbitMQ.NewRabbitMQ(rabbitMQConfig)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	err = rabbitMQ.QueueDeclareSimple("TestQ")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	err = rabbitMQ.AddProducer("")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	err = rabbitMQ.Publish("", "", "TestQ", "TestQ test ,essage")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	//============================================================================
}
