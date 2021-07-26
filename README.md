# REST  API

## Server
Server runs on `localhost:8585`

## Endpoint
Produce a message on a kafka topic:<br>
`POST /topics/:topic?key=value`

### Body
The body can be any string. It will convert it to bytes and send it on the topic
#### Plain text
```text
A simple string
```

#### JSON
```json
{
  "hello":"world"
}
```
