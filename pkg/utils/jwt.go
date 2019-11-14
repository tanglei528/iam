package utils

import (
    "github.com/dgrijalva/jwt-go"
    "time"
)

//https://segmentfault.com/a/1190000013297828

var jwtSecret = []byte("iam secret")

type CustomClaims struct {
    UserName string `json:"username"`
    Password string `json:"password"`
    jwt.StandardClaims
}

func GenerateToken(username, password string) (string, error) {
    claims := CustomClaims{
        username,
        password,
        jwt.StandardClaims{
            //ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
            ExpiresAt:  time.Now().Add(time.Second * 60).Unix(),
        },
    }
    tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    token, err := tokenClaims.SignedString(jwtSecret)

    return token, err
}

func ParseToken(token string) (*CustomClaims, error) {
    tokenClaims, err := jwt.ParseWithClaims(token, &CustomClaims{}, func(token *jwt.Token) (i interface{}, e error) {
        return jwtSecret, nil
    })

    if tokenClaims != nil {
        if claims, ok := tokenClaims.Claims.(*CustomClaims); ok && tokenClaims.Valid {
            return claims, nil
        }
    }
    return nil, err
}

