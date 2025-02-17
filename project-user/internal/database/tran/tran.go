package tran

import "github.com/WildWeber-xiaocheng/ms_project/project-user/internal/database"

// 事务的操作一定和数据库有关，所以需要注入数据库的连接 gorm.db
type Transaction interface {
	Action(func(conn database.DbConn) error) error
}
