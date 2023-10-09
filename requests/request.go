package requests

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

type Status string

const (
	Blocked   Status = "Blocked"
	Unblocked Status = "Unblocked"
)

type Domain struct {
	Domain string `json:"domain"`
	Status Status `json:"status"`
}

type TrustPositif struct {
	Values   []Domain `json:"values"`
	Response int      `json:"response"`
}

func ParseURI(domains []string) string {
	uri := "name="
	sep := "%0A"
	switch len(domains) {
	case 0:
		return uri
	case 1:
		return uri + domains[0]
	default:
		for _, domain := range domains {
			uri += domain + sep
		}
	}
	return uri
}

func NewRequest(domain string) (*TrustPositif, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	uri := fmt.Sprintf("name=%s", domain)
	payload := strings.NewReader(uri)
	req, err := createHTTPRequest(ctx, payload)
	if err != nil {
		return nil, err
	}

	recordName, err := executeHTTPRequest(req)
	if err != nil {
		return nil, err
	}

	return recordName, nil
}

func createHTTPRequest(ctx context.Context, payload *strings.Reader) (*http.Request, error) {
	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		"https://trustpositif.kominfo.go.id/Rest_server/getrecordsname_home",
		payload,
	)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	return req, nil
}

func executeHTTPRequest(req *http.Request) (*TrustPositif, error) {
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		if err == context.DeadlineExceeded {
			return nil, fmt.Errorf("request timed out")
		}
		return nil, err
	}

	var recordName TrustPositif
	if err := decodeJSON(res, &recordName); err != nil {
		return nil, err
	}

	return &recordName, nil
}

func decodeJSON(res *http.Response, v interface{}) error {
	defer res.Body.Close()
	return json.NewDecoder(res.Body).Decode(v)
}
