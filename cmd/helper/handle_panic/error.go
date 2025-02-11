package handle_panic

import (
	"autentikasi1/configs/configs_logrus"
	"fmt"
	"github.com/sirupsen/logrus"
	"log"
	"runtime"
)

var (
	Logging = configs_logrus.SetupLogger()
)

func PanicIfErr(message string, err error) {
	if err != nil {
		panic(fmt.Sprintf("PANIC (%s::%s)", message, err.Error()))
	}
}

func DeferFunc(logLevel logrus.Level, message string, err ...func(errorMsg any)) {
	if r := recover(); r != nil {
		buf := make([]byte, 1024)
		runtime.Stack(buf, false)
		log.Printf("\n\nFROM DEFER:\nLogging: %v\nPesan Custom: %s\n\n\n%s", r, message, buf)

		if logLevel == logrus.PanicLevel || logLevel == logrus.FatalLevel {
			Logging.Fatal(r)
		} else {
			Logging.Log(logLevel, r)
		}

		if len(err) > 0 {
			err[0](r)
		}
	}
}
