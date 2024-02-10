package handler

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/basriyasin/sp-user/repository"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

const (
	claimUserID     = "id"
	claimUserName   = "name"
	claimUserPhone  = "phone"
	claimExpire     = "exp"
	tokenExpireTime = time.Hour * 999999

	HeaderAuthorization = "Authorization"
)

type jwtClaim struct {
	repository.User
	Exp int64 `json:"exp"`
}

// generate signed JWT token with RSA key and include the user profile in the jwt claim
func generateToken(key *rsa.PrivateKey, user repository.User) (token string, err error) {
	var claim jwt.MapClaims
	c, _ := json.Marshal(jwtClaim{
		User: user,
		Exp:  time.Now().Add(tokenExpireTime).Unix(),
	})
	json.Unmarshal(c, &claim)

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	return jwtToken.SignedString(x509.MarshalPKCS1PrivateKey(key))
}

// verify JWT token and return the user information inside jwt claim
func verifyToken(ctx echo.Context, key *rsa.PrivateKey) (user repository.User, err error) {
	auth := strings.Split(ctx.Request().Header.Get(HeaderAuthorization), " ")
	if len(auth) != 2 {
		err = fmt.Errorf("invalid token")
		return
	}

	jwtToken, err := jwt.Parse(auth[1], func(token *jwt.Token) (interface{}, error) {
		return x509.MarshalPKCS1PrivateKey(key), nil
	})
	if err != nil {
		return
	}

	var claim jwtClaim
	c, _ := json.Marshal(jwtToken.Claims)
	json.Unmarshal(c, &claim)
	user = claim.User
	return
}
