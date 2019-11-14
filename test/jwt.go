package main

import (
    "github.com/dgrijalva/jwt-go"
    "iam/pkg/utils"
    "time"
)
var (
    iamKey = []byte("My secret")
)


func getToken() string {
    claims := &jwt.StandardClaims{
        //ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
        ExpiresAt: time.Now().Add(time.Second * 10).Unix(),
    }

    // Create token with claims
    tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

    // Generate encoded token and send it as response.
    token, err := tokenClaims.SignedString(iamKey)
    if err != nil {
        panic(err)
    }
    return token
}

func main() {
    //token := getToken()
    token, _ := utils.GenerateToken("", "")
    print(token)
    claims, _ := utils.ParseToken(token)
    print(claims)
}