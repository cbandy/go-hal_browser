Go HAL-browser
==============

A Go [http.Handler](http://golang.org/pkg/net/http/#Handler) that serves up
a statically compiled version of Mike Kelly's
[HAL-browser](https://github.com/mikekelly/hal-browser).

Examples
========

```go
package main

import "net/http"
import "github.com/cbandy/go-hal_browser"

func main() {
	// To serve the HAL-browser at "/browser.html"
	http.Handle("/", hal_browser.Handler)
}
```
```go
	// To serve the HAL-browser at "/explore.html"
	http.Handle("/", hal_browser.At("/explore.html"))
```
```go
	// To serve the HAL-browser at "/explorer/browser.html"
	http.Handle("/explorer/", http.StripPrefix("/explorer", hal_browser.Handler))
```
```go
	// To serve the HAL-browser at "/explorer/"
	http.Handle("/explorer/", http.StripPrefix("/explorer", hal_browser.At("/")))
```
