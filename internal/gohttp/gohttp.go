package gohttp

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var InfoLog *logrus.Logger
var WarnLog *logrus.Logger
var AccessLog *logrus.Logger

func init() {
	InfoLog = logrus.New()
	WarnLog = logrus.New()
}

func GoHttpHandler(c *gin.Context) {
	name := c.DefaultQuery("name", "world")

	res := gin.H{
		"message": "hello " + name,
	}

	AccessLog.WithFields(logrus.Fields{
		"request":  fmt.Sprintf("%v%v", c.Request.Host, c.Request.URL),
		"response": res,
	}).Info()

	c.JSON(200, res)
}
