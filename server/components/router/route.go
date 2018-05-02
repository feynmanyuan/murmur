package router

import (
	"sync"
	"net/http"
	"strings"
	"regexp"
	"fmt"
	"net/url"
	"reflect"
	"github.com/elazarl/go-bindata-assetfs"
	"bytemurmur.com/server/components"
)

var (
	once                  sync.Once
	routerServiceInstance *router
	defaultStaticPath     = map[string]string{"/static/": "/static/"}
)

func init() {
	once.Do(func() {
		routerServiceInstance = &router{
			lock:		&sync.RWMutex{},
			handlers: 	make([]*routeRegex, 0),
		}
		fs := assetfs.AssetFS{
			Asset:     asset.Asset,
			AssetDir:  asset.AssetDir,
			AssetInfo: asset.AssetInfo,
		}
		routerServiceInstance.fsHandler = http.FileServer(&fs)
	})
}

type HTTPHandler struct {}

func (h *HTTPHandler)ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := &Context{
		Request:  r,
		Response: w,
	}

	routerServiceInstance.route(ctx)
}

type router struct {
	lock 		*sync.RWMutex
	StaticPath  map[string]string
	handlers 	[]*routeRegex
	fsHandler   http.Handler
}

func (r *router) add(rr *routeRegex) {
	r.lock.Lock()
	defer r.lock.Unlock()

	r.handlers = append(r.handlers, rr)
}

func (r *router) setStaticMapping(path string, dir string) {
	r.lock.Lock()
	defer r.lock.Unlock()

	if r.StaticPath == nil {
		r.StaticPath = make(map[string]string)
	}
	r.StaticPath[path] = dir
}

func (r *router) getStaticPath() map[string]string {
	if r.StaticPath == nil {
		return defaultStaticPath
	}
	return r.StaticPath
}

func (r *router)route(ctx *Context) {
	m := r.getStaticPath()

	var severForRequest bool
	var requestPath = ctx.Request.URL.Path

	for k, _ := range m {
		if strings.HasPrefix(requestPath, k) {
			//file := v + ctx.Request.URL.Path[len(k):]
			//log.Println(file)
			r.fsHandler.ServeHTTP(ctx.Response, ctx.Request)
			//http.ServeFile(ctx.Response, ctx.Request, file)
			severForRequest = true
			return
		}
	}

	for _, v := range r.handlers {

		if !v.regex.MatchString(requestPath) {
			continue
		}

		matches := v.regex.FindStringSubmatch(requestPath)

		if len(matches[0]) != len(requestPath) {
			continue
		}

		params := make(map[string]string)
		if len(v.params) > 0 {
			values := ctx.Request.URL.Query()

			for idx, match:= range matches[1:] {
				values.Add(v.params[idx], match)
				params[v.params[idx]] = match
			}

			ctx.Request.URL.RawPath = url.Values(values).Encode() + "&" + ctx.Request.URL.RawPath

			ctx.Params = params
		}

		ins := reflect.ValueOf(&v.handler).Elem()
		var method reflect.Value

		switch ctx.Request.Method {
		case "GET":
			method = ins.MethodByName("Get")
		case "POST":
			method = ins.MethodByName("Post")
		case "UPDATE":
			method = ins.MethodByName("Update")
		case "DELETE":
			method = ins.MethodByName("Delete")
		case "HEAD":
			method = ins.MethodByName("Head")
		case "PUT":
			method = ins.MethodByName("put")
		case "PATCH":
			method = ins.MethodByName("Patch")
		case "OPTIONS":
			method = ins.MethodByName("Options")
		}


		if method.IsValid() {
			in := make([]reflect.Value, 1)
			in[0] = reflect.ValueOf(ctx)
			method.Call(in)
			severForRequest = true
		}
		break
	}

	if !severForRequest {
		http.NotFound(ctx.Response, ctx.Request)
	}
}

type routeRegex struct {

	regex 		*regexp.Regexp

	params 		[]string

	handler 	Handler

}

func StaticMapping(path string, dir string) {
	path = fmt.Sprintf("/%s/", strings.Join(strings.Split(strings.Trim(path, "/"),"/"), "/"))
	dir = fmt.Sprintf("/%s/", strings.Join(strings.Split(strings.Trim(dir, "/"),"/"), "/"))

	routerServiceInstance.setStaticMapping(path, dir)
}

func Register(regexPath string, handler Handler) {
	pathSlice := strings.Split(regexPath, "/")

	expr := "([^/]+)"

	params := make([]string, 0)
	for i, path := range pathSlice {
		if strings.HasPrefix(path, ":") {
			index := len(path)
			if paramCutFor := strings.Index(path, "(");index != -1 {
				expr = path[paramCutFor:]
				index = paramCutFor
			}
			params = append(params, path[1:index])
			pathSlice[i] = expr
		}
	}

	pattern := strings.Join(pathSlice, "/")

	regex, regexErr := regexp.Compile(pattern)

	if regexErr != nil {
		panic(regexErr)
		return
	}


	routerServiceInstance.add(&routeRegex{
		regex:		regex,
		params: 	params,
		handler:	handler,
	})
}

