package controller

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/labstack/echo"
)

//create request
type CreateRequest struct {
	Title string `json:"title"` //tag
	Description string `json:"description"` //tag
}

//response
type TodoResponse struct {
	Id 			int 	`json:"id"`
	Title 		string 	`json:"title"`
	Description string 	`json:"description"`
	Done 		bool 	`json:"done"`
}

//request update
type UpdateRequest struct {
	Title string `json:"title"`
	Description string `json:"description"`
}

//check response
type CheckRequest struct {
	Done bool `json:"done"`
}

func NewGetAllTodosController(e *echo.Echo, db *sql.DB) {
	//routing GET data
	e.GET("/todos", func(ctx echo.Context) error {
		//query db
		rows, err := db.Query("SELECT * FROM todos")
		//error handle
		if err != nil {
			return ctx.String(http.StatusInternalServerError, err.Error())
		}

		//definisikan reponse baru di atas
		//definisikan array di golang
		var res []TodoResponse

		//looping untuk nemampilkan data rows
		for rows.Next(){
			//definisikan variable
			var id int
			var title string
			var description string
			var done int

			//mengambil data dengan pointer
			err := rows.Scan(&id, &title, &description, &done)
			if err != nil {
				return ctx.String(http.StatusInternalServerError, err.Error())
			}

			//mapping data dari struct TodoResponse
			var todo TodoResponse
			todo.Id = id
			todo.Title = title
			todo.Description = description
			// todo.Done = done ini akan error karena pada struct meminta bool namun pada db meminta int
			//perlu dirubah
			if done == 1 {
				todo.Done = true
			} else {
				todo.Done = false
			}

			//masukkan dalam done
			res = append(res, todo)
		}
		
		return ctx.JSON(http.StatusOK, res)
	})
}

func NewCreateTodosController(e *echo.Echo, db *sql.DB) {
	//routing POST data
	e.POST("/todos", func(ctx echo.Context) error {
		// mapping request body to struct

		//inisialiasasi
		var request CreateRequest
		json.NewDecoder(ctx.Request().Body).Decode(&request) //diisi oleh struct yang sudah dibuat -> CreateRequest dan dikirim sebagai pointer

		//insert into db menggunakan Exec, done = 0 berarti default valuenya adalah false
		_, err := db.Exec("INSERT INTO todos (title, description, done) VALUES (?, ?, 0)", request.Title, request.Description) // ? ? itu memparsing dengan Title dan Description dari request

		//check ada error tidak
		if err != nil {
			// return http error
			return ctx.String(http.StatusInternalServerError, err.Error())
		}

		return ctx.String(http.StatusOK, "OK")
	})	
}

func NewUpdateTodosController(e *echo.Echo, db *sql.DB) {
	//routing UPDATE data
	e.PATCH("/todos/:id", func(ctx echo.Context) error {
		//ambil id
		id := ctx.Param("id")

		//inisialiasasi
		var request UpdateRequest
		json.NewDecoder(ctx.Request().Body).Decode(&request) //diisi oleh struct yang sudah dibuat -> CreateRequest dan dikirim sebagai pointer

		//insert into db menggunakan Exec, done = 0 berarti default valuenya adalah false
		_, err := db.Exec("UPDATE todos SET title = ?, description = ? WHERE id = ?", request.Title, request.Description, id) // karena ada 3 tanda ? yang dikirimkan maka di dikirimkan dengan 3 variable juga dari request dan id

		//check ada error tidak
		if err != nil {
			// return http error
			return ctx.String(http.StatusInternalServerError, err.Error())
		}

		return ctx.String(http.StatusOK, "OK")
	})
}

func NewCheckTodosController(e *echo.Echo, db *sql.DB) {
	//routing endpoint check unchek
	e.PATCH("/todos/:id/check", func(ctx echo.Context) error {
		//ambil id
		id := ctx.Param("id")

		//inisialiasasi
		var request CheckRequest
		json.NewDecoder(ctx.Request().Body).Decode(&request) //diisi oleh struct yang sudah dibuat -> CreateRequest dan dikirim sebagai pointer

		//perlu di mapping ulang karena data dari db int dan dari request berupa bool
		var doneInt int
		if request.Done {
			doneInt = 1
		}

		//insert into db menggunakan Exec, done = 0 berarti default valuenya adalah false
		_, err := db.Exec("UPDATE todos SET done = ? WHERE id = ?", doneInt, id) //cukup dengan 2 variable saja done dan id 

		//check ada error tidak
		if err != nil {
			// return http error
			return ctx.String(http.StatusInternalServerError, err.Error())
		}

		return ctx.String(http.StatusOK, "OK")
	})
}

func NewDeleteTodosController(e *echo.Echo, db *sql.DB) {
//routing DELETE data
	e.DELETE("/todos/:id", func(ctx echo.Context) error {
		id := ctx.Param("id")

		//query
		_, err := db.Exec("DELETE FROM todos WHERE id = ?", id) // ? untuk memparsing id dari url

		//err
		if err != nil {
			return ctx.String(http.StatusInternalServerError, err.Error())
		}

		return ctx.String(http.StatusOK, "OK")
	})
}