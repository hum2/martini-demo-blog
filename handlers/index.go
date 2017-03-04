package handlers

import (
	"github.com/martini-contrib/render"
	. "../common"
	. "../models"
	"database/sql"
	"fmt"
)

type IndexViewModel struct {
	Title       string
	Description string
}

func IndexRender(r render.Render, db *sql.DB) {
	var article = Article{}
	var q = "select id, title, description, created_at, updated_at from blogs limit 1"
	err := db.QueryRow(q).Scan(
		&article.Id, &article.Subject, &article.Body, &article.CreatedAt, &article.UpdatedAt)
	PanicIf(err)
	fmt.Printf("article.id = %d\n", article.Id)

	viewModel := IndexViewModel{
		"Martini Demo",
		"This is martini demo pages.",
	}

	r.HTML(200, "index", viewModel)
}