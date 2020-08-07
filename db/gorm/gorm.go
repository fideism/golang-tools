package gorm

import (
	"fmt"
	"github.com/jinzhu/gorm"
	// postgres
	_ "github.com/jinzhu/gorm/dialects/postgres"

	// sqlite
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// ConfigDialect db driver
type ConfigDialect string

// ConfigDialectPostgre postgre
const ConfigDialectPostgre ConfigDialect = "postgres"

// ConfigDialectSqliteV3 sqlite3
const ConfigDialectSqliteV3 ConfigDialect = `sqlite3`

// String ConfigDialect string
func (c ConfigDialect) String() string {
	return string(c)
}

// Config 数据库配置
type Config struct {
	Dialect  ConfigDialect
	Server   string
	Port     int
	User     string
	Database string
	Password string
}

// Conn conn
type Conn struct {
	*gorm.DB
}

// Build build
func Build(config Config) Conn {
	db, err := connect(config)

	if nil != err {
		panic("链接数据库失败")
	}

	return Conn{DB: db}

}

func connect(config Config) (*gorm.DB, error) {
	if config.Dialect == ConfigDialectPostgre {
		return gorm.Open(config.Dialect.String(), postgre(config))
	}

	if config.Dialect == ConfigDialectSqliteV3 {
		return gorm.Open(config.Dialect.String(), sqlite(config))
	}

	return nil, fmt.Errorf("没有数据库连接信息")
}

func postgre(config Config) string {
	return fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
		config.Server, config.Port, config.User, config.Database, config.Password)
}

func sqlite(config Config) string {
	return config.Server
}
