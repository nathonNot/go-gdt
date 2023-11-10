package log

import (
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"time"
)
import "github.com/panjf2000/ants/v2"

func init() {
	InitLog()
}

func InitLog() {
	logrus.SetReportCaller(true)
	logrus.SetFormatter(&logrus.JSONFormatter{
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyFile: "file_name",
		},
	})
	go func() {
		os.Mkdir("log", os.ModePerm)
		file, _ := os.OpenFile("log/logrus.log", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
		logrus.SetOutput(os.Stdout)
		//设置output,默认为stderr,可以为任何io.Writer，比如文件*os.File
		writers := []io.Writer{
			file,
			os.Stdout}
		//同时写文件和屏幕
		fileAndStdoutWriter := io.MultiWriter(writers...)
		logrus.SetOutput(fileAndStdoutWriter)
	}()
	Logger = &Wrapper{logrus.WithFields(logrus.Fields{})}
	ants.WithLogger(Logger)
	ants.WithExpiryDuration(time.Second * 10)
}
