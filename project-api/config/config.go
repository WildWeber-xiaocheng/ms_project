package config

import (
	"github.com/WildWeber-xiaocheng/ms_project/project-common/logs"
	"github.com/spf13/viper"
	"log"
	"os"
)

var Conf = InitConfig()

type Config struct {
	viper      *viper.Viper
	SC         *ServerConfig
	GC         *GrpcConfig
	EtcdConfig *EtcdConfig
}

type ServerConfig struct {
	Name string
	Addr string
}

type GrpcConfig struct {
	Name string
	Addr string
}

type EtcdConfig struct {
	Addrs []string
}

func InitConfig() *Config {
	v := viper.New()
	conf := &Config{viper: v}
	workDir, _ := os.Getwd()
	conf.viper.SetConfigName("config")
	conf.viper.SetConfigType("yaml")
	conf.viper.AddConfigPath(workDir + "/config")

	err := conf.viper.ReadInConfig()
	if err != nil {
		log.Fatalln(err)
		return nil
	}
	conf.ReadServerConfig()
	conf.InitZapLog()
	conf.ReadEtcdConfig()
	return conf
}

func (c *Config) ReadServerConfig() {
	sc := &ServerConfig{}
	sc.Name = c.viper.GetString("server.name")
	sc.Addr = c.viper.GetString("server.addr")
	c.SC = sc
}

func (c *Config) InitZapLog() {
	//从配置中读取日志配置，初始化日志
	lc := &logs.LogConfig{
		DebugFileName: "C:\\goCode\\ms_project\\logs\\debug\\project-debug.log",
		InfoFileName:  "C:\\goCode\\ms_project\\logs\\info\\project-info.log",
		WarnFileName:  "C:\\goCode\\ms_project\\logs\\error\\project-error.log",
		MaxSize:       500,
		MaxAge:        28,
		MaxBackups:    3,
	}
	err := logs.InitLogger(lc)
	if err != nil {
		log.Fatalln(err)
	}
}

func (c *Config) ReadEtcdConfig() {
	ec := &EtcdConfig{}
	var addrs []string
	err := c.viper.UnmarshalKey("etcd.addrs", &addrs)
	if err != nil {
		log.Fatalln(err)
	}
	ec.Addrs = addrs
	c.EtcdConfig = ec
}
