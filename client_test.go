package databridge

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestClient_NewClient(t *testing.T) {

	config := Config{APIKey: "api-key", TopicSlug: "my-topic-slug", ConsumerID: "consumer-id"}
	client, err := NewClient(config)

	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	require.Equal(t, "api-key", client.APIKey)
	require.Equal(t, "consumer-id", client.ConsumerID)
	require.Equal(t, "my-topic-slug", client.TopicSlug)
	require.Equal(t, "https://data-bridge.flipsidecrypto.com/api/v1", client.BaseURL)
}
