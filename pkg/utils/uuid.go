package utils

import "github.com/satori/go.uuid"

func GenerateUUID() string {
    uid := uuid.Must(uuid.NewV4()).String()
    return uid
}
