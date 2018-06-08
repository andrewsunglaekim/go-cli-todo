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
  todo := findTodo(todoID)
  todo.Body = todoBody
  db.DB.Save(&todo)
}

func toggleTodo(todoID uint64) Todo{
  todo := findTodo(todoID)
  todo.Completed = !todo.Completed
  db.DB.Save(&todo)
  return todo
}

func deleteTodo(todoID uint64) {
  todo := findTodo(todoID)
  db.DB.Delete(&todo)
}

func findTodo(todoID uint64) Todo{
  todo := Todo{}
  db.DB.First(&todo, todoID)
  return todo
}
