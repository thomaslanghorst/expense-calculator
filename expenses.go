package main

import (
	"strings"
)

type Expense struct {
	Category string
	Total    float64
}

func CalculateExpenses(records []*CsvRecord, categories []Category) []Expense {

	mapCategories(records, categories)

	expenses := make([]Expense, 0)
	expensesMap := make(map[string]float64, 0)

	for _, record := range records {
		if _, ok := expensesMap[record.Category]; !ok {
			expensesMap[record.Category] = record.Price
		} else {
			expensesMap[record.Category] = expensesMap[record.Category] + record.Price
		}
	}

	for category, total := range expensesMap {
		expenses = append(expenses, Expense{
			Category: category,
			Total:    total,
		})
	}

	return expenses
}

func mapCategories(records []*CsvRecord, categories []Category) {

	//TODO: build a cache for already known mappings

	for _, record := range records {
		for _, category := range categories {
			for _, shop := range category.Shops {
				if strings.Contains(strings.ToLower(record.Shop), shop) {
					record.Category = category.Name
				}
			}
		}
	}
}
