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
	record, err := client.GetNextRecord("b053b974-608e-4f1b-9969-871a02cfbf92")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if record == nil {
		t.Fatal("record is nil")
	}

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
	config := Config{APIKey: "15d4fbbc-d12e-4b08-a4d6-b55b92200649", TopicSlug: "dev-alert-fcas-events"}
	client, err := NewClient(config)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	return client
}
