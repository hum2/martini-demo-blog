package handlers

import (
	"github.com/go-martini/martini"
	"fmt"
)

func GetBookRender(p martini.Params) {
	fmt.Printf("id = %s\n", p["id"])
}

func NewBookRender() {
	fmt.Println("New Book")
}
