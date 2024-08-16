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

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GetAuthCallback(c *fiber.Ctx) error {
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

	var tokenResponse struct {
		AccessToken  string `json:"access_token"`
		IDToken      string `json:"id_token"`
		RefreshToken string `json:"refresh_token"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&tokenResponse); err != nil {
		log.Println("Error decoding token response", err)
		return err
	}

	claims, err := services.VerifyIDToken(tokenResponse.IDToken)
	if err != nil {
		log.Println("Error verifying ID token", err)
		return err
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

	log.Println("User claims", userClaims)

	DB := database.GetDatabase()
	var user models.User
	err = DB.First(&user, "sub = ?", userClaims.Sub).Error
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

	log.Println("User", user)

	return c.Redirect("/")
}

func createUser(user *models.User, claims *Claims) error {
	DB := database.GetDatabase()

	user.Sub = claims.Sub
	user.Email = claims.Email
	user.EmailVerified = claims.EmailVerified
	user.FamilyName = claims.FamilyName
	user.Name = claims.Name
	user.PreferredUsername = claims.PreferredUsername
	user.GivenName = claims.GivenName

	if err := DB.Create(&user).Error; err != nil {
		log.Println("Error creating user", err)
		return err
	}

	return nil
}

type Claims struct {
	Sub               int    `json:"sub"`
	Email             string `json:"email"`
	EmailVerified     bool   `json:"email_verified"`
	FamilyName        string `json:"family_name"`
	Name              string `json:"name"`
	PreferredUsername string `json:"preferred_username"`
	GivenName         string `json:"given_name"`
}
