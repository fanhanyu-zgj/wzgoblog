package routes

import (
	"net/http"
	"wz/app/http/controllers"
	"wz/app/http/middlewares"

	"github.com/gorilla/mux"
)

func RegisterWebRoutes(router *mux.Router) {
	// 静态页面
	pc := new(controllers.PageController)
	router.HandleFunc("/about", pc.AboutHandler).Methods("GET").Name("about")
	ac := new(controllers.ArticlesController)
	router.HandleFunc("/", ac.Index).Methods("GET").Name("home")
	router.HandleFunc("/articles/{id:[0-9]+}", ac.Show).Methods("GET").Name("articles.show")
	router.HandleFunc("/articles", ac.Index).Methods("GET").Name("articles.index")
	router.HandleFunc("/articles", middlewares.Auth(ac.Store)).Methods("POST").Name("articles.store")
	router.HandleFunc("/articles/create", middlewares.Auth(ac.Create)).Methods("GET").Name("articles.create")
	router.HandleFunc("/articles/{id:[0-9]+}/edit", middlewares.Auth(ac.Edit)).Methods("GET").Name("articles.edit")
	router.HandleFunc("/articles/{id:[0-9]+}/delete", middlewares.Auth(ac.Delete)).Methods("POST").Name("articles.delete")
	router.HandleFunc("/articles/{id:[0-9]+}", middlewares.Auth(ac.Update)).Methods("POST").Name("articles.update")

	router.NotFoundHandler = http.HandlerFunc(pc.NotFoundHandler)

	// router.PathPrefix("/css/").Handler(http.FileServer(http.Dir("./public")))
	// router.PathPrefix("/js/").Handler(http.FileServer(http.Dir("./public")))

	// 中间件：强制内容类型为 HTML
	//router.Use(middlewares.ForceHTMLMiddleware)

	// 用户认证
	auc := new(controllers.AuthController)
	router.HandleFunc("/auth/register", middlewares.Guest(auc.Register)).Methods("GET").Name("auth.register")
	router.HandleFunc("/auth/do-register", middlewares.Guest(auc.DoRegister)).Methods("POST").Name("auth.doregister")
	router.HandleFunc("/auth/login", middlewares.Guest(auc.Login)).Methods("GET").Name("auth.login")
	router.HandleFunc("/auth/dologin", middlewares.Guest(auc.DoLogin)).Methods("POST").Name("auth.dologin")
	router.HandleFunc("/auth/logout", middlewares.Auth(auc.Logout)).Methods("POST").Name("auth.logout")
	uc := new(controllers.UserController)
	router.HandleFunc("/users/{id:[0-9]+}", middlewares.Auth(uc.Show)).Methods("GET").Name("users.show")
	cc := new(controllers.CategoriesController)
	router.HandleFunc("/categories/create", middlewares.Auth(cc.Create)).Methods("GET").Name("categories.create")
	router.HandleFunc("/categories/store", middlewares.Auth(cc.Store)).Methods("POST").Name("categories.store")
	router.HandleFunc("/categories/{id:[0-9]+}", middlewares.Auth(cc.Show)).Methods("GET").Name("categories.show")

	// 开启会话
	router.Use(middlewares.StartSession)
}
