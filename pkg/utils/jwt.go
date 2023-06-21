package utils

import "github.com/golang-jwt/jwt/v5"

type Claims struct {
	Id       uint   `json:"id"`
	UserName string `json:"user_name"`
	jwt.MapClaims
}
