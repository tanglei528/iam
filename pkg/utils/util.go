package utils

import (
    "golang.org/x/crypto/bcrypt"
    "strconv"
)

func IntToBool(p int) (bool, error) {
    s := strconv.Itoa(p)
    b, err := strconv.ParseBool(s)
    return b, err
}

func ConvErrorToSlice(err error, errorsInfo []string) []string {
    errorsInfo = append(errorsInfo, err.Error())
    return errorsInfo
}

func GeneratePassword(password string) (string, error) {
    hashPWD, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        return "", err
    }
    return string(hashPWD), nil
}

func CheckPassword(dbPassword, webPassword string) error {
    err := bcrypt.CompareHashAndPassword([]byte(dbPassword), []byte(webPassword))
    return err
}