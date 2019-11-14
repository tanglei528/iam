package exception

var MsgFlags = map[int]string {
    Success : "success",
    SuccessCreate:  "创建成功",
    SuccessResponse: "请求成功",
    Error : "fail",
    InvalidParams : "请求参数错误",

    ErrorGetAppsFail : "获取多个应用失败",
    ErrorGetAppFail : "获取单个应用失败",
    ErrorCountAppFail  : "获取应用数量失败",
    ErrorCreateAppFail  : "创建应用失败",

    ErrorAuthCheckTokenFail : "Token鉴权失败",
    ErrorAuthCheckTokenTimeout : "Token已超时",
    ErrorAuthToken : "Token生成失败",
    ErrorAuth : "Token错误",

    ErrorConvIntToBool : "整型转布尔型错误",
}

func GetMsg(code int) string {
    msg, ok := MsgFlags[code]
    if ok {
        return msg
    }

    return MsgFlags[Error]
}
