package sslmate_cert_search_api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func search[T any](any T, query *Query, api *SSLMateSearchAPI) error {
	url, _ := url.Parse(SSLMATE_API_URL)
	urlQuery := url.Query()

	urlQuery.Add("domain", query.Domain)
	urlQuery.Add("include_subdomains", fmt.Sprintf("%t", query.IncludeSubdomains))
	urlQuery.Add("match_wildcards", fmt.Sprintf("%t", query.MatchWildcards))
	urlQuery.Add("after", query.After)
	urlQuery.Add("expand", query.Expand)
	url.RawQuery = urlQuery.Encode()

	fmt.Printf("url: %v\n", url.String())

	req := http.Request{
		Method: http.MethodGet,
		URL:    url,
		Header: http.Header{
			"Content-Type": []string{"application/json"},
		},
	}

	if api.Token != "" {
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", api.Token))
	}

	client := new(http.Client)
	resp, err := client.Do(&req)
	if err != nil {
		return fmt.Errorf("%v: %v", ERROR_FAILED_TO_FETCH, err)
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("%v: %v", ERROR_SERVER_RETURN_NOT_200, resp.StatusCode)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("%v: %v", ERROR_FAILED_TO_READ_BODY, err)
	}

	err = json.Unmarshal(body, any)
	if err != nil {
		return fmt.Errorf("%v: %v", ERROR_FAILED_TO_MARSHAL_JSON, err)
	}
	return err
}

func (api *SSLMateSearchAPI) Search(query *Query) ([]SSLMateCertEntry, error) {
	var result []SSLMateCertEntry
	var json []SSLMateCertJson

	query.Expand = "cert_der"

	err := search(&json, query, api)
	if err != nil {
		return nil, err
	}

	for _, j := range json {
		cert, err := convertJsonToEntry(&j)

		if err != nil {
			return nil, err
		}

		result = append(result, *cert)
	}

	return result, nil
}
