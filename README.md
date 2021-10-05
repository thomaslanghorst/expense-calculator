# expense-calculator
This project is designed to read an CSV files in the CSV-MT940 format containing a list of expenses. These CSV files can typically be downloaded on the webpage of your online bank.

### Setup
1. First you need to rename the _categories-example.json_ file using `mv categories-example.json categories.json`. Afterwards you can simply extend the _categories.json_ file to meet your personal needs.

2. After downloading your expenses CSV file you can simply copy it into the root of this project and rename it to _expenses.CSV_.

### Building the expense-calculator
    go build 

### Running the expense-calculator
    ./expenses-calculator

The output should look something like this

    groceries: 500.50
    gas: 150.37
    amazon: 199.99