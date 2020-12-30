package logging

import (
	"os"
	"path/filepath"
	"squirtle/src/common/util"
	"strconv"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	log "github.com/sirupsen/logrus"
)

func GetLogLevel(s string) log.Level {
	switch s {
	case "Debug":
		return log.DebugLevel
	case "Info":
		return log.InfoLevel
	case "Warn":
		return log.WarnLevel
	case "Error":
		return log.ErrorLevel
	case "Fatal":
		return log.FatalLevel
	case "Panic":
		return log.PanicLevel
	default:
		return log.InfoLevel
	}
}

func newLfsHook(logName string, level log.Level, maxRemain uint) log.Hook {
	writer, err := rotatelogs.New(
		logName+"."+strconv.Itoa(os.Getpid())+".%Y%m%d%H",
		// WithLinkName为最新的日志建立软连接，以方便随着找到当前日志文件
		rotatelogs.WithLinkName(logName+"."+strconv.Itoa(os.Getpid())),

		// WithRotationTime设置日志分割的时间，这里设置为一小时分割一次
		rotatelogs.WithRotationTime(time.Hour*1),

		// WithMaxAge和WithRotationCount二者只能设置一个，
		// WithMaxAge设置文件清理前的最长保存时间，
		// WithRotationCount设置文件清理前最多保存的个数。
		rotatelogs.WithMaxAge(time.Hour*(time.Duration)(maxRemain)),
		//rotatelogs.WithRotationCount(maxRemain),
	)
	if err != nil {
		log.Errorf("config local file system for logger error: %v", err)
	}

	log.SetLevel(level)

	lfsHook := lfshook.NewHook(lfshook.WriterMap{
		log.DebugLevel: writer,
		log.InfoLevel:  writer,
		log.WarnLevel:  writer,
		log.ErrorLevel: writer,
		log.FatalLevel: writer,
		log.PanicLevel: writer,
		//}, &log.TextFormatter{DisableColors: true})
	}, &log.JSONFormatter{})

	return lfsHook
}

func Setup2(serverName, path string, level log.Level, maxRemain uint) {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})
	//	log.SetFormatter(&log.TextFormatter{FullTimestamp: true, DisableColors: false})
	log.SetReportCaller(true)

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	//log.SetLevel(log.InfoLevel)

	if !util.Exists(path) {
		if !util.Mkdir(path) {
			log.WithFields(log.Fields{
				"err": "can not mkdir",
			}).Fatal()
		}
	}

	filename := filepath.Join(path, serverName)
	log.AddHook(newLfsHook(filename, level, maxRemain))
}
