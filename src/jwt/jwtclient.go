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

func IsValidTokenRequest(r *http.Request) bool {
  token, err := jwt.ParseFromRequest(r, func(token *jwt.Token) ([]byte, error) {
    return publicKey, nil
  })
  if token.Valid {
  	return true
  } else {
    return false
  }
}