package main

import (
	"crypto/sha512"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/anaskhan96/go-password-encoder"

	"crypto/md5"
	"database/sql"
	"encoding/hex"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"

	"github.com/liuyongbing/hello-go-srvs/goods_srv/model"
)

type Product struct {
	gorm.Model
	Code  sql.NullString
	Price uint
}

func genMd5(code string) string {
	Md5 := md5.New()
	_, _ = io.WriteString(Md5, code)
	return hex.EncodeToString(Md5.Sum(nil))
}

/*
initTable:创建表 user
*/
func initTable(db *gorm.DB, user *model.User) {
	db.AutoMigrate(&user)
}

func batchCreateUser(db *gorm.DB) error {
	// Using custom options
	options := &password.Options{SaltLen: 16, Iterations: 100, KeyLen: 32, HashFunction: sha512.New}
	salt, encodedPwd := password.Encode("generic password", options)
	newPassword := fmt.Sprintf("$pbkdf2-sha512$%s$%s", salt, encodedPwd)
	fmt.Println(len(newPassword))
	fmt.Println(newPassword)

	for i := 0; i < 10; i++ {
		user := model.User{
			Nickname: fmt.Sprintf("Nickname_%d", i),
			Mobile:   fmt.Sprintf("1881234567%d", i),
			Password: newPassword,
		}
		db.Save(&user)
	}

	return nil
}

func main() {
	// @link https://gorm.io/docs/logger.html
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			Colorful:                  true,        // Disable color
		},
	)
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	dsn := "root:mysql.root@tcp(127.0.0.1:3306)/hello_gosrvs?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: newLogger,
	})
	if err != nil {
		panic(err)
	}

	// 初始化数据表
	//err = db.AutoMigrate(&model.User{})
	// 初始化用户数据
	err = batchCreateUser(db)
	if err != nil {
		panic(err)
	}

}
