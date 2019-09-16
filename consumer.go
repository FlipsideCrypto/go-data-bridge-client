package databridge

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/pkg/errors"
)

// Consumer are processes that you run to iterate over the records stored in a topic. In order to retrieve records to process, you must register a Consumer with the Data Bridge.
// Upon registration, the Data Bridge will generate a "consumer_id". You must include this "consumer_id" when ever requesting to view a topic-level record.
// Topics are a lot like append-only lists. As you traverse a list you need to keep track of your index in the list. The Data Bridge conveniently utilizes your "consumer_id" to keep track of what topic level records you've seen and have not seen.
// It is possible to register multiple consumers, per account, if you would like to process topic records in a parralell fashion.
type Consumer struct {
	ID string `json:"id"`
}

// GetRegisteredConsumers get a list of consumers that have registered with the data bridge under this api key
func (c Client) GetRegisteredConsumers() ([]Consumer, error) {
	consumers := make([]Consumer, 0)

	url := fmt.Sprintf("%s/consumers?api_key=%s", c.BaseURL, c.APIKey)
	req, _ := http.NewRequest("GET", url, nil)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return consumers, err
	}

	if res.StatusCode != 200 {
		return consumers, errors.New(fmt.Sprintf("databridge responded with non-200 for %s", url))
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return consumers, err
	}

	err = json.Unmarshal([]byte(body), &consumers)
	if err != nil {
		return consumers, err
	}

	return consumers, nil
}

// GetAvailableConsumers returns the consumers associated with this api key that are not assigned to a record.
func (c Client) GetAvailableConsumers() ([]Consumer, error) {
	consumers := make([]Consumer, 0)

	url := fmt.Sprintf("%s/consumers/available?api_key=%s", c.BaseURL, c.APIKey)
	req, _ := http.NewRequest("GET", url, nil)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return consumers, err
	}

	if res.StatusCode != 200 {
		return consumers, errors.New(fmt.Sprintf("databridge responded with non-200 for %s", url))
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return consumers, err
	}

	err = json.Unmarshal([]byte(body), &consumers)
	if err != nil {
		return consumers, err
	}

	return consumers, nil
}

// RegisterConsumer registers a consumer with the Data Bridge to use when consuming topic-level records.
func (c Client) RegisterConsumer() (*Consumer, error) {
	url := fmt.Sprintf("%s/consumers?api_key=%s", c.BaseURL, c.APIKey)
	req, _ := http.NewRequest("POST", url, nil)
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

	var consumer Consumer

	err = json.Unmarshal([]byte(body), &consumer)
	if err != nil {
		return nil, err
	}

	return &consumer, nil
}
