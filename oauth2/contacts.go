package oauth2

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/sirupsen/logrus"
)

const (
	hubspotTokenURL    = "https://api.hubapi.com/oauth/v1/token"
	grantType          = "HUBSPOT_GRANT_TYPE"
	clientID           = "HUBSPOT_CLIENT_ID"
	clientSecret       = "HUBSPOT_CLIENT_SECRET"
	code               = "HUBSPOT_AUTHORIZATION_CODE"
	hubSpotAuthCodeURL = "https://app.hubspot.com/oauth/authorize?"
	client_id          = "ced2da20-a4cb-4eec-85e6-dae3d0d81bb0"
	redirect_uri       = "https://www.savannahinformatics.com/"
	scope              = "contacts%20content%20reports%20social%20automation%20forms%20files%20hubdb%20tickets"
)

type AuthorizationCode struct {
	Code string `json:"code"`
}

func GetAuthorizationCode() (*AuthorizationCode, error) {
	oauthCode := url.Values{}
	oauthCode.Set("client_id", client_id)
	oauthCode.Set("redirect_uri", redirect_uri)
	oauthCode.Set("scope", scope)

	request, err := http.NewRequest(
		http.MethodPost,
		hubSpotAuthCodeURL,
		strings.NewReader(oauthCode.Encode()),
	)
	if err != nil {
		logrus.Println("AN ERROR OCCURRED", err)
		return nil, fmt.Errorf("an error occured: %s", err)
	}

	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	return nil, nil
}
