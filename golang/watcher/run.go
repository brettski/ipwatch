package watcher

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"

	"github.com/brettski/go-ipwatch/config"
)

type WatcherEndpoint struct {
	Host         string `json:"host"`
	ForwardFor   string `json:"x-forwarded-for"`
	ForwardProto string `json:"x-forwarded-proto"`
	Forwarded    string `json:"forwarded"`
}

var appConfig config.IpWatchConfig
var isVerbose bool = false

func RunIpWatcherCheck(verbose bool) (net.IP, error) {
	isVerbose = verbose
	appConfig = config.GetConfig()
	if appConfig.CheckEndpoint == "" {
		return nil, errors.New("config.CheckEndpoint not set, unable to continue")
	}
	if isVerbose {
		log.Printf("Configured endpoint: %s\n", appConfig.CheckEndpoint)
	}
	watcherResult, err := callWatcherEndpoint()
	if err != nil {
		return nil, err
	}
	ip := parseIp(watcherResult)
	if ip == nil {
		return nil, errors.New("No IP parsed from ForwardedFor")
	}

	return ip, nil
}

// get endpoint and parse body
func callWatcherEndpoint() (WatcherEndpoint, error) {
	req, err := http.NewRequest(http.MethodGet, appConfig.CheckEndpoint, nil)
	if err != nil { //log.Fatalf("Error call NewRequest: %v", err)
		return WatcherEndpoint{}, fmt.Errorf("Error call NewRequest: %w", err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return WatcherEndpoint{}, fmt.Errorf("Error call Do: %w", err)

	}

	if isVerbose {
		log.Printf("Response struct: %+v\n", res)
	}
	resBody, err := io.ReadAll(res.Body)

	if err != nil {
		log.Fatalf("Error reading res body: %v", resBody)
		return WatcherEndpoint{}, fmt.Errorf("Error reading resBody: %w", err)
	}
	defer res.Body.Close()
	if isVerbose {
		log.Printf("Parsed response body: %s\n", resBody)
	}

	var newWatcherEndpoint WatcherEndpoint
	err = json.Unmarshal(resBody, &newWatcherEndpoint)
	if isVerbose {
		log.Printf("WatcherEndpoint struct: %+v\n\n", newWatcherEndpoint)
	}

	return newWatcherEndpoint, nil
}

// Pulls ip from struct and returns it as a net.IP type
func parseIp(we WatcherEndpoint) net.IP {
	return net.ParseIP(we.ForwardFor)
}
