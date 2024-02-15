package emailer

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"

	"github.com/brettski/go-ipwatch/config"
)

type PostmarkPayload struct {
	From     string `json:"From"`
	To       string `json:"To"`
	Subject  string `json:"Subject"`
	TextBody string `json:"TextBody"`
	Tag      string `json:"Tag"`
}

var appConfig config.IpWatchConfig
var isVerbose bool = false

func SendIpChangeEmail(ipAddr net.IP, verbose bool) {
	isVerbose = verbose
	appConfig = config.GetConfig()
	isValid, result := appConfig.ValidatePostmark()
	if !isValid {
		log.Fatalf("Required env values missing to send email: %s\n", result)
	}
	msg := fmt.Sprintf("A new IP address has been discovered at %s. %s\n\nNo DNS updates have been made", appConfig.Postmark.TestLocation, ipAddr)
	postmarkPayload := PostmarkPayload{
		From:     appConfig.Postmark.EmailFrom,
		To:       appConfig.Postmark.EmailTo,
		Subject:  fmt.Sprintf("New IP address discovered for site: %s", appConfig.Postmark.TestLocation),
		TextBody: msg,
		Tag:      "New_IP_Discovered",
	}

	err := sendPostmarkEmail(postmarkPayload)
	if err != nil {
		log.Fatalf("SendIpChangeEmail: Error sending postmark email: %s", err)
	}
}

func sendPostmarkEmail(pmValues PostmarkPayload) error {
	payload, err := json.Marshal(pmValues)
	if err != nil {
		return err
	}
	if isVerbose {
		log.Printf("sendPostmarkEmail payload json: %+s", payload)
	}

	req, err := http.NewRequest(http.MethodPost, "https://api.postmarkapp.com/email", bytes.NewReader(payload))
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("X-Postmark-Server-Token", appConfig.Postmark.Token)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if isVerbose {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Printf("Error reading postmark response body: %s", err)
		}
		log.Printf("Postmark response body: %s", body)
	}

	return nil
}
