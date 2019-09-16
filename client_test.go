package databridge

import (
	"fmt"
	"os"
	"testing"
)

func TestClient_GetUnreadCount(t *testing.T) {
	client := getClient(t)
	c, err := client.GetUnreadCount()
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if c == nil {
		t.Fatal("count is nil")
	}

	fmt.Fprintln(os.Stdout, "GetUnreadCount")
	fmt.Fprintln(os.Stdout, *c)
}

func TestClient_GetNextRecord(t *testing.T) {
	client := getClient(t)
	r, err := client.GetNextRecord()
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if r == nil {
		t.Fatal("Result is nil")
	}

	record := string(*r)
	fmt.Fprintln(os.Stdout, "GetNextRecord")
	fmt.Fprintln(os.Stdout, record)
}

func TestClient_GetRegisteredConsumers(t *testing.T) {
	client := getClient(t)

	consumers, err := client.GetRegisteredConsumers()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	fmt.Fprintln(os.Stdout, "GetRegisteredConsumers")
	fmt.Fprintln(os.Stdout, consumers)
}

func TestClient_GetAvailableConsumers(t *testing.T) {
	client := getClient(t)

	consumers, err := client.GetAvailableConsumers()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	fmt.Fprintln(os.Stdout, "GetConsumers")
	fmt.Fprintln(os.Stdout, consumers)
}

func TestClient_RegisterConsumer(t *testing.T) {
	client := getClient(t)

	consumer, err := client.RegisterConsumer()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	fmt.Fprintln(os.Stdout, "RegisterConsumer")
	fmt.Fprintln(os.Stdout, consumer)
}

func getClient(t *testing.T) Client {
	config := Config{APIKey: "15d4fbbc-d12e-4b08-a4d6-b55b92200649", TopicSlug: "dev-alert-fcas-events", ConsumerID: "84d3cba6-8be8-4848-b759-59093f2e74e6"}
	client, err := NewClient(config)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	return client
}
