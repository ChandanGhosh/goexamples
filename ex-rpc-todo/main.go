package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
)

//Todo ..
type Todo struct {
	Title string
	Body  string
}

// API ..
type API int

var todos []Todo

// GetTodos ..
func (a *API) GetTodos(title string, reply *[]Todo) error {
	*reply = todos
	return nil
}

// GetTodoByTitle ..
func (a *API) GetTodoByTitle(t Todo, reply *Todo) error {
	var todoItem Todo
	for _, todo := range todos {
		if todo.Title == t.Title {
			todoItem = todo
		}
	}
	*reply = todoItem
	return nil
}

// AddTodo ..
func (a *API) AddTodo(todoItem Todo, reply *Todo) error {
	todos = append(todos, todoItem)
	log.Printf("Added item: %q\n", todoItem)
	*reply = todoItem
	return nil
}

// EditTodo ..
func (a *API) EditTodo(todoItem Todo, reply *Todo) error {
	var editedTodo Todo
	for id, todo := range todos {
		if todo.Title == todoItem.Title {
			todos[id] = todoItem
			editedTodo = todos[id]
		}
	}
	log.Println("Updated todo", editedTodo)
	*reply = editedTodo
	return nil
}

// DeleteTodo ..
func (a *API) DeleteTodo(todoItem Todo, reply *Todo) error {
	var deletedTodo Todo
	for id, todo := range todos {
		if todo.Title == todoItem.Title {
			todos = append(todos[:id], todos[id+1:]...)
			deletedTodo = todoItem
		}
	}
	log.Println("Deleted todo: ", deletedTodo)
	*reply = deletedTodo
	return nil
}

func main() {
	var api = new(API)
	err := rpc.Register(api)
	if err != nil {
		log.Fatal("Error registering the API", err)
	}

	rpc.HandleHTTP()
	listener, err := net.Listen("tcp", ":4040")
	if err != nil {
		log.Fatal("Listener error", err)
	}
	log.Println("Server listening on port ", 4040)
	err = http.Serve(listener, nil)
	if err != nil {
		log.Fatal("Error creating server", err)
	}

}
