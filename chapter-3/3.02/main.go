package main

import "fmt"

func main() {
	creditScore := 500
	monthlyIncome := float32(1000)
	loanAmount := float32(1000)
	loanTerm := 24
	rate := getInterestRate(creditScore)
	monthlyPayment := getMonthlyPayment(loanAmount, loanTerm, rate)
	totalCost := (monthlyPayment * float32(loanTerm)) - loanAmount
	approved := false
	if monthlyIncome > monthlyPayment && monthlyPayment <= getMaxMonthlyPayments(monthlyIncome, creditScore) {
		approved = true
	}
	fmt.Println("Credit Score    :", creditScore)
	fmt.Println("monthlyIncome   :", monthlyIncome)
	fmt.Println("Loan Amount     :", loanAmount)
	fmt.Println("Loan Term       :", loanTerm)
	fmt.Println("Monthly Payment :", monthlyPayment)
	fmt.Println("Rate            :", rate)
	fmt.Println("Total Cost      :", totalCost)
	fmt.Println("Approved        :", approved)
}

func getMonthlyPayment(loanAmount float32, loanTerm int, interestRate float32) float32 {
	interest := (loanAmount * interestRate) / float32(loanTerm)
	principal := loanAmount / float32(loanTerm)
	return principal + interest
}

func getInterestRate(creditScore int) float32 {
	if creditScore >= 450 {
		return 0.15
	}
	return 0.2
}

func getMaxMonthlyPayments(monthlyIncome float32, creditScore int) float32 {
	if creditScore >= 450 {
		return monthlyIncome * 0.2
	}
	return monthlyIncome * 0.1
}
