package upbit

import (
	jwt "github.com/dgrijalva/jwt-go"
	"net/url"
	"time"
	"log"
)

type JWTtoken struct {
	AccessKey string `json:"access_key"`
	Nonce  int64    `json:"nonce"`
	Query string	`json:"query"`
	jwt.StandardClaims
}

func NewJWTtoken(queries map[string]string, accessKey string) *JWTtoken{
	v := url.Values{}
	if queries != nil{
		for key,value := range queries {
			v.Add(key, value)
		}
		query := v.Encode()
		return &JWTtoken{AccessKey:accessKey,Nonce:time.Now().Unix(),Query:query}
	}
	return &JWTtoken{AccessKey:accessKey,Nonce:time.Now().Unix()}
}


func (jwtToken *JWTtoken)CreateTokenString(secretKey string) (string, error) {
	// Embed User information to `token`
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), jwtToken)
	// token -> string. Only server knows this secret (foobar).
	tokenstring, err := token.SignedString([]byte(secretKey))

	if err != nil {
		log.Fatalln(err)
		return "", err
	}

	return tokenstring, nil
}