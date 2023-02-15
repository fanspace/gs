package core

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"os"
	"time"
)

type Config struct {
	ReleaseMode      bool
	DaprMode         bool
	Smark            string
	AppName          string
	HttpSettings     *HttpSettings
	Database         *Database
	CorsSettings     *CorsSettings
	LogSettings      *LogSettings
	MinioSettings    *MinioSettings
	GrpcSettings     *GrpcSettings
	RabbitMqSettings *RabbitMqSettings
	FtpSettings      *FtpSettings
}
type Database struct {
	MysqlSettings *MysqlSettings
	RedisSettings *RedisSettings
}
type MysqlSettings struct {
	DriverName   string
	Url          string
	MaxIdleConns int32
	MaxOpenConns int32
	QueryTimeout int32
	UseCache     bool
}

type RedisSettings struct {
	Addr        string
	DB          int
	Password    string
	PoolSize    int
	MaxIdle     int
	MaxActive   int
	IdleTimeout int
}
type LogSettings struct {
	EnableConsole bool
	ConsoleLevel  string
	ConsoleJson   *bool
	EnableFile    bool
	FileLevel     string
	FileJson      *bool
	FileLocation  string
}
type CorsSettings struct {
	AllowOrigins []string
	AllowMethods []string
}
type HttpSettings struct {
	ListenAddress string
	ReadTimeout   int
	WriteTimeout  int
}

type MinioSettings struct {
	KeyID         string
	AccessKey     string
	PublicBucket  string // 公用
	PrivateBucket string // 私有
	EndPoint      string
	EntryPoint    string
}

type GrpcSettings struct {
	TimeOut  time.Duration
	EndPoint map[string]string
}

type RabbitMqSettings struct {
	Url  string
	Port int32
	User string
	Pwd  string
}

type FtpSettings struct {
	Ftpurl  string
	Ftpuser string
	Ftppwd  string
}

var Cfg *Config = &Config{}
var VOptions *viper.Viper

func init() {
	log.Println("loading system configure......")
	vs := viper.New()
	vs.SetConfigName("config.yaml")
	vs.AddConfigPath("./config")
	vs.SetConfigType("yaml")
	if err := vs.ReadInConfig(); err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}
	if err := vs.Unmarshal(Cfg); err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}
	// 以下为判断
	if Cfg.Smark == "" {
		Cfg.Smark = GetIpStr()
	}
	initOptions()
	Cfg.PrintConfig()
}

func initOptions() {
	log.Println("加载自定义配置文件......")
	VOptions = viper.New()
	VOptions.SetConfigName("options.yaml")
	VOptions.AddConfigPath("./conf")
	VOptions.SetConfigType("yaml")
	if err := VOptions.ReadInConfig(); err != nil {
		log.Println(err.Error())
		//os.Exit(1)
	}

}

func (s *LogSettings) PrintLogConfig() {
	log.Println(fmt.Sprintf("EnableConsole:%v", s.EnableConsole))
	log.Println(fmt.Sprintf("ConsoleLevel:%v", s.ConsoleLevel))
	log.Println(fmt.Sprintf("ConsoleJson:%v", *s.ConsoleJson))
	log.Println(fmt.Sprintf("EnableFile:%v", s.EnableFile))
	log.Println(fmt.Sprintf("FileLevel:%v", s.FileLevel))
	log.Println(fmt.Sprintf("FileJson:%v", *s.FileJson))
	log.Println(fmt.Sprintf("FileLocation:%v", s.FileLocation))
}

func (o *Config) PrintConfig() {
	o.LogSettings.PrintLogConfig()
	log.Println(fmt.Sprintf("ReleaseMode:%v", o.ReleaseMode))
}
