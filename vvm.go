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
  displayTodosWithBuffer()
  todoIDString := getInput("Type the number of the todo you would like to edit")
  todoID, err := strconv.ParseUint(todoIDString, 10, 32)
  if err !=nil{
    fmt.Println(err)
  }
  editText := getInput("Type the new text for your todo: ")
  editTodo(todoID, editText)
  fmt.Println("Todo has been edited!")
}

func toggleCompletedEvents(){
  displayTodosWithBuffer()
  todoIDString := getInput("Type the number of the todo you would like to toggle(Complete/Not Complete)")
  todoID, err := strconv.ParseUint(todoIDString, 10, 32)
  if err !=nil{
    fmt.Println(err)
  }
  todo := toggleTodo(todoID)
  completedStatus := "Complete!"
  if todo.Completed {
    completedStatus = "Not Complete."
  }
  fmt.Println(todo.Body, "is now", completedStatus)
}

func deleteTodoEvents(){
  displayTodosWithBuffer()
  todoIDString := getInput("Type the number of the todo you would like to delete")
  todoID, err := strconv.ParseUint(todoIDString, 10, 32)
  if err !=nil{
    fmt.Println(err)
  }
  deleteTodo(todoID)
  fmt.Println("Todo deleted!")
}
