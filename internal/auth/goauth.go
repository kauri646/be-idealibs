package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var googleOauthConfig = &oauth2.Config{
	ClientID:     "312444234427-vm0nluh6ml1l8kvjrmbudepv2i0kk3cb.apps.googleusercontent.com",
	ClientSecret: "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",
	RedirectURL:  "http://localhost:8080/auth/google/callback",
	Scopes:       []string{"profile", "email"},
	Endpoint:     google.Endpoint,
}

var store = session.New(session.Config{
	Expiration:   60,
	CookieSecure: false,
})

func GoogleOauth(ctx *fiber.Ctx) error {
	from := ctx.Query("from", "/")
	url := googleOauthConfig.AuthCodeURL(from)
	return ctx.Redirect(url, http.StatusTemporaryRedirect)
}

func Profil(ctx *fiber.Ctx) error {
	session, err := store.Get(ctx)
	if err != nil {
		panic(err)
	}

	user := session.Get("user")
	if user == nil {
		return ctx.SendStatus(401)
	}

	var userInfo map[string]interface{}
	errJson := json.Unmarshal([]byte(user.(string)), &userInfo)

	if errJson != nil {
		panic(errJson)
	}

	return ctx.JSON(fiber.Map{
		"data": userInfo,
	})
}

func GoogleCallback(ctx *fiber.Ctx) error {
	state := ctx.Query("state")
	var pathUrl string = "/"
	if state != "" {
		pathUrl = state
	}

	code := ctx.Query("code")
	if code == "" {
		ctx.SendStatus(401)
	}

	token, err := googleOauthConfig.Exchange(context.Background(), code)
	if err != nil {
		ctx.SendStatus(401)
	}

	userInfo, err := fetchGoogleUserInfo(token)
	if err != nil {
		return ctx.SendStatus(500)
	}

	// if userInfo["email"] != "kautsarrifqi1@gmail.com" {
	// 	return ctx.SendStatus(401)
	// }

	session, err := store.Get(ctx)
	if err != nil {
		panic(err)
	}

	jsonBytes, err := json.Marshal(userInfo)
	if err != nil {
		panic(err)
	}

	session.Set("user", string(jsonBytes))
	session.Save()

	// jwtToken := "asdfghjkjhgf"

	// return ctx.JSON(fiber.Map{
	// 	"status": "success",
	// 	"token":  jwtToken,
	// 	"data":   userInfo,
	// })
	return ctx.Redirect(fmt.Sprint("http://localhost:8080", pathUrl), http.StatusTemporaryRedirect)
}

func fetchGoogleUserInfo(token *oauth2.Token) (map[string]interface{}, error) {
	client := googleOauthConfig.Client(context.Background(), token)

	response, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		return nil, fmt.Errorf("failed to fetch user info %v", err)
	}

	defer response.Body.Close()

	var userInfo map[string]interface{}
	if err := json.NewDecoder(response.Body).Decode(&userInfo); err != nil {
		return nil, fmt.Errorf("failed to decode user info %v", err)
	}

	return userInfo, nil
}
