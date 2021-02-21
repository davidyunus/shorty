# shorty

How to run shorty



1. Migrate url table from migration file
2. `go run cmd/main.go` 

note: docker and test file will be updated soon

### POST /v1/shorten

```
POST /shorten
Content-Type: "application/json"

{
  "url": "https://google.com",
  "shortcode": "mbahgugel"
}
```
#### Response
```
{
    "code": "Created",
    "data": {
        "shortcode": "mbahgugel"
    }
}
```

### GET /v1/:short
#### Response
```
{
    "code": "Found",
    "data": "https://google.com"
}
```


### GET /v1/:short/stats
#### Response
```
{
    "code": "OK",
    "data": {
        "id": 5,
        "url": "https://google.com",
        "shortcode": "mbahgugel",
        "redirectCount": 0,
        "startDate": "2021-02-21T22:33:15.523016Z",
        "lastSeenDate": "2021-02-21T22:33:15.523016Z"
    }
}
```
