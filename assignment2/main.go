package main

import (
	"github.com/Kaffeineman/assignment2-06/assignment2/route"

	"gorm.io/gorm"
)

var db *gorm.DB

func main() {
	r := route.SetupRouter()
	r.Run(":8080")
}
