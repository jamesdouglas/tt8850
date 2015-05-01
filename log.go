package skypatrolTT8850

import (
	"github.com/op/go-logging"
	"os"
)

var log = logging.MustGetLogger("skypatrolTT8850")

func init() {
	// setup logger
	stdoutBackend := logging.NewLogBackend(os.Stdout, "", 0)
	format := logging.MustStringFormatter("%{color}%{time:15:04:05.000} %{shortfunc}[%{level:.4s} %{id:03x}]%{color:reset} %{message}")
	backendFmt := logging.NewBackendFormatter(stdoutBackend, format)
	logging.SetBackend(backendFmt)
}
