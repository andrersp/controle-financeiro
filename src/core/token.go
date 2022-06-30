package core

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// CreateToken Make a jwt token
func CreateToken(userID uint) (string, error) {
	roles := jwt.MapClaims{}

	roles["authorized"] = true
	roles["exp"] = time.Now().Add(time.Hour * 6).Unix()
	roles["userID"] = userID
	// secret
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, roles)

	return token.SignedString([]byte(SECRET_KEY))

}

func ValidateToken(r *http.Request) error {

	tokenStr := getToken(r)

	token, erro := jwt.Parse(tokenStr, verifyTokenKey)
	if erro != nil {
		return erro
	}

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return nil
	}

	// fmt.Println(token)
	return errors.New("Invalid Token")

}

func getToken(r *http.Request) string {
	token := r.Header.Get("Authorization")

	if len(strings.Split(token, " ")) == 2 {
		return strings.Split(token, " ")[1]
	}
	return ""
}

func verifyTokenKey(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("Token not valid! %v", token.Header["alg"])
	}
	return SECRET_KEY, nil
}

func getIdFromToken(roles jwt.MapClaims) (uint64, error) {
	userId, erro := strconv.ParseUint(fmt.Sprintf("%.0f", roles["userID"]), 10, 64)

	if erro != nil {
		return 0, erro
	}
	return userId, nil

}

func GetUserID(r *http.Request) (uint64, error) {
	tokenStr := getToken(r)

	token, erro := jwt.Parse(tokenStr, verifyTokenKey)
	if erro != nil {
		return 0, erro
	}

	if roles, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userID, erro := getIdFromToken(roles)
		if erro != nil {
			return 0, erro
		}
		return userID, nil

	}

	return 0, errors.New("Invalid Token")

}

func CheckTokenRequest(reqID, tokenID uint64) error {

	if reqID != tokenID {
		return errors.New("User id and request id dont macth")
	}
	return nil
}
