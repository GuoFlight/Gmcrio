package db

import (
	"Gmicro/conf"
	"Gmicro/logger"
	"context"
	"github.com/GuoFlight/gerror/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

var Conn *gorm.DB

func Init(ctx context.Context) {
	var gerr *gerror.Gerr
	Conn, gerr = Connect(ctx)
	if gerr != nil {
		logger.GLogger.Fatal(gerr.Error())
	}
}
func Exit() {
	logger.GLogger.Info("Gorm即将退出")
	sqlDB, err := Conn.DB()
	if err != nil {
		logger.GLogger.Errorf("gorm异常: %s", err.Error())
	}
	err = sqlDB.Close()
	if err != nil {
		logger.GLogger.Errorf("gorm关闭异常: %s", err.Error())
	}
	logger.GLogger.Info("Gorm退出完成")
}
func Connect(ctx context.Context) (*gorm.DB, *gerror.Gerr) {
	db, err := gorm.Open(mysql.Open(conf.GConf.Db.Server), &gorm.Config{})
	if err != nil {
		gerr := logger.ErrWithCtx(ctx, "数据库连接失败: "+err.Error())
		return nil, gerr
	}
	sqlDB, err := db.DB()
	if err != nil {
		gerr := logger.ErrWithCtx(ctx, err.Error())
		return nil, gerr
	}

	sqlDB.SetMaxIdleConns(conf.GConf.Db.MaxIdleConns)
	sqlDB.SetMaxOpenConns(conf.GConf.Db.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(conf.GConf.Db.ConnMaxLifetime * time.Second)
	return db, nil
}
