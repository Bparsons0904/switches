package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"switches/services"
	"switches/templates/pages"

	"github.com/gofiber/fiber/v2"
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
	json.NewDecoder(resp.Body).Decode(&tokenResponse)

	claims, err := services.VerifyIDToken(tokenResponse.IDToken)
	if err != nil {
		log.Println("Error verifying ID token", err)
		return err
	}

	log.Println("claims: ", claims)
	// TODO
	// Createa struct for the claims
	// Check if user exists
	// Create user if user doesn exisit
	return c.Redirect("/")
	return Render(pages.HomePage(), pages.Home())(c)
	// return Render(pages.SwitchesPage(), pages.Switches())(c)
}
