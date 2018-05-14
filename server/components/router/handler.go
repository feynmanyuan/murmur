package router

import "net/http"

type Handler interface {
    Get(ctx *Context)
}

type PostHandler interface {
    Get(ctx *Context)

    Post(ctx *Context)
}

type RESTHandler interface {
    Get(ctx *Context)

    Update(ctx *Context)

    Delete(ctx *Context)

    Put(ctx *Context)

    Patch(ctx *Context)

    Options(ctx *Context)
}

type Context struct {
    Request *http.Request

    Response http.ResponseWriter

    Params map[string]string
}
