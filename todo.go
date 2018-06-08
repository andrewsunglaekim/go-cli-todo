package main

import (
  "github.com/jinzhu/gorm"
)

type Todo struct {
  gorm.Model
  Body string `gorm:"default:'galeone'"`
  Completed bool
}
