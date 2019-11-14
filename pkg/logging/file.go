package logging

import (
    "fmt"
    "iam/pkg/file"
    "iam/pkg/settings"
    "os"
    "time"
)

func getLogFilePath() string {
    return fmt.Sprintf("%s%s", settings.AppSetting.RuntimeRootPath, settings.AppSetting.LogSavePath)
}

func getLogFileName() string {
    return fmt.Sprintf("%s%s.%s",
        settings.AppSetting.LogSaveName,
        time.Now().Format(settings.AppSetting.TimeFormat),
        settings.AppSetting.LogFileExt,
    )
}

func openLogFile(fileName, filePath string) (*os.File, error) {
    // os.Stat：返回文件信息结构描述文件。如果出现错误，会返回*PathError
    dir, err := os.Getwd()
    if err != nil {
        return nil, fmt.Errorf("os.Getwd err: %v", err)
    }

    src := dir + "/" + filePath
    perm := file.CheckPermission(src)
    if perm == true {
        return nil, fmt.Errorf("file.CheckPermission Permission denied src: %s", src)
    }

    err = file.IsNotExistMkDir(src)
    if err != nil {
        return nil, fmt.Errorf("file.IsNotExistMkDir src: %s, err: %v", src, err)
    }

    f, err := file.Open(src + fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        return nil, fmt.Errorf("Fail to openLogFile :%v", err)
    }

    return f, nil
}
