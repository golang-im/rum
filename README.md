# Rum 
[WIP]

Rum provides middlewares for Golang [http.Client](https://golang.org/pkg/net/http/#Client). As it is easy to use http handler(server-side) middlewares  with [Gin](https://github.com/gin-gonic/gin), Rum aims to make using client-side middlewares as easy as Gin when use http.Client.

[![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white&style=flat-square)](https://pkg.go.dev/github.com/YouEclipse/rum)
[![](https://travis-ci.org/YouEclipse/rum.svg?branch=master&style=flat-square)](https://travis-ci.org/YouEclipse/rum) 



## Install

#### Requirments
Go version 1.13+

#### Install
```
go get github.com/YouEclipse/rum
```


## Quick start
```golang



import (
	"net/http"
	"github.com/YouEclipse/rum/pkg/rum"
)

func main() {
	httpClient := http.Client{}
	transport := &rum.Transport{}
	m := NewAuthenticationMiddleware()
	transport.Use(m.BasicAuth)

    httpClient.Transport = transport
    ...
}



```


## Features v0.1.0

#### built-in middlewares
- [x] cache 
- [ ] mock 
- [ ] throttle 
- [ ] authentication 
- [ ] retry
- [ ] open-tracing

...


## License
[Apache 2.0](./LICENSE)

