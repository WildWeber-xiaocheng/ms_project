package config

import (
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"log"
	"os"
	"test.com/project-common/logs"
)

var Conf = InitConfig()

type Config struct {
	viper *viper.Viper
	SC    *ServerConfig
}

type ServerConfig struct {
	Name string
	Addr string
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

func (c *Config) InitRedisOptions() *redis.Options {
	return &redis.Options{
		Addr:     c.viper.GetString("redis.host") + ":" + c.viper.GetString("redis.port"),
		Password: c.viper.GetString("redis.password"),
		DB:       c.viper.GetInt("redis.db"), // use default DB
	}
}
