package models

import (
    "fmt"
    "github.com/gin-gonic/gin"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
    e "iam/pkg/exception"
    "iam/pkg/gredis"
    "iam/pkg/logging"
    setting "iam/pkg/settings"
    "iam/pkg/utils"
    "log"
    "net/http"
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

func AppCheck() gin.HandlerFunc {
    return func(c *gin.Context) {
        code := e.Success
        appID := c.GetHeader("APP_ID")
        appKey := c.GetHeader("APP_KEY")
        if appID == "" || appKey == "" {
            code = e.ErrorAppHeader
        } else {
            key := fmt.Sprintf("%s_%s-%s", utils.CACHE_APP, appID, appKey)
            data, err := gredis.Get(key)
            if err != nil {
                logging.Error(err)
            } else if data != "" {
                logging.Info(fmt.Sprintf("Key: %s is existed.", key))
            } else { // err为nil且数据为空的情况下，会重新从数据库获取app
                fmt.Println(data)
                maps := map[string]interface{}{
                    "access_key": appID,
                    "secret_key": appKey,
                }
                apps, err := GetApps(0, 1, maps)
                if err != nil || len(apps) == 0 {
                    logging.Error(e.GetMsg(e.ErrorAppIDKey))
                    code = e.ErrorAppIDKey
                }
                err = gredis.Set(key, apps[0].Name, time.Hour * 24)
                if err != nil {
                    logging.Error(err)
                }
            }
        }

        if code != e.Success {
            c.JSON(http.StatusInternalServerError, gin.H{
                "code": http.StatusInternalServerError,
                "msg": e.GetMsg(code),
            })
            c.Abort()
            return
        }
        c.Next()
    }
}
