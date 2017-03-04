package models

import "html/template"

type Article struct {
	Id        int
	Subject   string `form:"subject"`
	Body      template.HTML  `form:"body"`
	CreatedAt string
	UpdatedAt string
}
