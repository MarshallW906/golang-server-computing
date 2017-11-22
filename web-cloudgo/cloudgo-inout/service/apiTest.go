package service

import (
	"math/rand"
	"net/http"

	"github.com/unrolled/render"
)

func apiTestHandler(formatter *render.Render) http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		formatter.JSON(w, http.StatusOK, struct {
			RandToken int `json:"randtoken"`
		}{RandToken: rand.Intn(2000)})
	}
}
