package jwt
import (
  jwt "github.com/dgrijalva/jwt-go"
  "io/ioutil"
  "time"
)

var (
  privateKey []byte
  default_ttl int64
)

func ServerInit( key_path string, ttl int64) {
  privateKey, _ = ioutil.ReadFile(key_path)
  default_ttl = ttl
}

func CreateToken(user_id int, ttl int64) (string, *jwt.Token) {
	 token := jwt.New(jwt.GetSigningMethod("RS256"))
	 token.Claims["ID"] = user_id
	 if ttl == 0 {
	 	ttl = default_ttl
	 }
     token.Claims["exp"] = time.Now().Unix() + ttl
     tokenString, err := token.SignedString(privateKey)
     if err != nil {
     	//TODO: do something
     }
	 return tokenString, token 
}

