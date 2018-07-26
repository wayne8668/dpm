package common

import (
	"dpm/vars"
	"fmt"
	"github.com/op/go-logging"
	"os"
)

const (
	LOG_FORMAT = `%{color}%{time:0102 15:04:05.999999} %{longfile} %{longfunc} > %{level:.4s} %{id:03x}%{color:reset} %{message}`
)

var (
	filePath = vars.Cfg.Get("logger.file_path").(string)
	Logger   = logging.MustGetLogger(vars.PROJECT_NAME)
	format   = logging.MustStringFormatter(LOG_FORMAT)
)

type (
	Password string
)

func (p Password) Redacted() interface{} {
	return logging.Redact(string(p))
}

func init() {
	logFile, err := os.OpenFile(filePath, os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println(err)
	}
	backendFile := logging.NewLogBackend(logFile, "", 0)
	backendStderr := logging.NewLogBackend(os.Stderr, "", 0)

	backendFileFormatter := logging.NewBackendFormatter(backendFile, format)
	backendStderrFormatter := logging.NewBackendFormatter(backendStderr, format)

	backendFileLeveled := logging.AddModuleLevel(backendFileFormatter)
	backendStderrLeveled := logging.AddModuleLevel(backendStderrFormatter)

	backendFileLeveled.SetLevel(logging.DEBUG, "")
	backendStderrLeveled.SetLevel(logging.DEBUG, "")

	logging.SetBackend(backendFileLeveled, backendStderrLeveled)

	// Logger.Debugf("debug %s", Password("secret"))
	// Logger.Info("info")
	// Logger.Notice("notice")
	// Logger.Warning("warning")
	// Logger.Error("xiaorui.cc")
	// Logger.Critical("太严重了")
}
