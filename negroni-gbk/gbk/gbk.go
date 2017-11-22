package gbk

import (
	"net/http"

	"github.com/urfave/negroni"
	"golang.org/x/text/transform"
)

const (
	headerRequestContentType = "application/x-www-form-urlencoded;charset:gbk"
	headerContentTypeCharset = "text/html;charset=gbk"

	headerAcceptEncoding  = "Accept-Encoding"
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
	// headers := gbkrw.ResponseWriter.Header()

	gbkrw.ResponseWriter.WriteHeader(code)
	gbkrw.wroteHeader = true
}

func (gbkrw *gbkResponseWriter) Write(b []byte) {

}

type handler struct {
	// pool sync.Pool
}

func GBK() *handler {
	// h.pool.New = func() interface{} {}
	return &handler{}
}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

}
