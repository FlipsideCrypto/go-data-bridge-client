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
