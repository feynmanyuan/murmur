package handler

import (
    "net/http"
    "fmt"
    "strings"
    "bytemurmur.com/server/components/router"
)

type DashboardHandler struct {
}

func (d *DashboardHandler) Get(ctx *router.Context) {
    w, r := ctx.Response, ctx.Request
    r.ParseForm()
    fmt.Println(r.Form)
    fmt.Println("path", r.URL.Path)
    fmt.Println("scheme", r.URL.Scheme)
    fmt.Println(r.Form["url_long"])
    for k, v := range r.Form {
        fmt.Println("key:", k)
        fmt.Println("val:", strings.Join(v, ""))
    }
    fmt.Fprintf(w, "Hello fey!")
}

func (d *DashboardHandler) Post(ctx *router.Context) {
}

func DashboardIndex(w http.ResponseWriter, r *http.Request) {
    r.ParseForm()
    fmt.Println(r.Form)
    fmt.Println("path", r.URL.Path)
    fmt.Println("scheme", r.URL.Scheme)
    fmt.Println(r.Form["url_long"])
    for k, v := range r.Form {
        fmt.Println("key:", k)
        fmt.Println("val:", strings.Join(v, ""))
    }
    fmt.Fprintf(w, "Hello astaxie!")
}
