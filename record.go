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

func getUnreadCount(c Client) (*int32, error) {
	url := fmt.Sprintf("%s/topics/%s?api_key=%s", c.BaseURL, c.TopicSlug, c.APIKey)
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("content-type", "application/json")
	req.Header.Add("api_key", c.APIKey)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("error making databridge http request for %s", url))
	}

	if res.StatusCode != 200 {
		return nil, errors.Wrap(err, fmt.Sprintf("databridge - error getting unread count for %s", url))
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, errors.Wrap(err, "error trying to read databridge response body")
	}

	var js struct {
		Count int32 `json:"unread_count"`
	}

	err = json.Unmarshal([]byte(body), &js)
	if err != nil {
		return nil, errors.Wrap(err, "error trying to unmarshal response body to json")
	}

	return &js.Count, nil
}

func getNextRecord(c Client) (*Record, error) {
	count, err := getUnreadCount(c)
	if err != nil {
		return nil, err
	} else if *count == 0 {
		return nil, nil
	}

	url := fmt.Sprintf("%s/topics/%s/records/next?consumer_id=%s&api_key=%s", c.BaseURL, c.TopicSlug, c.ConsumerID, c.APIKey)
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("content-type", "application/json")
	req.Header.Add("api_key", c.APIKey)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != 200 {
		return nil, errors.New(fmt.Sprintf("databridge responded with non-200 for %s", url))
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

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	if res.StatusCode != 200 {
		return errors.New(fmt.Sprintf("databridge update record state responded with non-200 for %s", url))
	}

	return nil
}
