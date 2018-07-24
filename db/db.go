package db

import (
	"github.com/echoturing/prospect-crawler/etc"
	"github.com/echoturing/prospect-crawler/logger"
	"github.com/go-sql-driver/mysql"
	"github.com/gocraft/dbr"
	"go.uber.org/zap"
)

func NewConn(cfg *etc.Config) (*dbr.Connection, error) {
	mysqlConfig := mysql.NewConfig()
	mysqlConfig.User = cfg.Mysql.User
	mysqlConfig.Passwd = cfg.Mysql.Password
	mysqlConfig.Net = cfg.Mysql.Net
	mysqlConfig.Addr = cfg.Mysql.Addr
	mysqlConfig.ParseTime = cfg.Mysql.ParseTime
	mysqlConfig.DBName = cfg.Mysql.DBName
	mysqlConfig.Params = make(map[string]string)
	mysqlConfig.Params["charset"] = cfg.Mysql.Charset
	dsn := mysqlConfig.FormatDSN()
	log := logger.GetLogger()
	log.Info("parse dsn", zap.String("dsn", dsn))
	conn, err := dbr.Open("mysql", dsn, nil)
	if err != nil {
		return nil, err
	}
	return conn, nil
}
