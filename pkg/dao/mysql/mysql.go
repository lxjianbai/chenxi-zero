package mysql

import (
	"errors"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Mysql struct {
	Host          string // 服务器地址
	Port          int    `json:",default=3306"`                                               // 端口
	Config        string `json:",default=charset%3Dutf8mb4%26parseTime%3Dtrue%26loc%3DLocal"` // 高级配置
	DbName        string // 数据库名
	Username      string // 数据库用户名
	Password      string // 数据库密码
	MaxIdleConns  int    `json:",default=10"`                               // 空闲中的最大连接数
	MaxOpenConns  int    `json:",default=10"`                               // 打开到数据库的最大连接数
	LogMode       string `json:",default=dev,options=dev|test|prod|silent"` // 是否开启Gorm全局日志
	LogZap        bool   // 是否通过zap写入日志文件
	SlowThreshold int64  `json:",default=1000"`
}

func (m *Mysql) Dsn() string {
	return m.Username + ":" + m.Password + "@tcp(" + m.Host + ":" + fmt.Sprintf("%d", m.Port) + ")/" + m.DbName + "?" + m.Config
}

func ConnectMysql(m Mysql) (*gorm.DB, error) {
	if m.DbName == "" {
		return nil, errors.New("database name is empty")
	}
	mysqlCfg := mysql.Config{
		DSN: m.Dsn(),
	}
	db, err := gorm.Open(mysql.New(mysqlCfg), &gorm.Config{
		Logger: NewGormLogger(m.LogMode),
	})
	if err != nil {
		return nil, err
	} else {
		sqldb, _ := db.DB()
		sqldb.SetMaxIdleConns(m.MaxIdleConns)
		sqldb.SetMaxOpenConns(m.MaxOpenConns)
		return db, nil
	}
}
