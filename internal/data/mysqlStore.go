package data

import (
	"errors"
	"github.com/2pgcn/auth/internal/conf"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"gorm.io/plugin/dbresolver"
	gormopentracing "gorm.io/plugin/opentracing"
)

type mysqlStore struct {
	die chan struct{}
	db  *gorm.DB
}

func OpenDBConnection(c *conf.Data_Mysql) (db *gorm.DB, err error) {
	//newLogger := New(
	//	Config{
	//		SlowThreshold:             200 * time.Millisecond, // 慢 SQL 阈值
	//		LogLevel:                  getGormLevel(),         // 日志级别
	//		IgnoreRecordNotFoundError: true,                   // 忽略ErrRecordNotFound（记录未找到）错误
	//		Colorful:                  false,                  // 禁用彩色打印
	//	},
	//)
	ormConfig := &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		//Logger: newLogger,
	}
	db, err = gorm.Open(mysql.New(mysql.Config{
		DSN:                       c.GetDsnMaster(), // DSN data source name
		DefaultStringSize:         256,              // string 类型字段的默认长度
		DisableDatetimePrecision:  true,             // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,             // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,             // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false,            // 根据版本自动配置
	}), ormConfig)
	if err != nil {
		panic(err)
		return
	}
	replicas := []gorm.Dialector{}
	for i := 0; i < len(c.DsnSlave); i++ {
		replicas = append(replicas, mysql.Open(c.DsnSlave[i]))
	}
	db.Use(dbresolver.Register(dbresolver.Config{
		Sources:  []gorm.Dialector{mysql.Open(c.GetDsnMaster())},
		Replicas: replicas,
		Policy:   dbresolver.RandomPolicy{},
	}).
		SetMaxIdleConns(int(c.GetMaxIdleConnection())).
		SetConnMaxLifetime(c.GetConnectionMaxLifetime().AsDuration()).
		SetMaxOpenConns(int(c.GetMaxOpenConnection())),
	)
	db.Use(gormopentracing.New())
	return
}

// RecordNotFound First、Last、Take 这些预期会返回结果的方法查询记录时，才会返回 ErrRecordNotFound
func RecordNotFound(db *gorm.DB) (ret bool) {
	if db.Error != nil {
		if errors.Is(db.Error, gorm.ErrRecordNotFound) {
			ret = true
		}
	}
	return
}
