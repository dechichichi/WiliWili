package client

import (
	"context"
	"fmt"
	"wiliwili/pkg/constants"
	"wiliwili/pkg/utils"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func InitMySQL() (db *gorm.DB, err error) {
	dsn, err := utils.GetMysqlDSN()
	if err != nil {
		return nil, err
	}
	db, err = gorm.Open(mysql.Open(dsn),
		&gorm.Config{
			PrepareStmt:            true,  // 在执行任何 SQL 时都会创建一个 prepared statement 并将其缓存，以提高后续的效率
			SkipDefaultTransaction: false, // 不禁用默认事务(即单个创建、更新、删除时使用事务)
			TranslateError:         true,  // 允许翻译错误
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true, // 使用单数表名
			},
		})
	if err != nil {
		return nil, err
	}
	sqlDB, _ := db.DB() // 尝试获取 DB 实例对象

	sqlDB.SetMaxIdleConns(constants.MaxIdleConns)       // 最大闲置连接数
	sqlDB.SetMaxOpenConns(constants.MaxConnections)     // 最大连接数
	sqlDB.SetConnMaxLifetime(constants.ConnMaxLifetime) // 最大可复用时间
	sqlDB.SetConnMaxIdleTime(constants.ConnMaxIdleTime) // 最长保持空闲状态时间
	db = db.WithContext(context.Background())
	// 进行连通性测试
	if err = sqlDB.Ping(); err != nil {
		return nil, fmt.Errorf("failed to connect to mysql: %v", err)
	}
	return db, nil
}
