package models

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	. "../common"
	"fmt"
	"github.com/gorilla/sessions"
	"net/http"
	"html/template"
	"github.com/martini-contrib/render"
	"github.com/go-martini/martini"
	"crypto/sha1"
	"encoding/base64"
)

/**
確認方法

mysql -uroot blog -e 'select * from articles;'
 */
func SetDatabase() *sql.DB {
	var db *sql.DB

	db, err := sql.Open("mysql", "root:@/blog")
	PanicIf(err)

	return db
}

/**

 */
func IsAuthorizeBlog(id string, pass string, db *sql.DB) (bool) {
	fmt.Printf("pass = %s\n", GeneratePassword([]byte(pass)))
	authorized := 0
	db.QueryRow("SELECT 1 FROM blogs WHERE blog_id = ? AND password = ? LIMIT 1", id, GeneratePassword([]byte(pass))).
		Scan(&authorized)
	return authorized == 1
}

/**
	Get single blog
 */
func GetBlog(id string, db *sql.DB) (b Blog) {
	db.QueryRow("SELECT id, title, description, created_at, updated_at  FROM blogs WHERE blog_id = ? LIMIT 1", id).
		Scan(&b.Id, &b.Title, &b.Description, &b.CreatedAt, &b.UpdatedAt)
	return b
}

/**
	Get blog list
 */
func GetBlogs(db *sql.DB) (blogs []Blog) {
	var blog Blog
	rows, err := db.Query("SELECT blog_id, title, description FROM blogs")
	PanicIf(err)
	defer rows.Close()
	for rows.Next() {
		blog = Blog{}
		err = rows.Scan(&blog.BlogId, &blog.Title, &blog.Description)
		PanicIf(err)
		blogs = append(blogs, blog)
	}

	return blogs
}

/**
	Get article row
 */
func GetArticle(blog_id int, article_id string, db *sql.DB) (article Article) {
	var body string
	db.QueryRow("SELECT id, subject, body, created_at, updated_at FROM articles WHERE blog_id = ? AND id = ? LIMIT 1", blog_id, article_id).
		Scan(&article.Id, &article.Subject, &body, &article.CreatedAt, &article.UpdatedAt)
	article.Body = template.HTML(body)

	return article
}

/**
	Get article list
 */
func GetArticles(blog_id int, db *sql.DB) (articles []Article) {
	var body string
	var article Article
	rows, err := db.Query("SELECT id, subject, body, created_at, updated_at FROM articles WHERE blog_id = ?", blog_id)
	PanicIf(err)
	for rows.Next() {
		article = Article{}
		err = rows.Scan(&article.Id, &article.Subject, &body, &article.CreatedAt, &article.UpdatedAt)
		PanicIf(err)
		article.Body = template.HTML(body)
		articles = append(articles, article)
	}

	return articles
}

type LoginModel struct {
	Title  string
	BlogId string
}

func LoginAuthorize(p martini.Params, r render.Render, R *http.Request) {
	model := LoginModel{}
	model.Title = "hoge"
	model.BlogId = p["blog_id"]

	store := sessions.NewCookieStore([]byte("secret123"))
	sess, err := store.Get(R, "session-name")
	if err != nil || sess.Values["blog_id"] == nil || p["blog_id"] != sess.Values["blog_id"] {
		r.Redirect(fmt.Sprintf("/%s/login/%s", p["blog_id"], p["article_id"]))
	}
}

func showLoginPage(r render.Render, model LoginModel) {
	r.HTML(200, "blog/login", model)
}

func GeneratePassword(pw []byte) (string) {
	hasher := sha1.New()
	hasher.Write(pw)
	sha := base64.URLEncoding.EncodeToString(hasher.Sum(nil))

	return sha
}
