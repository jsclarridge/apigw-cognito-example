package api

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func CallAPI(appURL, accessToken string) (string, error) {
	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", appURL, nil)
	if err != nil {
		return "", err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", accessToken))
	resp, err := httpClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
