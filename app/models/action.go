package models

import "github.com/jinzhu/gorm"

type Action struct {
    Model
    Name            string      `gorm:"column:name;not null"`
    DisplayName     string      `gorm:"column:display_name;not null"`
    IsActive        bool        `gorm:"column:is_active;size:1;not null"`
    AppID           int         `gorm:"column:app_id;not null"`
    App             App
}

func (action *Action) AddAppAction() (*Action, error) {
    if err := db.Create(action).Error; err != nil {
        return nil, err
    }
    return action, nil
}

func (action *Action) DeleteAppAction(id int) error {
    if err := db.Where("id = ?", id).Delete(Action{}).Error; err != nil {
        return err
    }
    return nil
}

func (action *Action) EditAppAction(id int, data map[string]interface{}) (*Action, error) {
    if err := db.Model(&action).Where("id = ?", id).Updates(data).Error; err != nil {
        return nil, err
    }
    r, err := action.GetAppActionByID(id)
    return r, err
}

func (action *Action) ListAppActions(pageNum int, pageSize int, maps interface{}) ([]Action, error) {
    var actions []Action
    err := db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&actions).Error
    if err != nil && err != gorm.ErrRecordNotFound {
        return nil, err
    }
    return actions, nil
}

func (action *Action) CountAppActions(maps interface{}) (int, error) {
    var count int
    if err := db.Model(&Action{}).Where(maps).Count(&count).Error; err != nil {
        return 0, err
    }
    return count, nil
}

func (action *Action) GetAppActionByID(maps interface{}) (*Action, error) {
    if err := db.Where(maps).First(&action).Error; err != nil {
        return nil, err
    }
    return action, nil
}

