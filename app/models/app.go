package models

import (
    "github.com/jinzhu/gorm"
    "iam/pkg/utils"
)

//App 逻辑先写了，参数等可以参看role做优化

type App struct {
    Model
    Name           string       `gorm:"size:50;unique;not null"`
    IsActive       bool         `gorm:"column:is_active;size:1;not null"`
    Description    string       `gorm:"size:255"`
    IndexUrl       string       `gorm:"column:index_url;size:255;not null"`
    LoginUrl       string       `gorm:"column:login_url;size:255;not null"`
    //Deleted        int          `gorm:"size:1;not null;DEFAULT:0"`
    AccessKey      string       `gorm:"column:access_key;not null"`
    SecretKey      string       `gorm:"column:secret_key;not null"`
}

func AddApp(data map[string]interface{}) (*App, error) {
    app := App{
        Name:        data["name"].(string),
        IsActive:    data["is_active"].(bool),
        Description: data["description"].(string),
        IndexUrl:    data["index_url"].(string),
        LoginUrl:    data["login_url"].(string),
        AccessKey:   utils.GenerateUUID(),
        SecretKey:   utils.GenerateUUID(),
    }
    if err := db.Create(&app).Error; err != nil {
        return nil, err
    }
    return &app, nil
}

func DeleteApp(id int) error {
    if err := db.Where("id = ?", id).Delete(App{}).Error; err != nil {
        return err
    }
    return nil
}

func EditApp(id int, data map[string]interface{}) (*App, error) {
    if err := db.Model(&App{}).Where("id = ?", id).Updates(data).Error; err != nil {
        return nil, err
    }
    app, err := GetAppByID(id)
    return app, err
}

func ExistAppByName(name string) (bool, error) {
    var app App
    err := db.Select("id").Where("id = ?", name).First(&app).Error
    if err != nil && err != gorm.ErrRecordNotFound {
        return false, err
    }

    if app.ID > 0 {
        return true, nil
    }
    return false, nil
}

func GetApps(pageNum int, pageSize int, maps interface{}) ([]App, error) {
    var apps []App
    err := db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&apps).Error
    if err != nil && err != gorm.ErrRecordNotFound {
        return nil, err
    }

    return apps, nil
}

func GetAppByID(id int) (*App, error) {
    var app App
    err := db.Where("id = ?", id).First(&app).Error
    if err != nil {
        return nil, err
    }
    return &app, nil
}

func GetTotalApp(maps interface{}) (int, error) {
    var count int
    if err := db.Model(&App{}).Where(maps).Count(&count).Error; err != nil {
        return 0, err
    }
    return count, nil
}