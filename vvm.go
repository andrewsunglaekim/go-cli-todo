package main

import (
  "fmt"
  "strconv"
)

func displayAllTodos() []Todo{
  fmt.Println("All Todos")
  makeBuffer()
  todos := getAllTodos()
  for _, todo := range todos{
    completedStatus := "Complete!"
    if todo.Completed {
      completedStatus = "Not Complete."
    }
    fmt.Println(fmt.Sprint(todo.ID) + ") " + todo.Body + " - "+ completedStatus)
  }
  return todos
}

func createTodoEvents() {
  todoText := getInput("Please enter the text of your todo:")
  createTodo(todoText)
  fmt.Println("Todo created!")
}

func editTodoEvents() {
  // todoEdit := Todo{}
  makeBuffer()
  displayAllTodos()
  makeBuffer()
  todoIDString := getInput("Type the number of the todo you would like to edit")
  todoID, err := strconv.ParseUint(todoIDString, 10, 32)
  if err !=nil{
    fmt.Println(err)
  }
  editText := getInput("Type the new text for your todo: ")
  editTodo(todoID, editText)
  fmt.Println("Todo has been edited!")
  // todoIndexPlus1
}

func toggleCompletedEvents() {
  makeBuffer()
  displayAllTodos()
  makeBuffer()
  todoIDString := getInput("Type the number of the todo you would like to edit")
  todoID, err := strconv.ParseUint(todoIDString, 10, 32)
  if err !=nil{
    fmt.Println(err)
  }
}
