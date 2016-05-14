package jwt
import (
   jwt "github.com/dgrijalva/jwt-go"
   "io/ioutil"
   "net/http"
   "log"
)

var (
  publicKey []byte
)

func ClientInit(keyFile string) {
  publicKey, _ = ioutil.ReadFile(keyFile)
}

func IsValidTokenRequest(r *http.Request) (bool, *jwt.Token) {
  token, err := jwt.ParseFromRequest(r, func(*jwt.Token) (interface{}, error) {
    return publicKey, nil
  })
  if err != nil {
    log.Printf("Error %s", err)
  } else {
    if token.Valid {
      return true, token
    } else {
      return false, token
    }
  }
  return false, nil

}