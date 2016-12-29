# Go URL Mapper [![Build Status](https://travis-ci.org/assertis/url-mapper.svg?branch=master)](https://travis-ci.org/assertis/url-mapper)

Package mapping query string to struct using reflection and tags.

Example:

```go
type Request struct {
    Origin string `query:"origin"` 
    Destination string `query:"destination"`
    NumOfPassengers int `query:"adults"`
    OutwardDate time.Time `query:"outward"`
    ReturnDate time.Time `query:"inward,omitempty"`
}
```
