# go-data-bridge-client
Go client for accessing Data Bridge

## Usage

### Initialization
```
config := Config{APIKey: "api-key", TopicSlug: "my-topic-slug", ConsumerID: "consumer-id"}
client, err := NewClient(config)
```

### Get Next Record
```
record, err := client.GetNextRecord()
```

### Mark Record Completed
```
err := client.CompleteRecord(record)
```

### Mark Record Failed
```
err := client.FailRecord(record)
```



