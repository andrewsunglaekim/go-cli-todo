package main

import "todo/db"

func getAllTodos() []Todo {
	todos := []Todo{}
	db.DB.Find(&todos)
	return todos
}

func createTodo(todoBody string) {
  todo := Todo{Body: todoBody, Completed: false}
  db.DB.Create(&todo)
}

func editTodo(todoID uint64, todoBody string) {
  todo := Todo{}
  db.DB.First(&todo, todoID)
  todo.Body = todoBody
  db.DB.Save(&todo)
}
