package main

import "fmt"

type cetak interface {
	getNama() string
	getNIM() string
}

type mahasiswa struct {
	nama string
	nim  string
}

func (m mahasiswa) getNama() string {
	return m.nama
}

func (m mahasiswa) getNIM() string {
	return m.nim
}

func cobaInterface(n cetak) {
	fmt.Println(n)
}

func main() {
	// configs.ConnectDB()
	// e := echo.New()

	// // group routes for author
	// eAuthor := e.Group("/author")
	// eAuthor.POST("/", controllers.InsertAuthor)
	// eAuthor.GET("/:id", controllers.GetAuthorById)
	// eAuthor.GET("/search/:name", controllers.GetAuthorByName)
	// eAuthor.GET("/", controllers.GetAuthor)

	// // group routes for category
	// eCategory := e.Group("/category")
	// eCategory.POST("/", controllers.InsertCategory)
	// eCategory.GET("/", controllers.GetCategory)

	// e.Start(":8000")
	a := mahasiswa{"naufal", "01"}
	fmt.Println(a.getNIM())
	cobaInterface(a)

}
