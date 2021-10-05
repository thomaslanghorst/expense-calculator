package main

import (
	"fmt"
	"log"
)

var (
	csvFile        = "./expenses-example.CSV"
	categoriesFile = "./categories-example.json"
)

func main() {

	records, err := ReadCsv()
	if err != nil {
		log.Fatalf("error parsing csv file. Error: %s\n", err.Error())
	}

	categories, err := ReadCategories()
	if err != nil {
		log.Fatalf("error categories file. Error: %s\n", err.Error())
	}

	expenses := CalculateExpenses(records, categories)

	for _, expense := range expenses {
		fmt.Printf("%s: %.2f\n", expense.Category, expense.Total)
	}

}
