package databridge

type Config struct {
	APIKey     string
	ConsumerID string
	TopicSlug  string
}

type Client struct {
	BaseURL    string
	APIKey     string
	TopicSlug  string
	ConsumerID string
}

func NewClient(config Config) (Client, error) {
	c := Client{}
	c.APIKey = config.APIKey
	c.BaseURL = "https://data-bridge.flipsidecrypto.com/api/v1"
	c.TopicSlug = config.TopicSlug
	c.ConsumerID = config.ConsumerID

	return c, nil
}

func (c Client) GetNextRecord() (*Record, error) {
	return getNextRecord(c)
}

func (c Client) CompleteRecord(r Record) error {
	return r.updateRecordState(c, "completed")
}

func (c Client) FailRecord(r Record) error {
	return r.updateRecordState(c, "failed")
}
