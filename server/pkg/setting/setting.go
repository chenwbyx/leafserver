package setting

import (
	"log"
	"time"
	"github.com/go-ini/ini"
)

var (
	// gate conf
	PendingWriteNum        = 2000
	MaxMsgLen       uint32 = 4096
	HTTPTimeout            = 10 * time.Second
	LenMsgLen              = 2
	LittleEndian           = false

	// skeleton conf
	GoLen              = 10000
	TimerDispatcherLen = 10000
	AsynCallLen        = 10000
	ChanRPCLen         = 10000
)

type App struct {
	RuntimeRootPath string

	ExportSavePath string

	LogSavePath string
	LogSaveName string
	LogFileExt  string
	TimeFormat  string
	LogFlag     int
}

var AppSetting = &App{}

type Server struct {
	LogLevel     string
	HttpPort     int
	TcpAddr      int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	MaxConnNum   int
	ConsolePort  int
	ProfilePath  string

	WSAddr      string
	CertFile    string
	KeyFile     string
	TCPAddr     string
}

var ServerSetting = &Server{}

type Log struct {
	ServiceName string
	App         string
	Http        string
	Default     string
	Login       string
}

var LogSetting = &Log{}

type Database struct {
	Type        string
	User        string
	Password    string
	Host        string
	Name        string
	TablePrefix string
}

var DatabaseSetting = &Database{}

type Redis struct {
	Host     string
	Password string
	DB       int
	Cluster  bool
	Hosts    []string
}

var RedisSetting = &Redis{}

var cfg *ini.File

// Setup initialize the configuration instance
func Setup() {
	var err error
	cfg, err = ini.Load("conf/app.ini")
	if err != nil {
		log.Fatalf("setting.Setup, fail to parse 'conf/app.ini': %v", err)
	}

	mapTo("app", AppSetting)
	mapTo("server", ServerSetting)
	mapTo("database", DatabaseSetting)
	mapTo("redis", RedisSetting)
	mapTo("log", LogSetting)

	AppSetting.LogFlag = log.LstdFlags
	ServerSetting.ReadTimeout = ServerSetting.ReadTimeout * time.Second
	ServerSetting.WriteTimeout = ServerSetting.WriteTimeout * time.Second
	RedisSetting.IdleTimeout = RedisSetting.IdleTimeout * time.Second
}

// mapTo map section
func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("Cfg.MapTo %s err: %v", section, err)
	}
}