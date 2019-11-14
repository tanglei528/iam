package models

import (
    "fmt"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
    setting "iam/pkg/settings"
    "log"
    "time"
)

var db *gorm.DB

type Model struct {
    ID             int          `gorm:"PRIMARY_KEY;AUTO_INCREMENT;UNIQUE"`
    CreatedAt      time.Time
    UpdatedAt      time.Time
    DeletedAt      *time.Time
}

func Setup() {
    var err error
    db, err = gorm.Open(
        setting.DatabaseSetting.Type,
        fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
        setting.DatabaseSetting.User,
        setting.DatabaseSetting.Password,
        setting.DatabaseSetting.Host,
        setting.DatabaseSetting.Name,
    ))

    if err != nil {
        log.Fatalf("models.Setup err: %v", err)
    }

    // 设置表名前缀
    //gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
    //    return setting.DatabaseSetting.TablePrefix + defaultTableName
    //}

    // 表名单数形式
    db.SingularTable(true)
    db.DB().SetMaxIdleConns(10)
    db.DB().SetMaxOpenConns(100)
    db.LogMode(true)
}

func CloseDB() {
    defer db.Close()
}