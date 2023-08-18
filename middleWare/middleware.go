package middleWare

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"

	firebase "firebase.google.com/go"
	firebaseAuth "firebase.google.com/go/auth"
	"github.com/labstack/echo/v4"
	"google.golang.org/api/option"
)

func Auth() echo.MiddlewareFunc {
	return auth
}

func AuthAllowGuest() echo.MiddlewareFunc {
	return authAllowGuest
}

func auth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		opt := option.WithCredentialsFile(os.Getenv("GOOGLE_APPLICATION_CREDENTIALS"))
		app, err := firebase.NewApp(context.Background(), nil, opt)
		if err != nil {
			fmt.Println(err.Error())
			return err
		}

		client, err := app.Auth(context.Background())
		if err != nil {
			fmt.Println(err.Error())
			return err
		}

		auth := c.Request().Header.Get("Authorization")
		idToken := strings.Replace(auth, "Bearer ", "", 1)
		token, err := client.VerifyIDToken(context.Background(), idToken)
		if err != nil {
			fmt.Println(err.Error())
			return c.JSON(http.StatusUnauthorized, err.Error())
		}

		c.Set("token", token)
		return next(c)
	}
}

func authAllowGuest(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		opt := option.WithCredentialsFile(os.Getenv("GOOGLE_APPLICATION_CREDENTIALS"))
		app, err := firebase.NewApp(context.Background(), nil, opt)
		if err != nil {
			fmt.Println(err.Error())
			return err
		}

		client, err := app.Auth(context.Background())
		if err != nil {
			fmt.Println(err.Error())
			return err
		}

		auth := c.Request().Header.Get("Authorization")
		idToken := strings.Replace(auth, "Bearer ", "", 1)
		if idToken == "undefined" {
			c.Set("token", &firebaseAuth.Token{})
			return next(c)
		}
		token, err := client.VerifyIDToken(context.Background(), idToken)
		if err != nil {
			fmt.Println(err.Error())
			return c.JSON(http.StatusUnauthorized, err.Error())
		}

		c.Set("token", token)
		return next(c)
	}
}
