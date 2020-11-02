package main

import (
	"MOAja/payment-service/middleware/log"
	"database/sql"
	"fmt"
	"os"

	"github.com/bagustyo92/wms/modules/inventory/controller"

	"github.com/bagustyo92/wms/modules/inventory/repo"
	"github.com/bagustyo92/wms/modules/inventory/service"

	"github.com/bagustyo92/wms/config"
	"github.com/bagustyo92/wms/modules/inventory/model"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

func main() {
	// INIT APP
	envFilename := os.Getenv("NODE_ENV")
	if envFilename == "" {
		envFilename = "local" // default env is local
	}
	config.InitApp("env/" + envFilename + ".env")

	// Create log file
	if _, err := os.Stat("logs"); os.IsNotExist(err) {
		os.Mkdir("logs", os.ModePerm)
	}

	// DB Connection
	gormClient, err := gorm.Open("mysql", fmt.Sprintf("%s?charset=utf8&parseTime=True&loc=Local", config.DBConnection))
	if err != nil {
		log.MakeLogEntry(nil).Panic(fmt.Sprintf("create gorm client instance failed with message: %s", err.Error()))
		os.Exit(1)
	}

	gormClient.AutoMigrate(
		&model.Product{}, &model.Stock{}, &model.Inbound{}, &model.Outbound{},
	)

	defer gormClient.Close()

	dbClient, err := sql.Open("mysql", fmt.Sprintf("%s?charset=utf8&parseTime=True&loc=Local", config.DBConnection))
	if err != nil {
		log.MakeLogEntry(nil).Panic(fmt.Sprintf("create db client instance failed with message: %s", err.Error()))
		os.Exit(1)
	}
	defer dbClient.Close()

	err = dbClient.Ping()
	if err != nil {
		log.MakeLogEntry(nil).Panic(fmt.Sprintf("db ping failed with message: %s", err.Error()))
		os.Exit(1)
	}

	// Create Redis connection
	// redisConf := &redis.RedisConfig{
	// 	Username:  config.REDIS_USERNAME,
	// 	Password:  config.REDIS_PASSWORD,
	// 	URL:       config.REDIS_URL,
	// 	Port:      config.REDIS_PORT,
	// 	MaxIdle:   10,
	// 	MaxActive: 1000,
	// 	Timeout:   1,
	// 	Wait:      false,
	// }

	// rdb := redis.NewRedisDatabase(redisConf)

	// if err := rdb.Ping(); err != nil {
	// 	log.MakeLogEntry(nil).Panic(err)
	// 	panic(err)
	// }

	// if err := rdb.FlushAll(); err != nil {
	// 	log.MakeLogEntry(nil).Panic(err)
	// 	panic(err)
	// }

	// Declare Echo
	e := echo.New()

	e.Use(log.Logging)
	// e.Use(echoMiddleware.Recover())

	inventoryRepo := repo.NewProductsRepo(gormClient)
	inventoryService := service.NewProductService(inventoryRepo)
	controller.ApplyController(e, inventoryService)

	lock := make(chan error)
	// cron := make(chan error)

	go func(lock chan error) { lock <- e.Start(":" + config.AppPort) }(lock)
	// go func(cron chan error) { cron <- core.ExecuteCronJobPayment(1) }(cron)

	// time.Sleep(1 * time.Millisecond)

	err = <-lock
	if err != nil {
		log.MakeLogEntry(nil).Warning(err)
	}
}
