package main

import (
	"errors"
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
)

type BankAccount struct {
	user_id      int
	first_name   string
	last_name    string
	credit_score int
	balance      float64
	last_salary  string
}

func (b *BankAccount) get_salary(salary_amount float64) (string, error) {
	salaryDate, err := time.Parse("2006-01-02", b.last_salary)
	if err != nil {
		return "", errors.New("the parsing of salaryDate has failed")
	}
	NewSalaryDate := salaryDate.AddDate(1, 0, 0).Format("2006-01-02")
	currentDate := time.Now().Format("2006-01-02")

	if currentDate >= NewSalaryDate {
		b.balance += salary_amount
		b.last_salary = currentDate
	} else {
		return "", errors.New("Its not the date yet.")
	}
	message := fmt.Sprintf("We have succesfully added %v to our bank account, the total balance now is %v", salary_amount, b.balance)
	return message, nil
}

func (b *BankAccount) deposit(amount float64) (string, float64, error) {
	if b.credit_score <= 100 {
		return "", 0, errors.New("Your credit score is too low")
	}
	b.balance += amount
	return "your deposit is succesful", b.balance, nil
}

func (b *BankAccount) withdraw(amount float64) (string, float64, error) {
	if b.balance < amount {
		return "", b.balance, errors.New("Insufficient Funds")
	}
	b.balance -= amount
	return "withdraw succesful", b.balance, nil
}

func main() {
	nadia := &BankAccount{
		user_id:      1,
		first_name:   "nadia",
		last_name:    "vidatalla",
		credit_score: 100,
		balance:      0,
		last_salary:  "2024-12-08",
	}
	logrus.Info("Bank account application started")

	fmt.Println(*nadia)
	check, err := nadia.get_salary(500)
	fmt.Println(*nadia, check, err)

	check_2, err := nadia.get_salary(500)

	fmt.Println(*nadia, check_2, err)
	a, b, c := nadia.deposit(1000)
	fmt.Println(*nadia, a, b, c)

	nadia.credit_score += 1
	d, e, f := nadia.deposit(1000)
	fmt.Println(*nadia, d, e, f)
	z, x, v := nadia.withdraw(500)
	fmt.Println(*nadia, z, x, v)
	j, k, l := nadia.withdraw(1001)
	fmt.Println(*nadia, j, k, l)
}
