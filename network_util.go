package utils

import (
	"io"
	"io/ioutil"
	"net/http"
)

type NetworkUtil struct {
}

func (this *NetworkUtil) HttpRequest(method string, url string, headers map[string]string, requestBody io.Reader) (responseHeader http.Header, body []byte, err error) {
	req, err := http.NewRequest(method, url, requestBody)
	if err != nil {
		return nil, nil, err
	}

	for k, v := range headers {
		req.Header.Add(k, v)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, nil, err
	}

	defer resp.Body.Close()

	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, err
	}

	return resp.Header, body, nil
}
