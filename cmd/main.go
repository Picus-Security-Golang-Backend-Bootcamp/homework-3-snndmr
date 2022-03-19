package main

import (
	"fmt"
	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-3-snndmr/internal/domain/book"
	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-3-snndmr/internal/helper"
	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-3-snndmr/internal/infrastructure"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func main() {
	// To initiate database and book repository object.
	db := infrastructure.NewMySqlDB("root:password123!@tcp(localhost:3306)/books")
	repository := book.NewRepository(db)
	repository.Migration()

	// To fill database with csv data.
	err := helper.FillDBFromCSV(repository, "resources/book.csv")
	if err != nil {
		return
	}

	args := os.Args

	// To check if arguments are entered.
	if len(args) == 1 {
		fmt.Printf("Parameters you can use in the %s application:\n%-10s%s\n%-10s%s", filepath.Base(args[0]), "search", "To search books by name.", "list", "To list the books.")
		return
	}

	// To operate by parameter.
	switch strings.ToLower(args[1]) {
	case "buy":
		// To check if the book id is given as an argument.
		if len(args) == 2 {
			fmt.Println("You must enter book id. (Ex: go run main.go buy 5 12)")
			return
		}

		// To check the validity of the book id.
		id, err := strconv.Atoi(args[2])
		if err != nil {
			fmt.Println("You must enter valid book id. (Ex: go run main.go buy 5 12)")
			return
		}

		// To check if the number of books is given as an argument
		if len(args) == 3 {
			fmt.Println("You must enter the number of books you want to buy. (Ex: go run main.go buy 5 12)")
			return
		}

		// To check the validity of the number of books to be purchased
		count, err := strconv.Atoi(args[3])
		if err != nil {
			fmt.Println("You must enter valid book count. (Ex: go run main.go buy 5 12)")
			return
		}

		err, result := repository.GetById(id)
		if err != nil {
			fmt.Printf("No book found with this %d.", id)
			return
		}

		err = result.DecreaseAmount(count)
		if err != nil {
			fmt.Println(err)
			return
		}

		err = repository.Update(result)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("Book purchased successfully!")
	case "delete":
		// To check if the book id is given as an argument
		if len(args) == 2 {
			fmt.Println("You must enter book id. (Ex: go run main.go delete 5)")
			return
		}

		// To check the validity of the book id
		id, err := strconv.Atoi(args[2])
		if err != nil {
			fmt.Printf("You must enter valid book id. (Ex: go run main.go delete 5)")
			return
		}

		err, result := repository.GetById(id)
		if err != nil {
			fmt.Printf("No book found with this %d.", id)
			return
		}

		err = result.Delete()
		if err != nil {
			fmt.Println(err)
			return
		}

		err = repository.Update(result)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("Book deleted successfully!")
	case "search":
		// To check if the substr is given as an argument
		if len(args) == 2 {
			fmt.Println("You must enter a value. (Ex: go run main.go search Chocolate Jesus)")
			return
		}

		// To concatenate and search disjointed strings
		bookSubstr := strings.Join(args[2:], " ")
		books := repository.Search(bookSubstr)

		if len(books) == 0 {
			fmt.Printf("%s not found.", bookSubstr)
			return
		}

		for _, result := range books {
			fmt.Println(result.ToString())
		}
	case "list":
		// To check if arguments are entered
		if len(args) > 2 {
			fmt.Println("List has no additional arguments (Ex: go run main.go list)")
			return
		}

		books := repository.List()
		for _, result := range books {
			fmt.Println(result.ToStringWithoutAuthor())
		}
	}
}
