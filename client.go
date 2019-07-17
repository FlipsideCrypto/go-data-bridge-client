package databridge

// Config allows a consuming app to set up API Key, Consumer ID, and Topic Slug
type Config struct {
	APIKey     string
	ConsumerID string
	TopicSlug  string
}

// Client allows access to the Databridge API
type Client struct {
	BaseURL    string
	APIKey     string
	TopicSlug  string
	ConsumerID string
}

// NewClient returns a new Databridge Client
func NewClient(config Config) (Client, error) {
	c := Client{}
	c.APIKey = config.APIKey
	c.BaseURL = "https://data-bridge.flipsidecrypto.com/api/v1"
	c.TopicSlug = config.TopicSlug
	c.ConsumerID = config.ConsumerID

	return c, nil
}

// GetNextRecord returns the topic's next record.  Will return nil without an error when there are no more records.
func (c Client) GetNextRecord() (*Record, error) {
	return getNextRecord(c)
}

// CompleteRecord allows the record to be marked as completed
func (c Client) CompleteRecord(r Record) error {
	return r.updateRecordState(c, "completed")
}

// FailRecord allows the record to be marked as failed
func (c Client) FailRecord(r Record) error {
	return r.updateRecordState(c, "failed")
}
