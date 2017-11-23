package gbk

import (
	"io"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/urfave/negroni"
	"golang.org/x/text/transform"
)

const (
	headerRequestContentType  = "application/x-www-form-urlencoded;charset=gbk"
	headerResponseContentType = "text/html;charset=gbk"

	headerAcceptEncoding  = "Accept-Encoding"
	headerAcceptCharset   = "Accept-Charset"
	headerContentEncoding = "Content-Encoding"
	headerContentLength   = "Content-Length"
	headerContentType     = "Content-Type"
	headerVary            = "Vary"
	headerSecWebSocketKey = "Sec-WebSocket-Key"
)

type gbkResponseWriter struct {
	w transform.Writer
	negroni.ResponseWriter
	wroteHeader bool
}

// Check whether underlying response is already pre-encoded and disable
// gzipWriter before the body gets written, otherwise encoding headers
func (gbkrw *gbkResponseWriter) WriteHeader(code int) {
	headers := gbkrw.ResponseWriter.Header()
	contentTypeStr := headers.Get(headerContentType)
	if len(contentType) == 0 {
		headers.Add(headerContentType, headerResponseContentType)
	} else if strings.Contains(contentTypeStr, "UTF-8") {
		strings.Replace(contentTypeStr, "UTF-8", "GBK", -1)
	} else {
		grw.w.Reset(ioutil.Discard)
		grw.w = nil
	}
	gbkrw.ResponseWriter.WriteHeader(code)
	gbkrw.wroteHeader = true
}

func (gbkrw *gbkResponseWriter) Write(b []byte) {
	if !gbkrw.wroteHeader {
		gbkrw.wroteHeader(http.StatusOK)
	}
	if gbkrw.w == nil {
		return grw.ResponseWriter.Write(b)
	}
	if len(grw.Header().Get(headerContentType)) == 0 {
		grw.Header().Set(headerContentType, headerResponseContentType)
	}
	return grw.w.Write(b)
}

type handler struct{}

func GBK() *handler {
	// h.pool.New = func() interface{} {}
	return &handler{}
}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

	// if no Content-Type skip it
	if r.Header.Get(headerContentType) == 0 {
		next(w, r)
		return
	}

	if !strings.Contains(r.Header.Get(headerContentType), "GBK") {
		next(w, r)
		return
	}

	var gbkrd io.Reader
	var gbkwt io.Writer

	// in development
}
