package inmobile

import (
	"fmt"
	"github.com/messagebird/sachet"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type Config struct {
	APIKey string `yaml:"api_key"`
}

var _ (sachet.Provider) = (*InMobile)(nil)

type InMobile struct {
	Config
	HTTPClient *http.Client // The HTTP client to send requests on.
}

func NewInMobile(config Config) *InMobile {
	return &InMobile{
		config,
		&http.Client{Timeout: time.Second * 20},
	}
}

func (im *InMobile) Send(message sachet.Message) error {
	smsURL := "https://mm.inmobile.dk/Api/V2/Get/SendMessages"
	queryParams := url.Values{
		"apiKey":     {im.APIKey},
		"flash":      {"false"},
		"sendername": {message.From},
		"recipients": {strings.Join(message.To[:], ",")},
		"text":       {message.Text},
		"encoding":   {"utf-8"},
	}

	request, err := http.NewRequest("GET", smsURL, nil)
	if err != nil {
		return err
	}

	request.URL.RawQuery = queryParams.Encode()
	response, err := im.HTTPClient.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode >= http.StatusBadRequest {
		return fmt.Errorf("Failed sending sms. statusCode: %d", response.StatusCode)
	}

	return nil
}
