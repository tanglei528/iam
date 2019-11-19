package models

import "github.com/jinzhu/gorm"

type User struct {
    Model
    Name     string     `gorm:"not null"`
    Email    string     `gorm:"unique;not null"`
    Password string     `gorm:"not null"`
    Phone    string     `gorm:"not null"`
    Address  string     `gorm:"not null"`
}

func (user *User) AddUser() (*User, error) {
    if err := db.Create(user).Error; err != nil {
        return nil, err
    }
    return user, nil
}

func (user *User) DeleteUser(id int) error {
    if err := db.Where("id = ?", id).Delete(user).Error; err != nil {
        return err
    }
    return nil
}

func (user *User) EditUser(id int, data map[string]interface{}) (*User, error) {
    if err := db.Model(&user).Where("id = ?", id).Updates(data).Error; err != nil {
        return nil, err
    }
    r, err := user.GetUserByID(id)
    return r, err
}

func (user *User) CountUsers(maps interface{}) (int, error) {
    var count int
    if err := db.Model(&user).Where(maps).Count(&count).Error; err != nil {
        return 0, err
    }
    return count, nil
}

func (user *User) ListUsers(pageNum int, pageSize int, maps interface{}) ([]User, error) {
    var users []User
    err := db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&users).Error
    if err != nil && err != gorm.ErrRecordNotFound {
        return nil, err
    }

    return users, nil
}

func (user *User) GetUserByID(id int) (*User, error) {
    if err := db.Where("id = ?", id).First(&user).Error; err != nil {
        return nil, err
    }
    return user, nil
}