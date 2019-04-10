package main

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "ponag",
		})
	})

	r.LoadHTMLGlob("templates/*.html")
	r.GET("/", func(c *gin.Context) {
		db, err := sql.Open("mysql", "root:@/_my_gin")
		if err != nil {
			panic(err.Error())
		}
		defer db.Close()

		rows, err := db.Query("SELECT * FROM todo")
		if err != nil {
			panic(err.Error())
		}
		defer rows.Close()

		var todos []todo
		for rows.Next() {
			var todo todo
			err = rows.Scan(&todo.Id, &todo.Title, &todo.Contents, &todo.Due)
			if err != nil {
				panic(err.Error())
			}

			todos = append(todos, todo)
		}

		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "My Gin",
			"todos": todos,
		})
	})

	r.Run()
}

type todo struct {
	Id       int
	Title    string
	Contents string
	Due      string
}
