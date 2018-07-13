package common

import (
	"net/http"
	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/gorilla/context"
	"time"
	"dpm/vars"
)

var (
	SecretMap = map[string]string{
		vars.PROJECT_NAME: "welcome to dmp",
	}

	IgnoreValidateRoute = map[string]bool{
		"users_register": true,
		"users_login":    true,
	}
)

const (
	CURRENT_USER = "CUSR"
	TOKEN_KEY    = "Authorization"
)

type UserClaims struct {
	*jwt.StandardClaims
	TokenType string
	*UserToken
}

type UserToken struct {
	Id   string
	Name string
	Pwd  string
}

func CreateToken(user *UserToken) (string, error) {
	// create a signer for rsa 256
	t := jwt.New(jwt.GetSigningMethod("HS256"))
	now := time.Now()
	issuedAt := time.Date(2018, time.May, 30, 0, 0, 0, 0, now.Location())
	// set our claims
	t.Claims = &UserClaims{
		&jwt.StandardClaims{
			// set the expire time
			// see http://tools.ietf.org/html/draft-ietf-oauth-json-web-token-20#section-4.1.4
			ExpiresAt: now.Add(time.Hour * 24).Unix(), //for dev
			// ExpiresAt: now.Add(time.Minute * 10).Unix(),
			Issuer:   vars.PROJECT_NAME,
			IssuedAt: issuedAt.Unix(),
		},
		"level1",
		user,
	}
	// Creat token string
	return t.SignedString([]byte(SecretMap[vars.PROJECT_NAME]))
}

func ValidateTokenHandlerFunc(inner http.Handler, routeName string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if !IgnoreValidateRoute[routeName] {
			// only accessible with a valid token
			// Get token from request
			token, err := request.ParseFromRequestWithClaims(r, request.AuthorizationHeaderExtractor, &UserClaims{},
				func(token *jwt.Token) (interface{}, error) {
					return []byte(SecretMap[vars.PROJECT_NAME]), nil
				})
			Logger.Info("token:", token)
			// If the token is missing or invalid, return error
			if err != nil {
				error := ErrTraceCode(http.StatusUnauthorized, err)
				panic(error)
			}

			// Token is valid
			// fmt.Fprintln(w, "Welcome,", token.Claims.(*UserClaims).Name)
			u := token.Claims.(*UserClaims).UserToken
			context.Set(r, CURRENT_USER, u)
			// Got the value like this : context.Get(r,"cusr").(*UserToken)
			Logger.Infof("username is:[%s],and pwd is:[%s]", u.Name, "*********")
		}

		inner.ServeHTTP(w, r)
	})
}
