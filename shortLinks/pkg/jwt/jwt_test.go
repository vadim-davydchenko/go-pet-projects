package jwt_test

import (
	"shortLinks/pkg/jwt"
	"testing"
)

func TestJWTCreate(t *testing.T) {
	const email = "a@a.ru"
	jwtService := jwt.NewJWT("80+gUP7R2QpY7qEUA5HPq/Hjk3oTYN3+7ClcWvbpx0E=")
	token, err := jwtService.Create(jwt.JWTData{
		Email: email,
	})

	if err != nil {
		t.Fatal(err)
	}
	isValid, data := jwtService.Parse(token)

	if !isValid {
		t.Fatal("Invalid token")
	}
	if data.Email != email {
		t.Fatal("Invalid data")
	}
}
