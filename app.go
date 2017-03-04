package main

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	. "./handlers"
	. "./models"
	"net/http"
	"github.com/martini-contrib/binding"
)

func getLogger() martini.Handler {
	return func(res http.ResponseWriter, req *http.Request, c martini.Context) {
		//fmt.Println("Logger")
	}
}
func newMartini() *martini.ClassicMartini {
	// 基本的には martini.Classic を実行しているのと同じ
	// いくつか独自
	r := martini.NewRouter()
	m := martini.New()
	// ミドルウェア
	m.Use(getLogger())

	// c.Next() でhttpリクエスト処理を先に行い、実行時間を計測する
	m.Use(func(c martini.Context) {
		//s := MakeTimestamp()
		c.Next()
		//e := MakeTimestamp() - s
		//fmt.Printf("process time = %v\n", e)
	})
	// サービス
	m.Map(SetDatabase())
	m.Use(martini.Static("public"))
	m.MapTo(r, (*martini.Routes)(nil))
	m.Action(r.Handle)

	return &martini.ClassicMartini{m, r}
}

func main() {
	m := newMartini()

	m.Use(render.Renderer(render.Options{
		Directory: "templates",
		Layout: "layouts/default",
		Extensions: []string{".tmpl"},
		Charset: "utf-8",
	}))

	m.NotFound(func(r render.Render) {
//		r.Redirect("/")
	})

	m.Get("/", BlogListRender)
	//m.Get("/about/?", AboutRender)
	//// 認証などを挟んでから AboutRender へ
	//m.Get("/about/:name", authorize, AboutRender)
	// 認証を挟みつつ book のまとまったグループへ渡す
	//m.Get("/blog", BlogListRender)
	m.Group("/:blog_id", func(r martini.Router) {
		// TODO
		r.Get("/", BlogIndexRender)
		r.Group("/login/(?P<article_id>[0-9]+)", func(r martini.Router) {
			r.Get("/?", BlogLoginViewRender)
			r.Post("/?", BlogLoginRender)
		})
		r.Group("/(?P<article_id>[0-9]+)", func(r martini.Router) {
			r.Get("/?", BlogDetailRender)
			r.Group("/edit", func(r martini.Router) {
				r.Get("/?", BlogEditRender)
				r.Post("/?", binding.Form(Article{}), BlogPostRender);
			}, binding.Form(Blog{}), LoginAuthorize)
		})

		//r.Get("/profile", BlogProfileRender)
		//r.Get("/:post_id", BlogPostRender)
		//r.Get("/:tag_name", BlogTagRender)
	})
	m.Run()
}
