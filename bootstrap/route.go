package bootstrap

import (
	"embed"
	"io/fs"
	"net/http"
	"wz/pkg/route"
	"wz/routes"

	"github.com/gorilla/mux"
)

func SetupRoute(staticFS embed.FS) *mux.Router {
	router := mux.NewRouter()
	routes.RegisterWebRoutes(router)

	route.SetRoute(router)
	//静态资源
	sub, _ := fs.Sub(staticFS, "public")
	router.PathPrefix("/").Handler(http.FileServer(http.FS(sub)))

	return router
}
