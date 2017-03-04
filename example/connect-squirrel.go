package main

import (
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"database/sql"
	"os"
	_ "github.com/go-sql-driver/mysql"
)

type Article struct {
	Id          int
	Title       string
	Description string
	CraetedAt   string
	UpdatedAt   string
}

/**
確認方法

mysql -uroot blog -e 'select * from articles;'
 */
func main() {
	var db *sql.DB

	db, err := sql.Open("mysql", "root:@/blog")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var article = Article{}
	err = sq.Select("*").From("blogs").Where("id = 2").Limit(1).
		RunWith(db).QueryRow().
		Scan(&article.Id, &article.Title, &article.Description, &article.CraetedAt, &article.UpdatedAt)
	if err != nil {
		panic(err)
	}
	fmt.Printf("id = %d/title = %s\n", article.Id, article.Title)

	id, subject := 0, ""
	err = sq.Select("blogs.id, subject").From("blogs").
		LeftJoin("articles on blogs.id = articles.id").
		Where("blogs.id = 1").Limit(1).
		RunWith(db).QueryRow().
		Scan(&id, &subject)
	if err != nil {
		panic(err)
	}
	fmt.Printf("id = %d/subject = %s\n", id, subject)

	db.Close()
}