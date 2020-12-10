package setting

import (
	"log"
	"time"

	"github.com/go-ini/ini"
)

type App struct {
	Version         string
	ServerID        int
	JwtSecret       string
	PageSize        int
	RuntimeRootPath string

	PrefixUrl string

	ImageSavePath  string
	ImageMaxSize   int
	ImageAllowExts []string

	ExportSavePath string

	QrCodeSavePath string

	LogSavePath string
	LogSaveName string
	LogFileExt  string
	TimeFormat  string

	StaticPath string
	NodeList   string

	AccessKeyId  string
	AccessSecret string
}

var AppSetting = &App{}

type Wx struct {
	AppId  string
	Secret string
}

var WxSetting = &Wx{}

type Http struct {
	RunMode      string
	Host         string
	Port         int
	WriteTimeout time.Duration
	ReadTimeout  time.Duration
}

var HttpSetting = &Http{}

type RpcDb struct {
	RunMode      string
	RPCHost      string
	RPCPort      int
	WriteTimeout time.Duration
	ReadTimeout  time.Duration
}

var RpcDbSetting = &RpcDb{}

type RpcHub struct {
	RunMode      string
	RPCHost      string
	RPCPort      int
	WriteTimeout time.Duration
	ReadTimeout  time.Duration
}

var RpcHubSetting = &RpcHub{}

type Backend struct {
	RunMode      string
	HTTPHost     string
	HTTPPort     int
	WriteTimeout time.Duration
	ReadTimeout  time.Duration
}

var BackendSetting = &Backend{}

type AdminDB struct {
	Type        string
	User        string
	Password    string
	Host        string
	Name        string
	TablePrefix string
}

var AdminDBSetting = &AdminDB{}

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
	Host        string
	Password    string
	MaxIdle     int
	MaxActive   int
	IdleTimeout time.Duration
}

var RedisSetting = &Redis{}

type Service struct {
	Host string
	Port int
}

var ServiceSetting = &Service{}

func Setup(filename string) {
	Cfg, err := ini.Load(filename)
	if err != nil {
		log.Fatalf("fail to parse '%s': %v", filename, err)
	}

	err = Cfg.Section("wx").MapTo(WxSetting)
	if err != nil {
		log.Fatalf("Cfg.MapTo WxSetting err: %v", err)
	}

	err = Cfg.Section("http").MapTo(HttpSetting)
	if err != nil {
		log.Fatalf("Cfg.MapTo HttpSetting err: %v", err)
	}

	err = Cfg.Section("app").MapTo(AppSetting)
	if err != nil {
		log.Fatalf("Cfg.MapTo AppSetting err: %v", err)
	}
	AppSetting.ImageMaxSize = AppSetting.ImageMaxSize * 1024 * 1024

	err = Cfg.Section("rpc_db").MapTo(RpcDbSetting)
	if err != nil {
		log.Fatalf("Cfg.MapTo RpcDbSetting err: %v", err)
	}
	RpcDbSetting.ReadTimeout = RpcDbSetting.ReadTimeout * time.Second
	RpcDbSetting.WriteTimeout = RpcDbSetting.ReadTimeout * time.Second

	err = Cfg.Section("rpc_hub").MapTo(RpcHubSetting)
	if err != nil {
		log.Fatalf("Cfg.MapTo RpcHubSetting err: %v", err)
	}
	RpcHubSetting.ReadTimeout = RpcHubSetting.ReadTimeout * time.Second
	RpcHubSetting.WriteTimeout = RpcHubSetting.ReadTimeout * time.Second

	err = Cfg.Section("backend").MapTo(BackendSetting)
	if err != nil {
		log.Fatalf("Cfg.MapTo BackendSetting err: %v", err)
	}
	BackendSetting.ReadTimeout = BackendSetting.ReadTimeout * time.Second
	BackendSetting.WriteTimeout = BackendSetting.ReadTimeout * time.Second

	err = Cfg.Section("database").MapTo(DatabaseSetting)
	if err != nil {
		log.Fatalf("Cfg.MapTo DatabaseSetting err: %v", err)
	}

	err = Cfg.Section("admindb").MapTo(AdminDBSetting)
	if err != nil {
		log.Fatalf("Cfg.MapTo AdminDBSetting err: %v", err)
	}

	err = Cfg.Section("redis").MapTo(RedisSetting)
	if err != nil {
		log.Fatalf("Cfg.MapTo RedisSetting err: %v", err)
	}
	RedisSetting.IdleTimeout = RedisSetting.IdleTimeout * time.Second

	err = Cfg.Section("service").MapTo(ServiceSetting)
	if err != nil {
		log.Fatalf("Cfg.MapTo ServiceSetting err: %v", err)
	}
}
