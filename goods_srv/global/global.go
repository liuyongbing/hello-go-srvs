package global

import (
	"gorm.io/gorm"

	"github.com/liuyongbing/hello-go-srvs/goods_srv/config"
)

var (
	DB           *gorm.DB
	ServerConfig config.ServerConfig
	NacosConfig  config.NacosConfig
)

func init() {
	// @link https://gorm.io/docs/logger.html
	// newLogger := logger.New(
	// 	log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
	// 	logger.Config{
	// 		SlowThreshold:             time.Second, // Slow SQL threshold
	// 		LogLevel:                  logger.Info, // Log level
	// 		IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
	// 		Colorful:                  true,        // Disable color
	// 	},
	// )
	// // 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	// dsn := "root:mysql.root@tcp(127.0.0.1:3306)/hello_gosrvs?charset=utf8mb4&parseTime=True&loc=Local"

	// // 全局模式
	// var err error
	// DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
	// 	NamingStrategy: schema.NamingStrategy{
	// 		SingularTable: true,
	// 	},
	// 	Logger: newLogger,
	// })
	// if err != nil {
	// 	panic(err)
	// }

}
