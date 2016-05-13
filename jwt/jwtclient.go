package jwt
import (
   jwt "github.com/dgrijalva/jwt-go"
   "io/ioutil"
   "net/http"
   "fmt"
)

var (
  publicKey []byte
)

func ClientInit(keyFile string) {
  publicKey, _ = ioutil.ReadFile(keyFile)
}

func IsValidTokenRequest(r *http.Request) (bool, *jwt.Token) {
  token, err := jwt.ParseFromRequest(r, func(token *jwt.Token) ([]byte, error) {
    return publicKey, nil
  })
  if token.Valid {
  	return true, token
  } else {
    return false, token
  }
}