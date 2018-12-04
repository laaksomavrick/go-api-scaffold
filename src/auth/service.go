package auth

import (
	"fmt"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/laaksomavrick/goals-api/src/user"
)

// TODO: real hmac secret; read from file
// TODO: should only need to read this once on app start, not every req?
// if keyData, e := ioutil.ReadFile("test/hmacTestKey"); e == nil {
// 	hmacSampleSecret = keyData
// } else {
// 	panic(e)
// }

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) GenerateToken(user user.User) (string, error) {
	// TODO: store this token in db; invalidate old token if relogging in
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": user.ID,
		"email":  user.Email,
		"nbf":    time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
	})
	tokenString, err := token.SignedString([]byte("asdqweasdqwe"))
	return tokenString, err
}

func (s *Service) ParseToken(tokenString string) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return hmacSampleSecret, nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println(claims["foo"], claims["nbf"])
	} else {
		fmt.Println(err)
	}
}
