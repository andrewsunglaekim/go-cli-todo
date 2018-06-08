package main

import (
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/postgres"
  "fmt"
)

func seed(){
  db, err := gorm.Open("postgres", "host=localhost port=5432 user=andrewkim dbname=todo password=sandhorsebobanus sslmode=disable")
  fmt.Println(err)
  defer db.Close()
  // var todos []Todo
  db.DropTableIfExists(&Todo{})
  db.CreateTable(&Todo{})
  // db.AutoMigrate(&Todo{})
  todo := Todo{Completed: true, Body: "laern how to do fuckin defaults x 1"}
  todo1 := Todo{Completed: true, Body: "laern how to do fuckin defaults x 2"}
  todo2 := Todo{Completed: true, Body: "laern how to do fuckin defaults x 3"}
  todo3 := Todo{Completed: true, Body: ""}
  // fmt.Println(db.NewRecord(todo))
  db.Create(&todo)
  db.Create(&todo1)
  db.Create(&todo2)
  db.Create(&todo3)
  // fmt.Println(db.NewRecord(todo))
}
