package main

import "fmt"

const (
	LoungeCapacity       = 150
	VIPCapacity          = 50
	VIPSectionTables     = 10
	VIPSectionCost       = 100
	GeneralAdmissionCost = 100
	MonthlyTurnover      = 1000000
)

func main() {
	var vipBookings int
	var generalBookings int

	// Allow bookings up to capacity
	for i := 0; i < LoungeCapacity; i++ {
		if i < VIPCapacity {
			// VIP booking
			if (i % VIPSectionTables) == 0 {
				vipBookings++
			}
		}

		// Calculate total profit and profit per month
		totalProfit := (vipBookings * VIPSectionCost) + (generalBookings * GeneralAdmissionCost)
		profitPerMonth := totalProfit * 30

		// Calculate minimum number of guests required to meet monthly target
		minimumGuests := (MonthlyTurnover / 30) / GeneralAdmissionCost

		// Print the results
		fmt.Printf("Total profit: R%d\n", totalProfit)
		fmt.Printf("Profit per month: R%d\n", profitPerMonth)
		fmt.Printf("Minimum number of guests for monthly target: %d\n", minimumGuests)
	}
}
