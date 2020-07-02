package main

import (
	"log"
	"net/rpc"
)

//Todo ..
type Todo struct {
	Title string
	Body  string
}

func main() {
	var reply Todo
	var replies []Todo

	client, err := rpc.DialHTTP("tcp", "localhost:4040")
	if err != nil {
		log.Fatal("Error dialing the rpc server", err)
	}

	// Adding some sample todos
	a := Todo{Title: "A", Body: "A item"}
	b := Todo{Title: "B", Body: "B item"}
	c := Todo{Title: "C", Body: "C item"}

	client.Call("API.AddTodo", a, &reply)
	client.Call("API.AddTodo", b, &reply)
	client.Call("API.AddTodo", c, &reply)

	// Get all todos
	log.Println("Getting all todos")
	if err = client.Call("API.GetTodos", "", &replies); err == nil {
		log.Println("Got all todos: ", replies)
	}

	// Get todo by a title
	log.Println("Getting a todo by title")
	if err = client.Call("API.GetTodoByTitle", c, &reply); err == nil {
		log.Println("Got todo item by title: ", reply)
	}

	// Editing an item
	log.Println("Editing a todo")
	b.Body = "B item is updated"
	if err = client.Call("API.EditTodo", b, &reply); err == nil {
		log.Println("Updated todo item: ", reply)
	}

	// Deleting an item
	log.Println("Deleting a todo")
	if err = client.Call("API.DeleteTodo", c, &reply); err == nil {
		log.Println("Deleted todo item: ", reply)
	}

}
