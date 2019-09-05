package request

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

func Post(url string, data []byte) ([]byte, error) {
	client := &http.Client{}

	var err error
	req, err := http.NewRequest("POST", url, bytes.NewReader(data))

	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Xeniro-Api", "0.0.1")

	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	defer func() {
		_ = resp.Body.Close()
	}()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	return body, err
}
