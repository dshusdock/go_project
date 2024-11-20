package jwtauthsvc

import (
	"fmt"
	"github.com/go-chi/jwtauth/v5"
)

// var secretKey = []byte("secret-key")

var tokenAuth *jwtauth.JWTAuth


func init() {
	fmt.Println("Init Token service")
	tokenAuth = jwtauth.New("HS256", []byte("somethingotherthansecrett"), nil, /*jwt.WithAcceptableSkew(30*time.Second)*/)

}

func GetToken() *jwtauth.JWTAuth {
	return tokenAuth
}

func CreateToken(username string) (string, error) {
	// _, tokenString, _ := tokenAuth.Encode(map[string]interface{}{"user_id": username, "exp": time.Now().Add(300*time.Second).Unix()})
	_, tokenString, _ := tokenAuth.Encode(map[string]interface{}{"user_id": username})	
	return tokenString, nil
}


