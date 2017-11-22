package service

import (
	"crypto/md5"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/unrolled/render"
)

func homeHandler(formatter *render.Render) http.HandlerFunc {
	crutime := time.Now().Unix()
	h := md5.New()
	io.WriteString(h, strconv.FormatInt(crutime, 10))
	token := fmt.Sprintf("%x", h.Sum(nil))

	return func(w http.ResponseWriter, req *http.Request) {
		formatter.HTML(w, http.StatusOK, "index", struct {
			ID      string `json:"id"`
			Content string `json:"content"`
			Token   string `json:"token"`
		}{ID: "8675309", Content: "Hello from Go! of homeHandler", Token: token})
	}
}
