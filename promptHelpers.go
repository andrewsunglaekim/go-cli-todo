package main

import (
  "fmt"
  "bufio"
  "os"
  "strings"
)

func getInput(prompt string) string {
  reader := bufio.NewReader(os.Stdin)
  fmt.Println(prompt)
  text, _ := reader.ReadString('\n')
  return strings.TrimSpace(text)
}

func makeBuffer(){
  fmt.Println("*****************************************************************")
}


func printMenu(){
  fmt.Println("1) View All Todos")
  fmt.Println("2) Create a Todo")
  fmt.Println("3) Edit a Todo")
  fmt.Println("4) Toggle a Todo(Complete/Not Complete)")
  fmt.Println("5) Delete a Todo")
}

func clearScreen(){
  fmt.Println("\033[H\033[2J")
}

func advanceWithEnter() {
  getInput("\nPress enter to continue.")
  clearScreen()
}

func displayTodosWithBuffer() {
  makeBuffer()
  displayAllTodos()
  makeBuffer()
}
