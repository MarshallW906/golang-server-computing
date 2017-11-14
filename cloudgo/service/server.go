package service

import (
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
)

// NewServer configures and returns a Server.
func NewServer() *negroni.Negroni {

	// use render.New to init a JsonFormatter
	formatter := render.New(render.Options{
		IndentJSON: true,
	})

	// init a negroni instance and a mux
	n := negroni.Classic()
	mx := mux.NewRouter()

	// add routes of /hello/{id} to the mux
	initRoutes(mx, formatter)

	// can be used as an http.Handler because type Router implements http.Handler interface
	n.UseHandler(mx)
	return n
}

func initRoutes(mx *mux.Router, formatter *render.Render) {
	// set a HandleFunc to the mux, supporting GET
	mx.HandleFunc("/hello/{id}", testHandler(formatter)).Methods("GET")
}

func testHandler(formatter *render.Render) http.HandlerFunc {
	// use Closure to apply a render.Render to a HandleFunc
	return func(w http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		id := vars["id"]
		formatter.JSON(w, http.StatusOK, struct{ Test string }{"Hello " + id})
	}
}
