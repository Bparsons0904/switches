package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"switches/database"
	"switches/models"
	"switches/services"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

func UserLogoutCallback(c *fiber.Ctx) error {
	sessionID := c.Cookies("sessionID")
	c.ClearCookie("sessionID")
	log.Println("User logout", sessionID)
	if sessionID != "" {
		retrievedSession, err := database.GetJSONKeyDB[services.Session]("session", sessionID)
		if err != nil {
			log.Println("Error getting session from keydb", err)
			c.ClearCookie("sessionID")
			return c.Redirect("/")
		}

		retrievedSession.IsLoggedIn = false
		err = database.SetJSONKeyDB("session", sessionID, retrievedSession, 15*time.Minute)
		if err != nil {
			log.Println("Error setting session in keydb", err)
		}
		err = database.DeleteUUIDKeyDB("user", retrievedSession.UserID)
		if err != nil {
			log.Println("Error deleting user from keydb", err)
		}
	}
	return c.Redirect("/")
}

func AuthLogin(c *fiber.Ctx) error {
	auth, err := services.GenerateAuthString()
	if err != nil {
		log.Println("Error generating auth string", err)
		return c.Status(fiber.StatusInternalServerError).SendString("Error generating auth string")
	}

	authURL := fmt.Sprintf(
		"https://%s/oauth/v2/authorize?client_id=%s&redirect_uri=%s&response_type=code&scope=%s&state=%s&nonce=%s&code_challenge=%s&code_challenge_method=S256",
		auth.AuthURL,
		auth.ClientID,
		url.QueryEscape(auth.CallbackURL),
		url.QueryEscape("openid profile email offline_access"),
		auth.State,
		auth.Nonce,
		auth.CodeChallenge,
	)

	return c.Redirect(authURL)
}

func UserLogout(c *fiber.Ctx) error {
	baseLogoutURL := fmt.Sprintf("https://%s/oidc/v1/end_session", viper.GetString("AUTH_URL"))
	params := url.Values{}
	params.Add("client_id", viper.GetString("AUTH_CLIENT_ID"))
	params.Add(
		"post_logout_redirect_uri",
		fmt.Sprintf("%s/auth/logout/callback", viper.GetString("BASE_URL")),
	)
	state, err := services.GenerateRandomString(32)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).
			SendString("Error generating random string")
	}
	params.Add("state", state)
	logoutURL := fmt.Sprintf("%s?%s", baseLogoutURL, params.Encode())

	return c.Redirect(logoutURL)
}

func AuthCallback(c *fiber.Ctx) error {
	log.Println("Get auth callback")
	code := c.Query("code")
	state := c.Query("state")

	auth, err := services.GetAuth(state)
	if err != nil {
		log.Println("Error getting auth")
		return c.Status(fiber.StatusInternalServerError).SendString("Error getting auth")
	}

	tokenURL := fmt.Sprintf("https://%s/oauth/v2/token", auth.AuthURL)
	resp, err := http.PostForm(tokenURL, url.Values{
		"grant_type":    {"authorization_code"},
		"client_id":     {auth.ClientID},
		"code":          {code},
		"redirect_uri":  {auth.CallbackURL},
		"code_verifier": {auth.CodeVerifier},
	})
	if err != nil {
		log.Println("Error getting token", err)
		return err
	}
	defer resp.Body.Close()

	var tokenResponse services.TokenResponse
	if err := json.NewDecoder(resp.Body).Decode(&tokenResponse); err != nil {
		log.Println("Error decoding token response", err)
		return err
	}

	claims, err := services.VerifyIDToken(tokenResponse.IDToken, false)
	if err != nil {
		log.Println("Error verifying ID token w/o force", err)
		claims, err = services.VerifyIDToken(tokenResponse.IDToken, true)
		if err != nil {
			log.Println("Error verifying ID token, /w force", err)
			return err
		}
	}

	sub, ok := claims["sub"].(string)
	if !ok {
		log.Println("Error getting sub claim from claims")
		return fmt.Errorf("invalid sub claim")
	}

	subInt, err := strconv.Atoi(sub)
	if err != nil {
		log.Println("Error converting sub to int", err)
		return err
	}

	userClaims := &services.Claims{
		Sub:               subInt,
		Email:             claims["email"].(string),
		EmailVerified:     claims["email_verified"].(bool),
		FamilyName:        claims["family_name"].(string),
		Name:              claims["name"].(string),
		PreferredUsername: claims["preferred_username"].(string),
		GivenName:         claims["given_name"].(string),
	}

	var user models.User
	err = database.DB.First(&user, "sub = ?", userClaims.Sub).Error

	if err == gorm.ErrRecordNotFound {
		log.Println("User does not exist, creating user")
		if err := createUser(&user, userClaims); err != nil {
			log.Println("Error creating user", err)
			return err
		}
	} else if err != nil {
		log.Println("Error getting user", err)
		return err
	}

	session := services.Session{
		AccessToken:  tokenResponse.AccessToken,
		UserID:       user.ID,
		Sub:          user.Sub,
		ExpiresAt:    time.Now().Add(time.Duration(tokenResponse.ExpiresIn) * time.Second),
		RefreshToken: tokenResponse.RefreshToken,
		RefreshBy:    time.Now().Add(30 * 24 * time.Hour), // 30 days
		IDToken:      tokenResponse.IDToken,
		IsLoggedIn:   true,
	}

	if err := database.SetJSONKeyDB("session", session.AccessToken, session, 30*24*time.Hour); err != nil {
		log.Println("Error setting session in keydb", err)
	}

	expiredIn := time.Now().Add(time.Duration(tokenResponse.ExpiresIn) * time.Second)
	cookie := &fiber.Cookie{
		Name:     "sessionID",
		Value:    session.AccessToken,
		Expires:  expiredIn,
		HTTPOnly: true,
		Secure:   true,
		SameSite: fiber.CookieSameSiteNoneMode,
	}

	c.Cookie(cookie)
	return c.Redirect("/")
}

func createUser(user *models.User, claims *services.Claims) error {
	user.Sub = claims.Sub
	user.Email = claims.Email
	user.EmailVerified = claims.EmailVerified
	user.FamilyName = claims.FamilyName
	user.Name = claims.Name
	user.PreferredUsername = claims.PreferredUsername
	user.GivenName = claims.GivenName

	if err := database.DB.Create(&user).Error; err != nil {
		log.Println("Error creating user", err)
		return err
	}

	return nil
}
