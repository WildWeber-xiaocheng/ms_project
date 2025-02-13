package config

import (
	"fmt"
	"github.com/WildWeber-xiaocheng/ms_project/project-project/internal/database/gorms"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var _db *gorm.DB

func (c *Config) ReConnMysql() {
	//配置MySQL连接参数
	username := c.MysqlConfig.Username //账号
	password := c.MysqlConfig.Password //密码
	host := c.MysqlConfig.Host         //数据库地址，可以是Ip或者域名
	port := c.MysqlConfig.Port         //数据库端口
	Dbname := c.MysqlConfig.Db         //数据库名
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", username, password, host, port, Dbname)
	var err error
	_db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		zap.L().Error("连接数据库失败", zap.Error(err))
		//panic("连接数据库失败, error=" + err.Error())
	}
	gorms.SetDB(_db)
}
