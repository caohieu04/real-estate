package main

import (
	"fmt"
	"os"
)

func main() {
	dsn := os.Getenv("DBConnectionStr")
	fmt.Print(dsn)

	// db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
}
