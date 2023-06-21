package dao

import (
	"context"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"strings"
	"todoList/conf"
)

var _db *gorm.DB

func MySqlInit() {
	conn := strings.Join([]string{conf.DbUser, ":", conf.DbPassWord, "@tcp(", conf.DbHost, ":", conf.DbPort, ")/", conf.DbName, "?charset=utf8&parseTime=true"}, "")
	var ormLogger logger.Interface
	if gin.Mode() == "debug" {
		ormLogger = logger.Default.LogMode(logger.Info)
	}

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       conn, // 连接数据库
		DefaultStringSize:         256,  // string类型的默认字段
		DisableDatetimePrecision:  true, // 禁用datetime精度，mysql5.6不支持
		DontSupportRenameIndex:    true, // 重命名索引时采用删除并新建的方式，MySQL 5.7+
		DontSupportRenameColumn:   true, // 重命名列时采用删除并新建的方式，MySQL 5.7+
		SkipInitializeWithVersion: false,
	}), &gorm.Config{
		Logger: ormLogger, // 打印日志
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 使用单数表名
		},
	})
	if err != nil {
		panic(err)
	}

	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(20)     // 设置连接池
	sqlDB.SetMaxOpenConns(100)    // 设置最大连接数
	sqlDB.SetConnMaxLifetime(300) // 设置连接最大存活时间
	_db = db
	migration()
}

func NewDBClient(ctx context.Context) *gorm.DB {
	db := _db
	return db.WithContext(ctx)
}
