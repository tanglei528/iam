package exception

const (
    Success = 200
    SuccessCreate = 201
    SuccessResponse = 204

    Error = 500
    InvalidParams = 400

    ErrorAuthCheckTokenFail = 10001
    ErrorAuthCheckTokenTimeout = 10002
    ErrorAuthToken = 10003
    ErrorAuth = 10004

    ErrorCountAppFail = 10011
    ErrorGetAppsFail  = 10012
    ErrorGetAppFail   = 10013
    ErrorCreateAppFail = 100014

    ErrorConvIntToBool  = 50001
)
