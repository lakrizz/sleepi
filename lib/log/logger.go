package log

const (
	LOG_PANIC = iota
	LOG_FATAL
	LOG_ERROR
	LOG_WARNING
	LOG_INFO
	LOG_DEBUG
)

var Logfile string = "prod.log"
var DebugLevel int = LOG_DEBUG // Logrus has six logging levels: Debug (5), Info (4), Warning (3), Error (2), Fatal (1) and Panic (0).

func InitLogger() error {

}
