package main

import (
    "fmt"
    "github.com/fvbock/endless"
    "iam/app/models"
    "iam/app/routers"
    "iam/pkg/logging"
    setting "iam/pkg/settings"
    "iam/pkg/validate"
    "log"
    "syscall"
)

func main() {
    setting.Setup()
    models.Setup()
    logging.Setup()
    //gredis.Setup()
    validate.SetUp()

    endless.DefaultReadTimeOut = setting.ServerSetting.ReadTimeout
    endless.DefaultWriteTimeOut = setting.ServerSetting.WriteTimeout
    endless.DefaultMaxHeaderBytes = 1 << 20
    endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)

    server := endless.NewServer(endPoint, routers.InitRouter())
    server.BeforeBegin = func(add string) {
        log.Printf("Actual pid is %d", syscall.Getpid())
    }

    err := server.ListenAndServe()
    if err != nil {
        log.Printf("Server err: %v", err)
    }
}