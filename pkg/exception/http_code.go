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

    ErrorCountAppRolesFail = 10111
    ErrorListAppRolesFail  = 10112
    ErrorGetAppRole  = 10113
    ErrorUpdateAppRole  = 10114

    ErrorConvIntToBool  = 50001
)
