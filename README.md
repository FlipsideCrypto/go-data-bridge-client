# go-data-bridge-client
Go client for accessing Data Bridge

https://data-bridge-docs.flipsidecrypto.com/#section/Introduction

## Usage

### Initialization
```
config := Config{APIKey: "api-key", TopicSlug: "my-topic-slug"}
client, err := NewClient(config)
```

### Get Registered Consumers 
```
consumers, err := client.GetRegisteredConsumers()
```

### Get Available Consumers 
```
consumers, err := client.GetAvailableConsumers()
```

### Register Consumer
```
consumer, err := client.RegisterConsumer()
```

### Get Next Record
```
record, err := client.GetNextRecord(consumerID)
```

### Mark Record Completed
```
err := client.CompleteRecord(record)
```

### Mark Record Failed
```
err := client.FailRecord(record)
```

### Publish New Record 
```
err := client.PublishRecord(data)
```

