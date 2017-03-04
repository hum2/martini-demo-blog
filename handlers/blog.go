package handlers

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"database/sql"
	"fmt"
	. "../models"
	. "../common"
	"net/http"
	"github.com/gorilla/sessions"
)

type BlogListModel struct {
	Title string
	Blogs []Blog
}

type BlogIndexModel struct {
	Title    string
	Blog     Blog
	Articles []Article
}

type BlogDetailModel struct {
	Title       string
	Blog        Blog
	Article     Article
	IsAuthorize bool
}

type BlogLoginModel struct {
	Title     string
	BlogId    string
	ArticleId string
	IsError   bool
	Message   string
}

type BlogEditModel struct {
	Title   string
	Blog    Blog
	Article Article
}

/**
	List blog
 */
func BlogListRender(r render.Render, db *sql.DB) {
	model := BlogListModel{}
	model.Title = "Blog List"
	model.Blogs = GetBlogs(db)

	r.HTML(200, "blog", model)
}

/**
	Blog top page
 */
func BlogIndexRender(p martini.Params, r render.Render, db *sql.DB) {
	var model BlogIndexModel
	model.Blog = GetBlog(p["blog_id"], db)
	model.Articles = GetArticles(model.Blog.Id, db)
	r.HTML(200, "blog/index", model)
}

/**

 */
func BlogDetailRender(p martini.Params, r render.Render, db *sql.DB) {
	var model BlogDetailModel
	blog := GetBlog(p["blog_id"], db)
	blog.BlogId = p["blog_id"]
	model.Blog = blog
	article := GetArticle(blog.Id, p["article_id"], db)
	model.Article = article
	r.HTML(200, "blog/detail", model)
}

func BlogLoginViewRender(p martini.Params, r render.Render, db *sql.DB) {
	var model BlogLoginModel
	blog := GetBlog(p["blog_id"], db)
	model.Title = blog.Title
	model.BlogId = p["blog_id"]
	model.ArticleId = p["article_id"]
	r.HTML(200, "blog/login", model)
}

func BlogLoginRender(p martini.Params, r render.Render, R *http.Request, w http.ResponseWriter, db *sql.DB) {
	var model BlogLoginModel
	authorized := IsAuthorizeBlog(p["blog_id"], R.PostFormValue("password"), db)
	fmt.Printf("authorized = %v\n", authorized)
	if authorized {
		store := sessions.NewCookieStore([]byte("secret123"))
		sess, err := store.Get(R, "session-name")
		fmt.Printf("sess = %v/err = %v\n", sess, err)
		if err != nil {
			r.Redirect(fmt.Sprintf("/%s/login/%s", p["blog_id"], p["article_id"]))
		}
		sess.Values["blog_id"] = p["blog_id"]
		sess.Save(R, w)
		fmt.Printf("blog_id = %s, session blog_id = %s\n", p["blog_id"], sess.Values["blog_id"])
		r.Redirect(fmt.Sprintf("/%s/", p["blog_id"]))
	}

	fmt.Println("HOGE")
	blog := GetBlog(p["blog_id"], db)
	model.Title = blog.Title
	model.BlogId = p["blog_id"]
	model.ArticleId = p["article_id"]
	model.IsError = true
	model.Message = "ID/PWが間違っています"
	fmt.Printf("is_error = %v, message = %s\n", model.IsError, model.Message)
	r.HTML(200, "blog/login", model)
}

/**
	edit article
 */
func BlogEditRender(p martini.Params, r render.Render, db *sql.DB) {
	var model = BlogEditModel{}
	if p["article_id"] != "" {
		blog := GetBlog(p["blog_id"], db)
		article := GetArticle(blog.Id, p["article_id"], db)
		model.Title = blog.Title
		model.Blog = blog
		model.Article = article
	}
	r.HTML(200, "blog/edit", model)
}

/**
	Process insert OR update article
 */
func BlogPostRender(p martini.Params, r render.Render, db *sql.DB, article Article) {
	blog := GetBlog(p["blog_id"], db)

	if p["article_id"] != "" && p["article_id"] != "0" {
		fmt.Printf("UPDATE articles SET subject = %s, body = %s WHERE blog_id = %v AND id = %s\n",
			article.Subject, article.Body, blog.Id, p["article_id"])
		db.Exec(
			"UPDATE articles SET subject = ?, body = ? WHERE blog_id = ? AND id = ?",
			article.Subject, article.Body, 3, p["article_id"])
		r.Redirect(fmt.Sprintf("/blog/%s/edit/%s", p["blog_id"], p["article_id"]))
	}
	res, err := db.Exec(
		"INSERT INTO articles(blog_id, subject, body) VALUES(?, ?, ?)",
		blog.Id, article.Subject, article.Body)
	PanicIf(err)
	article_id, err := res.LastInsertId()
	PanicIf(err)
	fmt.Printf("insert article id = %d\n", article_id)
	r.Redirect(fmt.Sprintf("/blog/%s/edit/%d", p["blog_id"], article_id))
}