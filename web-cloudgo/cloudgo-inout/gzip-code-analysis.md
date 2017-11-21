# negroni-gzip code Analysis

[Negroni-gzip](https://github.com/phyber/negroni-gzip) is a Gzip middleware for [Negroni](https://github.com/urfave/negroni).

It's quite tiny as the source code has only 127 lines (109 slocs), but useful.

<!-- TOC -->

- [negroni-gzip code Analysis](#negroni-gzip-code-analysis)
  - [Summary](#summary)
  - [Constants](#constants)
  - [type gzipResponseWriter struct](#type-gzipresponsewriter-struct)
    - [func (grw *gzipResponseWriter) WriteHeader(code int)](#func-grw-gzipresponsewriter-writeheadercode-int)
    - [func (grw *gzipResponseWriter) Write(b []byte)](#func-grw-gzipresponsewriter-writeb-byte)

<!-- /TOC -->

## Summary

Generally, it simply wrapped the `compress/gzip` to make it compatible with `Negroni`, providing a `handler` with a `ServeHTTP func(http.ResponseWriter, *http.Request, http.HandlerFunc)` so that it can be treated as an `negroni.Handler` (hence become a middleware for `Negroni`).

In the `ServeHTTP` func, it first wraps the `http.ResponseWriter` to create a `negroni.ResponseWriter`, then wraps the `negroni.ResponseWriter` into `gzipResponseWriter`.  After that, it call the `next` handler supplying the `gzipResponseWriter` instead of the original.

When everything has been handled and written (after it turned out of the call hierarchy of `next(rw, r)`), the `gzipResponseWriter` deleted the content length after we know we have been written to, and `Close()`.

## Constants

The constants for compression are all copied from the compress/gzip package, simply use them to specify the compression mode when calling Gzip() to get a gzip middleware instance from external code.

Constants with the prefix `header-` and `encodingGzip = 'gzip'`, are corresponding to the HTTP Header detection.

`headerVary`: Add `Vary: Accept-Encoding` to Header, to ensure the client can get a well-encoded response (not Garbled). To know more about it, click [this](http://www.webkaka.com/blog/archives/how-to-set-Vary-Accept-Encoding-header.html)(this site is in zh-CN).

## type gzipResponseWriter struct

A wrapper of `negroni.ResponseWriter`, providing an "overriden"(it's quite similar to C++'s override **but I havn't know whether they are same things yet**) `Write()` and `WriteHeader()` (using the original `Header()`, no need to "override")

```go
type gzipResponseWriter struct {
    w *gzip.Writer
    negroni.ResponseWriter
    wroteHeader bool
}
```

### func (grw *gzipResponseWriter) WriteHeader(code int)

First it get the original Header from `grw.ResponseWriter.Header()` and check it.

If there's nothing about `Content-Encoding`, add `Content-Encoding: gzip` & `Vary: Accept-Encoding` to specify gzip compression.

[Else of above] If there's already a `Content-Encoding` (means that the response has already been encoded, so we should not encode it for another time), simply reset the `*gzip.Writer` and set it to `nil` to free it.

Then, no matter whether `gzip` is used, it should always call `grw.ResponseWriter.WriteHeader(code)` to write the header, and set the `wroteHeader` to `true`.

### func (grw *gzipResponseWriter) Write(b []byte)

First check if the Header has been written, if not, call `WriteHeader()` to write it.

Then it checks if `grw.w == nil`, if so, that means we don't need to apply gzip compression this time, so simply call `ResponseWriter.Write()` and `return`.

If not, then we use `gzip.Write()`. But we should use the net/http library to detect the `Content-Type` first if it hasn't been specified. After that we call `grw.w.Write()` and return.

## type handler struct

```go
type handler struct {
    pool sync.Pool
}
```

### func Gzip(level int) *handler

It only contains a `sync.Pool` to specify a customized `New` func, which init and return a `*gzip.Writer`, and can re-use it because of `sync.Pool.Get()` & `sync.Pool.Put()`.

For more details of pointer type, you can see [this](http://blog.csdn.net/qq_21816375/article/details/78161723)(also in zh-CN).

### func (h *handler) ServeHTTP()

First it checked if the client support gzip encoding. If not, simply skip and call `next(w, r)`.

Then check if the client attempt **WebSocket** connection. If so, should also skip immediately and call `next(w, r)`. It is because there's something wrong with the compatibility between `WebSocket` and `gzip`. The issue is controversial, but most of the tone are telling that `WebSocket` cannot work well with `gzip`.

You can see [this site on StackOverflow](https://stackoverflow.com/questions/11646680/could-websocket-support-gzip-compression) and [this zh-CN site](http://www.bijishequ.com/detail/321220?p=) to get more.

After all the checks have been passed, it simply call `h.pool.Get()` & `h.pool.Put()` to init a `*gzip.Writer`, and then wrap the `negroni.ResponseWriter` with the `*gzip.Writer` and a `wroteHeader bool` to create a `gzipResponseWriter` (can also serve as a `ResponseWriter`), call `next(w, r)`, and Finally delete the content length after we know we have been written to, then `Close()` the `*gzip.Writer`.