package webportal

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// AppRouter is the application router instance
var AppRouter *mux.Router

// RunWebPortal starts the server and runs the webapp
func RunWebPortal(addr string) error {
	AppRouter = mux.NewRouter()

	AppRouter.HandleFunc("/", rootHandler)
	return http.ListenAndServe(addr, AppRouter)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Webportal running with gorilla mux %s", r.RemoteAddr)
}
