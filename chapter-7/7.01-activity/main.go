package main

import (
	"errors"
	"fmt"
)

type Payer interface {
	Pay() (string, float64)
}

type employee struct {
	id        int
	firstName string
	lastName  string
}

type developer struct {
	individual        employee
	hourlyRate        float64
	hoursWorkedInYear int
	review            map[string]interface{}
}

func (dev developer) Pay() (string, float64) {
	fullName := dev.FullName()
	return fullName, dev.hourlyRate * float64(dev.hoursWorkedInYear)
}

func (dev developer) FullName() string {
	return fmt.Sprintf("%s %s", dev.individual.firstName, dev.individual.lastName)
}

func (dev developer) GetAverageReview() {
	overallScore := 0
	categories := 0
	for _, val := range dev.review {
		score, err := getReviewScore(val)
		if err == nil {
			categories += 1
			overallScore += score
		}
	}
	if categories != 0 {
		fmt.Printf("%s has an overage rating of %v\n", dev.FullName(), float64(overallScore)/float64(categories))
	} else {
		fmt.Printf("%s has an overage rating of %v\n", dev.FullName(), float64(overallScore))
	}
}

func getReviewScore(initialVal any) (int, error) {
	switch v := initialVal.(type) {
	case string:
		return getScoreFromStringRating(v)
	case int:
		return v, nil
	default:
		return 0, errors.New("unknown rating type")
	}
}

func getScoreFromStringRating(val string) (int, error) {
	switch val {
	case "Excellent":
		return 5, nil
	case "Good":
		return 4, nil
	case "Fair":
		return 3, nil
	case "Poor":
		return 2, nil
	case "Unsatisfactory":
		return 1, nil
	default:
		return 0, errors.New("unknown rating")
	}
}

type manager struct {
	individual     employee
	salary         float64
	commissionRate float64
}

func (man manager) Pay() (string, float64) {
	fullName := man.FullName()
	return fullName, man.salary + (man.salary * man.commissionRate)
}

func (man manager) FullName() string {
	return fmt.Sprintf("%s %s", man.individual.firstName, man.individual.lastName)
}

func payDetails(p Payer) {
	fullName, amount := p.Pay()
	fmt.Printf("%s was compensated with %v\n", fullName, amount)
}

func main() {
	employeeReview := make(map[string]interface{})
	employeeReview["WorkQuality"] = 5
	employeeReview["TeamWork"] = 2
	employeeReview["Communication"] = "Poor"
	employeeReview["Problem-solving"] = 4
	employeeReview["Dependability"] = "Unsatisfactory"
	d := developer{individual: employee{id: 1, firstName: "Eric", lastName: "Davis"}, hourlyRate: 35, hoursWorkedInYear: 2400, review: employeeReview}
	m := manager{individual: employee{id: 2, firstName: "Mr.", lastName: "Boss"}, salary: 150000, commissionRate: .07}
	d.GetAverageReview()
	payDetails(d)
	payDetails(m)
}
