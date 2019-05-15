package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	picUpDropOff()

}
func picUpDropOff() {
	var pickUp, dropOff string
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("What is your pick up location?: ")
	scanner.Scan()
	pickUp = scanner.Text()
	var validateLocationConditioner = validateLocation(pickUp)
	if validateLocationConditioner == false {
		return
	}
	validateLocation(pickUp)
	fmt.Print("Where would you like to go?: ")
	scanner.Scan()
	dropOff = scanner.Text()
	if validateLocationConditioner == false {
		return
	}
	validateLocation(dropOff)
	if determineAmount(pickUp, dropOff) == 0 {
		return
	}
	fmt.Println("Your fare cost is ", determineAmount(pickUp, dropOff))
	time.Sleep(10 * time.Second)
	scanner = bufio.NewScanner(os.Stdin)
	fmt.Print("Enter the specified fare: ")
	scanner.Scan()
	passengerFare := scanner.Text()
	newPassengerFare, _ := strconv.Atoi(passengerFare)
	Fare := determineAmount(pickUp, dropOff)
	for underPaidFare := 0; underPaidFare <= 4; underPaidFare++ {
		if newPassengerFare == Fare {
			fmt.Println("Thank you for your patronage")
			break
		} else if newPassengerFare > Fare {
			fmt.Println("Here is your change: ", newPassengerFare-Fare)
			break
		} else if newPassengerFare < Fare {
			fmt.Println("The money you paid is not enough, please enter the required amount: ")
			scanner.Scan()
			passengerFare = scanner.Text()
			newPassengerFare, _ = strconv.Atoi(passengerFare)
			if underPaidFare == 4 {
				fmt.Print("Enter the right amount or you will be reported to the police: ")
				scanner.Scan()
				passengerFare = scanner.Text()
				newPassengerFare, _ = strconv.Atoi(passengerFare)
			}
		} else {
			fmt.Print("Invalid response")
			break
		}
	}

	requestTip(Fare)
}

func validateLocation(Locations string) (validLocation bool) {
	validLocation = false
	if Locations == "Alakahia" || Locations == "Aluu" || Locations == "Choba" || Locations == "Rumuosi" || Locations == "Rumuola" || Locations == "Rumuokoro" || Locations == "Mgbuoba" {
		return true
	} else if Locations == "" {
		fmt.Println("Invalid destination")
		return false
	} else {
		fmt.Print("Sorry our cab service does not operate in that area, bye laters!")
		return false
	}
	return
}

func determineAmount(fromLocation1, toLocation2 string) (Fare int) {
	if fromLocation1 == "Alakahia" && toLocation2 == "Choba" ||
		fromLocation1 == "Alakahia" && toLocation2 == "Rumuosi" ||
		fromLocation1 == "Choba" && toLocation2 == "Alakahia" ||
		fromLocation1 == "Choba" && toLocation2 == "Aluu" ||
		fromLocation1 == "Aluu" && toLocation2 == "Choba" ||
		fromLocation1 == "Choba" && toLocation2 == "Rumuosi" ||
		fromLocation1 == "Rumuosi" && toLocation2 == "Choba" ||
		fromLocation1 == "Rumuosi" && toLocation2 == "Alakahia" {
		return 50
	} else if fromLocation1 == "Alakahia" && toLocation2 == "Rumuokoro" ||
		fromLocation1 == "Choba" && toLocation2 == "Rumuokoro" ||
		fromLocation1 == "Rumuola" && toLocation2 == "Rumuokoro" ||
		fromLocation1 == "Rumuosi" && toLocation2 == "Rumuokoro" ||
		fromLocation1 == "Rumuokoro" && toLocation2 == "Rumuosi" ||
		fromLocation1 == "Rumuokoro" && toLocation2 == "Alakahia" ||
		fromLocation1 == "Rumuokoro" && toLocation2 == "Rumuola" ||
		fromLocation1 == "Rumuokoro" && toLocation2 == "Mgbuoba" ||
		fromLocation1 == "Mgbuoba" && toLocation2 == "Choba" ||
		fromLocation1 == "Mgbuoba" && toLocation2 == "Rumuokoro" {
		return 100
	} else if fromLocation1 == "Rumuola" && toLocation2 == "Choba" ||
		fromLocation1 == "Choba" && toLocation2 == "Rumuola" {
		return 200
	} else {
		return
	}
	return
}

func requestTip(tip int) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("What amount is your tip? : ")
	scanner.Scan()
	tipOffered := scanner.Text()
	newTipFromCustomer, _ := strconv.Atoi(tipOffered)
	switch newTipFromCustomer >= 0 {
	case newTipFromCustomer == 0:
		fmt.Println("You are a stingy fool")
	case newTipFromCustomer <= tip:
		fmt.Println("Thank you")
	case newTipFromCustomer >= tip:
		fmt.Println("Gracias mucho Mr. Reeves")
	}

}
