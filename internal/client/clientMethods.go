package client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) API(domain, namespace, identifier, output string) (Substances, error) {
	// url := commonURL + compound + "name/hydrogen/JSON"

	url := fmt.Sprintf("%s/%s/%s/%s/%s", commonURL, domain, namespace, identifier, output)

	fmt.Println(url)

	if data, ok := c.cache.Get(url); ok {
		substances := Substances{}
		err := json.Unmarshal(data, &substances)
		if err != nil {
			return Substances{}, err
		}
		return substances, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Substances{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return Substances{}, err
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return Substances{}, err
	}

	substances := Substances{}
	if err := json.Unmarshal(data, &substances); err != nil {
		return Substances{}, err
	}

	c.cache.Add(url, data)
	return substances, nil
}
