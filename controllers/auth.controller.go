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
	"switches/utils"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func UserLogout(c *fiber.Ctx) error {
	sessionID := c.Cookies("sessionID")
	c.ClearCookie("sessionID")
	log.Println("User logout", sessionID)
	if sessionID != "" {
		retrievedSession, err := database.GetJSONKeyDB[Session]("session", sessionID)
		if err != nil {
			log.Println("Error getting session from keydb", err)
			c.ClearCookie("sessionID")
			return c.Redirect("/")
		}

		log.Println("Deleting session", sessionID)
		database.KeyDB.Del(database.KeyDB.Context(), "session:"+sessionID)
		database.KeyDB.Del(database.KeyDB.Context(), "user:"+retrievedSession.UserID.String())

	}
	return c.Redirect("/")
}

func GetAuthCallback(c *fiber.Ctx) error {
	timer := utils.StartTimer("Get auth callback")

	log.Println("Get auth callback")
	code := c.Query("code")
	state := c.Query("state")

	auth, err := services.GetAuth(state)
	if err != nil {
		log.Println("Error getting auth")
		return c.Status(fiber.StatusInternalServerError).SendString("Error getting auth")
	}

	timer.LogTime("Get auth")

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
	timer.LogTime("Token request")
	defer resp.Body.Close()

	var tokenResponse struct {
		IDToken      string `json:"id_token"`
		AccessToken  string `json:"access_token"`
		RefreshToken string `json:"refresh_token"`
		ExpiresIn    int    `json:"expires_in"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&tokenResponse); err != nil {
		log.Println("Error decoding token response", err)
		return err
	}

	timer.LogTime("Token decode")
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

	userClaims := &Claims{
		Sub:               subInt,
		Email:             claims["email"].(string),
		EmailVerified:     claims["email_verified"].(bool),
		FamilyName:        claims["family_name"].(string),
		Name:              claims["name"].(string),
		PreferredUsername: claims["preferred_username"].(string),
		GivenName:         claims["given_name"].(string),
	}

	timer.LogTime("Get claims")

	var user models.User
	err = database.DB.First(&user, "sub = ?", userClaims.Sub).Error

	timer.LogTime("Query user")
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

	timer.LogTime("Have user")
	id, err := uuid.NewV7()
	if err != nil {
		log.Println("Error generating session id", err)
		return err
	}
	expiredAt := time.Now().Add(time.Duration(tokenResponse.ExpiresIn) * time.Second)
	session := Session{
		SessionID:    id,
		UserID:       user.ID,
		Sub:          user.Sub,
		AccessToken:  tokenResponse.AccessToken,
		RefreshToken: tokenResponse.RefreshToken,
		ExpiresAt:    expiredAt,
		IDToken:      tokenResponse.IDToken,
	}

	if err := database.SetJSONKeyDB("session", session.SessionID.String(), session); err != nil {
		log.Println("Error setting session in keydb", err)
	}
	timer.LogTime("Set session in keydb")

	timer.LogTime("Retrieved session")
	expiredIn := time.Now().Add(time.Duration(tokenResponse.ExpiresIn) * time.Second)
	cookie := &fiber.Cookie{
		Name:     "sessionID",
		Value:    session.SessionID.String(),
		Expires:  expiredIn,
		HTTPOnly: true,
		Secure:   true,
		SameSite: fiber.CookieSameSiteNoneMode,
	}

	c.Cookie(cookie)

	timer.LogTime("Set session cookie")
	return c.Redirect("/")
}

type Session struct {
	SessionID    uuid.UUID `json:"sessionId"`
	UserID       uuid.UUID `json:"userId"`
	Sub          int       `json:"sub"`
	AccessToken  string    `json:"accessToken"`
	RefreshToken string    `json:"refreshToken"`
	ExpiresAt    time.Time `json:"expiresAt"`
	IDToken      string    `json:"idToken"`
}

func createUser(user *models.User, claims *Claims) error {
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

type Claims struct {
	Sub               int      `json:"sub"`
	Email             string   `json:"email"`
	EmailVerified     bool     `json:"email_verified"`
	FamilyName        string   `json:"family_name"`
	Name              string   `json:"name"`
	PreferredUsername string   `json:"preferred_username"`
	GivenName         string   `json:"given_name"`
	Roles             []string `json:"roles"`
}
