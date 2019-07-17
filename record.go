package databridge

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Record represents a Data Bridge data record
type Record struct {
	ID   string                   `json:"id"`
	Data []map[string]interface{} `json:"data"`
}

func getNextRecord(c Client) (*Record, error) {
	url := fmt.Sprintf("%s/topics/%s/records/next?consumer_id=%s&api_key=%s", c.BaseURL, c.TopicSlug, c.ConsumerID, c.APIKey)

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("content-type", "application/json")
	req.Header.Add("api_key", c.APIKey)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode == 404 {
		return nil, nil
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var record Record
	err = json.Unmarshal([]byte(body), &record)
	if err != nil {
		return nil, err
	}

	return &record, nil
}

func (r Record) updateRecordState(c Client, state string) error {
	url := fmt.Sprintf("%s/records/%s/state/%s?api_key=%s", c.BaseURL, r.ID, state, c.APIKey)

	req, _ := http.NewRequest("PUT", url, nil)
	req.Header.Add("content-type", "application/json")
	req.Header.Add("api_key", c.APIKey)

	_, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	// todo - probably get more granular on checking this response

	return nil
}
