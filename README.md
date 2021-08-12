# Go Error Handling Service

Mencoba implementasi error handling untuk service seperti REST API dan GRPC.

## Context

- Cara handling error http json dengan mudah


## Optionals

- [Custom Handler](#custom-handler)

### Custom Handler

Branch : [custom handler](https://github.com/zeihanaulia/go-error-handling/tree/02-custom-handler)

Membuat custom handler untuk intrupsi error dan response.

```go

type customHandler func(http.ResponseWriter, *http.Request) error
func (fn customHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    // handling error
}

```

usage:

```go
func indexHandler(w http.ResponseWriter, r *http.Request) error {return nil}

http.Handle("/",  customHandler(indexHandler)})
```

