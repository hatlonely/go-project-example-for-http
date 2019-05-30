package main

import (
	"flag"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/hatlonely/go-http-project-example/internal/gohttp"
	"github.com/hatlonely/go-http-project-example/internal/logger"
	"github.com/spf13/viper"
	"os"
)

func RegisterHandler(r *gin.Engine) {
	r.GET("/hello", gohttp.GoHttpHandler)
}

// AppVersion name
var AppVersion = "unknown"

func main() {
	version := flag.Bool("v", false, "print current version")
	configfile := flag.String("c", "configs/gohttp.json", "config file path")
	flag.Parse()
	if *version {
		fmt.Println(AppVersion)
		os.Exit(0)
	}

	config := viper.New()
	config.SetConfigType("json")
	fp, err := os.Open(*configfile)
	if err != nil {
		panic(err)
	}
	err = config.ReadConfig(fp)
	if err != nil {
		panic(err)
	}

	infoLog, err := logger.NewTextLoggerWithViper(config.Sub("logger.infoLog"))
	if err != nil {
		panic(err)
	}
	warnLog, err := logger.NewTextLoggerWithViper(config.Sub("logger.warnLog"))
	if err != nil {
		panic(err)
	}
	accessLog, err := logger.NewJsonLoggerWithViper(config.Sub("logger.accessLog"))
	gohttp.InfoLog = infoLog
	gohttp.WarnLog = warnLog
	gohttp.AccessLog = accessLog

	infoLog.Infof("%v init success, port[%v]", os.Args[0], config.GetString("service.port"))

	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(cors.Default())

	RegisterHandler(r)

	if err := r.Run(config.GetString("service.port")); err != nil {
		panic(err)
	}
}
