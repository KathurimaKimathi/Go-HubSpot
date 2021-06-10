package hapikey

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/sirupsen/logrus"
)

const (
	APIKey     = "<YOUR API KEY>"
	hubSpotURL = "https://api.hubapi.com/crm/v3/%s?%s"
)

// CreateContact creates a new Hubspot contact
func CreateCRMContact() error {
	logrus.Print("API KEY: ", APIKey)
	contactsReqData := url.Values{}
	contactsReqData.Set("hapikey", APIKey)

	createCRMContactsURL := fmt.Sprintf(
		hubSpotURL,
		"objects/contacts",
		contactsReqData.Encode(),
	)
	contact := map[string]interface{}{
		"properties": map[string]string{
			"firstname": "Go",
			"lastname":  "Contact",
			"phone":     "+254700000000",
			"website":   "savannahinformatics.com",
		},
	}
	bs, err := json.Marshal(contact)
	if err != nil {
		return err
	}
	contactBuffer := bytes.NewBuffer(bs)

	request, err := http.NewRequest(
		http.MethodPost,
		createCRMContactsURL,
		contactBuffer,
	)
	if err != nil {
		return err
	}
	request.Header.Set("accept", "application/json")
	request.Header.Set("content-type", "application/json")

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}

	logrus.Print(string(body))
	return nil
}
