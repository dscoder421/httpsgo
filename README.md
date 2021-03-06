# httpsgo
Go package to make beginners easily start a GoLang server, and even fulfil the needs of a expert which is built only using the <a href="https://pkg.go.dev/net/http">http</a> package bundled with go

NOTE: The dot before the `"github.com/..."` will make sure you don't add package name before the function code i.e `httpsgo.StartHttpServer(":3000")` and it is optional. Remove if you don't like it. 

This is a repository for making servers easily like this:

```golang
package main

import (
  "net/http"
  ."github.com/dscoder421/httpsgo"
)

func main() {
  RegisterHandler("/hello", func(rw http.ResponseWriter, r *http.Request) {
    Send("Hello", rw)
  })
  StarthttpServer(":3000")
}
```
This code will run a server at `http://localhost:3000/hello` with a response `Hello`
And you can send response codes like this:
```golang
package main

import (
  "net/http"
  ."github.com/dscoder421/httpsgo"
)

func main() {
  RegisterHandler("/hello", func(rw http.ResponseWriter, r *http.Request) {
    SendStatus(200, rw) // OK
  })
  StartHttpServer(":3000")
}
```
This code will return the response of HTTP Code `200`

Setting headers like this:
```golang
package main

import (
  "net/http"
  ."github.com/dscoder421/httpsgo"
)

func main() {
  RegisterHandler("/hello", func(rw http.ResponseWriter, r *http.Request) {
    SetHeader("Content-Type", "application/json", rw)
  })
  StartHttpServer(":3000")
}
```
This code will set the response as JSON. Not only these features but you can also start a HTTPS server(currently under development):
```golang
package main

import (
  "net/http"
  ."github.com/dscoder421/httpsgo"
)

func main() {
  RegisterHandler("/hello", func(rw http.ResponseWriter, r *http.Request) {
    Send("Hello", rw)
  })
  StartHttpsServer(":3000", "cert", "key")
}
```
This code will only accept requests from `https://` links

More features will be added soon.
