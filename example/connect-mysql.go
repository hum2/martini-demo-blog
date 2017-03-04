package main

import (
	"os"
	"fmt"
	"database/sql"
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
	var q = "select * from blogs limit 1"
	err = db.QueryRow(q).Scan(
		&article.Id, &article.Title, &article.Description, &article.CraetedAt, &article.UpdatedAt)
	if err != nil {
		panic(err)
	}
	fmt.Printf("id = %d/title = %s\n", article.Id, article.Title)

	rows, err := db.Query(
		"SELECT id, title, description, created_at, updated_at FROM blogs",
	)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		id, title, desc, createdAt, updatedAt := 0, "", "", "", ""
		err = rows.Scan(&id, &title, &desc, &createdAt, &updatedAt)
		if err != nil {
			panic(err)
		}
		fmt.Println(id, title, desc, createdAt, updatedAt)
	}

	db.Close()
}