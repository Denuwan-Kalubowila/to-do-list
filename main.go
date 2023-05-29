package main

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	query "github.com/go-to-do/controllers"
	"github.com/joho/godotenv"
)

func main() {
	var err error
	err = godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	err = query.GetAllStudent()
	if err != nil {
		fmt.Println(err)
	}
	err = query.GetStudentbyId(2)
	if err != nil {
		fmt.Println(err)
	}
	err = query.UpdateById(4, "Geesara", "Imal", 23)
	if err != nil {
		fmt.Println(err)
	}
}
