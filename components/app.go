package components

import (
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/relentless-coder/vnotes-server/components/pages"
	"github.com/relentless-coder/vnotes-server/helpers"
)

var (
	Router *mux.Router
)

func SetupRoutes() *mux.Router {
	Router = mux.NewRouter()
	Router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		helpers.RespondWithJSON(w, http.StatusOK, "hello world")
	})
	pages.Routes(Router)
	return Router
}
