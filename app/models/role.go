package models

import "github.com/jinzhu/gorm"

type Role struct {
    Model
    Name           string       `gorm:"column:name;not null"`
    Description    string       `gorm:"column:description;not null"`
    Permission     string       `gorm:"permission;not null"`
    App            App
    AppID          int          `gorm:"app_id;not null"`
}

func (role *Role) AddAppRole() (*Role, error) {
    if err := db.Create(role).Error; err != nil {
        return nil, err
    }
    return role, nil
}

func (role *Role) DeleteAppRole(id int) error {
    if err := db.Where("id = ?", id).Delete(Role{}).Error; err != nil {
        return err
    }
    return nil
}

func (role *Role) EditAppRole(id int, data map[string]interface{}) (*Role, error) {
    if err := db.Model(&role).Where("id = ?", id).Updates(data).Error; err != nil {
        return nil, err
    }
    r, err := role.GetAppRoleByID(id)
    return r, err
}

func (role *Role) ListAppRoles(pageNum int, pageSize int, maps interface{}) ([]Role, error) {
    var roles []Role
    err := db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&roles).Error
    if err != nil && err != gorm.ErrRecordNotFound {
        return nil, err
    }
    return roles, nil
}

func (role *Role) CountAppRoles(maps interface{}) (int, error) {
    var count int
    if err := db.Model(&Role{}).Where(maps).Count(&count).Error; err != nil {
        return 0, err
    }
    return count, nil
}

func (role *Role) GetAppRoleByID(maps interface{}) (*Role, error) {
    if err := db.Where(maps).First(&role).Error; err != nil {
        return nil, err
    }
    return role, nil
}