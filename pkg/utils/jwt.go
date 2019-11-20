package utils

import (
    "github.com/dgrijalva/jwt-go"
    "github.com/gin-gonic/gin"
    e "iam/pkg/exception"
    "iam/pkg/logging"
    "net/http"
    "time"
)

//https://segmentfault.com/a/1190000013297828

var jwtSecret = []byte("iam secret")

type CustomClaims struct {
    UserName string `json:"username"`
    Password string `json:"password"`
    jwt.StandardClaims
}

// 生成token
func GenerateToken(username, password string) (string, error) {
    claims := CustomClaims{
        username,
        password,
        jwt.StandardClaims{
            ExpiresAt: time.Now().Add(time.Hour * 12).Unix(),
            //ExpiresAt:  time.Now().Add(time.Second * 60).Unix(),
        },
    }
    tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    token, err := tokenClaims.SignedString(jwtSecret)

    return token, err
}

// 解析token
func ParseToken(token string) (*CustomClaims, error) {
    tokenClaims, err := jwt.ParseWithClaims(token, &CustomClaims{}, func(token *jwt.Token) (i interface{}, e error) {
        return jwtSecret, nil
    })

    if tokenClaims != nil {
        //if claims, ok := tokenClaims.Claims.(*CustomClaims); ok && tokenClaims.Valid {
        if claims, ok := tokenClaims.Claims.(*CustomClaims); ok {
            return claims, nil
        }
    }
    return nil, err
}

// token认证
func JWT() gin.HandlerFunc {
    return func(c *gin.Context) {
        var code int
        var data interface{}

        code = e.Success
        token := c.GetHeader("token")
        if token == "" {
            logging.Error("Not found token in header")
            code = e.ErrorAuthTokenProvide
        } else {
            code = CheckToken(token)
        }

        if code != e.Success {
            c.JSON(http.StatusUnauthorized, gin.H{
                "code" : http.StatusUnauthorized,
                "msg" : e.GetMsg(code),
                "data" : data,
            })
            c.Abort()
            return
        }
        c.Next()
    }
}

func CheckToken(token string) int {
    code := e.Success
    claims, err := ParseToken(token)
    if err != nil {
        logging.Error(e.GetMsg(e.ErrorAuthCheckTokenFail))
        code = e.ErrorAuthCheckTokenFail
    } else if time.Now().Unix() > claims.ExpiresAt {
        logging.Warn(e.GetMsg(e.ErrorAuthCheckTokenTimeout))
        code = e.ErrorAuthCheckTokenTimeout
    }
    return code
}