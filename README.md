# My Pushbullet helper for golang

## 1. install [pushbullet-go](https://github.com/mitsuse/pushbullet-go) and this helper

```
$ go get -u github.com/mitsuse/pushbullet-go/...
$ go get -u github.com/meinside/pushbullet-go-helper/...
```

## 2. import this helper

```go
// pbhelper_test.go

package main

import (
	"fmt"
	"github.com/meinside/pushbullet-go-helper"
)
```

## 3. send pushes

```go
func main() {
	if pbhelper.SendNote("This is a note", "From Golang, through Pushbullet") {
		fmt.Println("Push note was successful.")
	}
	
	if pbhelper.SendLink("This is a link", "From Golang, through Pushbullet", "http://www.golang.org") {
		fmt.Println("Push link was successful.")
	}
}
```

## 4. build your binary

```
$ go build -o /some/directory/pbhelper_test pbhelper_test.go
```

## 5. put your Pushbullet access token in a file

Token file must be named `.pushbullet.token` and placed in the same directory as your binary file.

```
$ vi /some/directory/.pushbullet.token
```

## 6. run

```
$ /some/directory/pbhelper_test

Push note was successful.
Push link was successful.
```

