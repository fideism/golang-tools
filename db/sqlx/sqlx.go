package sqlx

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"github.com/fideism/golang-tools/logger"

	// Go postgres driver
	_ "github.com/lib/pq"
)

// Config 数据库配置
type Config struct {
	Dialect  string
	Server   string
	Port     int
	User     string
	Database string
	Password string
}

// Conn 链接对象
type Conn struct {
	*sqlx.DB
}

// Connection db connection
var Connection Conn

// Build 数据库连接
func Build(config Config) Conn {
	db, err := connect(config)
	if nil != err {
		panic("连接数据库失败")
	}

	return Conn{DB: db}
}

// 数据库连接
func connect(conf Config) (db *sqlx.DB, err error) {
	source := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
		conf.Server, conf.Port, conf.User, conf.Database, conf.Password)

	db, err = sqlx.Connect(conf.Dialect, source)
	if nil != err {
		return nil, err
	}

	return db.Unsafe(), err
}

// Get 重载 get 方法 记录日志
func (db Conn) Get(dest interface{}, query string, args ...interface{}) error {
	logger.NewEntry(logger.Channel("sql")).WithFields(logrus.Fields{
		"args":  args,
		"query": query,
	}).Info("GET")

	return db.DB.Get(dest, query, args...)
}

// Select 重载 select
func (db Conn) Select(dest interface{}, query string, args ...interface{}) error {
	logger.NewEntry(logger.Channel("sql")).WithFields(logrus.Fields{
		"args":  args,
		"query": query,
	}).Info("SELECT")

	return db.DB.Select(dest, query, args...)
}

// Queryx 重载 Queryx
func (db Conn) Queryx(query string, args ...interface{}) (*sqlx.Rows, error) {
	logger.NewEntry(logger.Channel("sql")).WithFields(logrus.Fields{
		"args":  args,
		"query": query,
	}).Info("QUERYX")

	return db.DB.Queryx(query, args...)
}

// QueryRowx 重载 QueryRowx
func (db Conn) QueryRowx(query string, args ...interface{}) *sqlx.Row {
	logger.NewEntry(logger.Channel("sql")).WithFields(logrus.Fields{
		"args":  args,
		"query": query,
	}).Info("QUERYROWX")

	return db.DB.QueryRowx(query, args...)
}
