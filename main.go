package main

import (
  // import "github.com/gin-gonic/gin"
  "todo/db"
  "fmt"
  // "github.com/jinzhu/gorm"
  // _ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
  if err := db.Open(); err != nil {
    fmt.Println(err)
  }
  defer db.Close()
  clearScreen()
  fmt.Println("Hello! Welcome to cli todo, brought to you by gorm and go")
  makeBuffer()
  var menuChoice string
  for menuChoice != "q" {
    printMenu()
    menuChoice = getInput("What would you like to do?(Please type a number and hit enter type 'q' to quit)")
    clearScreen()
    switch menuChoice {
    case "1":
      displayAllTodos()
      advanceWithEnter()
    case "2":
      createTodoEvents()
      advanceWithEnter()
    case "3":
      editTodoEvents()
      advanceWithEnter()
    case "4":
      toggleCompletedEvents()
      advanceWithEnter()
    case "5":
      deleteTodoEvents()
      advanceWithEnter()
    case "q":
      fmt.Println("Thanks for using! Goodbye!")
    default:
      fmt.Println("Sorry that's not an option! Please choose again or \"q\" to quit")
    }

  }

  // fmt.Println("boopawerfa")
  // fmt.Println("boopawerfa")
  // fmt.Println("boopawerfa")
  // fmt.Println("boopawerfa")
  // fmt.Println(add(1, 32))
  // fmt.Println(add(1, 32))
  // fmt.Println("boopawerfa")
	// r := gin.Default()
	// r.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "pong",
	// 	})
	// })
	// r.GET("/boop", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "boopy",
	// 	})
	// })
	// r.GET("/tom", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "ewsasasa",
	// 	})
	// })
	// r.Run() // listen and serve on 0.0.0.0:8080
}

// func db(){
//   db, _ := gorm.Open("postgres", "host=localhost port=5432 user=andrewkim dbname=todo password=sandhorsebobanus sslmode=disable")
//   return db
// }
