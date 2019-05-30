module github.com/hatlonely/go-project-example-for-http

replace (
	golang.org/x/sys v0.0.0 => github.com/golang/sys v0.0.0
	golang.org/x/text v0.3.2 => github.com/golang/text v0.3.2
)

require (
	github.com/gin-contrib/cors v1.3.0
	github.com/gin-gonic/gin v1.4.0
	github.com/hatlonely/go-http-project-example v0.0.0-20190530065635-f4f73df2d04a
	github.com/lestrrat-go/file-rotatelogs v2.2.0+incompatible
	github.com/sirupsen/logrus v1.4.2
	github.com/smartystreets/goconvey v0.0.0-20190330032615-68dc04aab96a
	github.com/spf13/viper v1.4.0
)
