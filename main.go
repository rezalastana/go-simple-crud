package main

import (
	"github.com/labstack/echo"
	"github.com/rezalastana/golang-todos/controller"
	"github.com/rezalastana/golang-todos/database"
)

func main() {

	//panggil database
	db := database.InitDb()
	defer db.Close() //close db setelah tidak dibutuhkan

	err := db.Ping()
	if err != nil {
		panic(err)
	}

	e := echo.New() //declare echo

	// ENDPOINT
	// import controller dan func getall
	controller.NewGetAllTodosController(e,db) //parsing echo-nya dan db nya

	// import controller dan func create
	controller.NewCreateTodosController(e,db)

	// import controtller dan func update
	controller.NewUpdateTodosController(e,db)

	// import controller dan func check
	controller.NewCheckTodosController(e,db)

	//import delete
	controller.NewDeleteTodosController(e, db)


	e.Start(":8000") //port 8080

	
}