package gomysql

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var (
	DefaultHost        = "127.0.0.1"
	DefaultPort        = 3306
	DefaultCharset     = "utf8mb4"
	DefaultCollation   = "utf8mb4_unicode_ci"
	DefaultEngine      = "INNODB"
	DefaultTimezone    = "PRC"
	DefaultMaxIdleConn = 10
	DefaultMaxOpenConn = 100
)

type Configs struct {
	Host        string
	Port        int
	User        string
	Password    string
	Database    string
	Charset     string
	Collation   string
	Engine      string
	Timezone    string
	MaxIdleConn int
	MaxOpenConn int
	LogMode     bool
}

type Config func(c *Configs)

// NewDB 新建连接
func NewDB(user, password, dbname string, configs ...Config) *gorm.DB {
	return newDB(user, password, dbname, configs...)
}

func Host(h string) Config {
	return func(c *Configs) {
		c.Host = h
	}
}

func Port(p int) Config {
	return func(c *Configs) {
		c.Port = p
	}
}

func Charset(s string) Config {
	return func(c *Configs) {
		c.Charset = s
	}
}

func Collation(s string) Config {
	return func(c *Configs) {
		c.Collation = s
	}
}

func Engine(s string) Config {
	return func(c *Configs) {
		c.Engine = s
	}
}

func Timezone(s string) Config {
	return func(c *Configs) {
		c.Timezone = s
	}
}

func MaxIdleConn(i int) Config {
	return func(c *Configs) {
		c.MaxIdleConn = i
	}
}

func MaxOpenConn(i int) Config {
	return func(c *Configs) {
		c.MaxOpenConn = i
	}
}

func LogMode(b bool) Config {
	return func(c *Configs) {
		c.LogMode = b
	}
}

func newConfigs(user, password, dbname string, configs ...Config) Configs {
	cnf := Configs{
		Host:        DefaultHost,
		Port:        DefaultPort,
		User:        user,
		Password:    password,
		Database:    dbname,
		Charset:     DefaultCharset,
		Collation:   DefaultCollation,
		Engine:      DefaultEngine,
		Timezone:    DefaultTimezone,
		MaxIdleConn: DefaultMaxIdleConn,
		MaxOpenConn: DefaultMaxOpenConn,
		LogMode:     false,
	}

	for _, c := range configs {
		c(&cnf)
	}

	return cnf
}

func newDB(user, password, dbname string, configs ...Config) *gorm.DB {
	c := newConfigs(user, password, dbname, configs...)
	db, err := gorm.Open("mysql",
		fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&collation=%s&parseTime=true&loc=%s",
			c.User,
			c.Password,
			c.Host,
			c.Port,
			c.Database,
			c.Charset,
			c.Collation,
			c.Timezone,
		))
	if err != nil {
		panic(err)
	}
	if err := db.DB().Ping(); err != nil {
		panic(err)
	}
	db.LogMode(c.LogMode)
	db.DB().SetMaxIdleConns(c.MaxIdleConn)
	db.DB().SetMaxOpenConns(c.MaxOpenConn)
	return db
}
