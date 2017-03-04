package models

import "database/sql"

type Blog struct {
	Id            int
	BlogId        string `form:"blog_id"`
	Title         string
	Description   string
	CreatedAt     string
	UpdatedAt     string
	IsAuthorize   int
	AuthorizePass sql.NullString `form:"password"`
}
