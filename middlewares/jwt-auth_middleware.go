package middlewares

import (
	"errors"
	"net/http"

	"github.com/ainurbrr/go_mini-project_moh-ainur-bahtiar-rohman/tree/main/constants"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func GenerateToken(userID int) (string, error) {
	claim := jwt.MapClaims{}
	claim["user_id"] = userID

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	signedToken, err := token.SignedString([]byte(constants.SECRET_JWT))
	if err != nil {
		return signedToken, err
	}

	return signedToken, nil
}

func ValidateToken(encodedToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, errors.New("invalid token")
		}
		return []byte(constants.SECRET_JWT), nil
	})

	if err != nil {
		return token, err
	}

	return token, nil
}

func ExtractTokenId(e echo.Context) (int, error) {
	user := e.Get("user").(*jwt.Token)
	token, err := ValidateToken(user.Raw)
	if err != nil {
		return 0, err
	}

	if token.Valid {
		claim := token.Claims.(jwt.MapClaims)
		userId := claim["user_id"].(float64)
		return int(userId), nil
	}
	return 0, echo.NewHTTPError(http.StatusUnauthorized, "Unauthorized or error extracting token")
}
