package main

//this is gonna hold my order struct and its function

import "fmt"

//Order struct that contains the basic information for an order at El toro
type Order struct {
	ID       string
	Campaign string
	Company  string
	URL      string
	Budget   float64
	daysLeft int
}

//Prints out and orderly looking
func (o Order) returnInfo() {
	fmt.Printf("Order ID: %s\nCampaign: %s\nCompany: %s\nURL: %s\nBudget: $%f\nDays Left: %d", o.ID, o.Campaign, o.Company, o.URL, o.Budget, o.daysLeft)
}
