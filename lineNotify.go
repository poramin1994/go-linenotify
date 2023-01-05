package lineNotify

import (
	"net/http"
	"strings"

	beego "github.com/beego/beego/v2/adapter"
)

// API for Client
type API struct {
	beego.Controller
}

const (
	lineNotifyAPI = "https://notify-api.line.me/api/notify"
)

func (api *API) LineNotify(accessToken string, data map[string]string) (err error) {
	// Create a new HTTP client
	client := &http.Client{}

	// Create a new POST request
	var messageNoti string
	messageNoti += " \n"

	for name, value := range data {
		messageNoti += name + " : " + value + " \n"
	}

	req, err := http.NewRequest("POST", lineNotifyAPI, strings.NewReader("message="+messageNoti))
	if err != nil {
		return
	}

	// Set the "Authorization" and "Content-Type" headers
	req.Header.Set("Authorization", "Bearer "+accessToken)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	// Check the response status code
	if resp.StatusCode != 200 {
		return
	}
	return
}
