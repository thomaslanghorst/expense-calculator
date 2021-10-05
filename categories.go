package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

type Category struct {
	Name  string   `json:"name"`
	Shops []string `json:"shops"`
}

func ReadCategories() ([]Category, error) {

	var categories []Category

	jsonFile, err := os.Open(categoriesFile)
	if err != nil {
		return categories, errors.New(fmt.Sprintf("unable to open categories file. Error: %s", err.Error()))
	}

	bytes, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return categories, errors.New(fmt.Sprintf("unable to read categories file. Error: %s", err.Error()))
	}

	err = json.Unmarshal(bytes, &categories)
	if err != nil {
		return categories, errors.New(fmt.Sprintf("unable to read unmarshal categories file. Error: %s", err.Error()))
	}

	return categories, nil
}
