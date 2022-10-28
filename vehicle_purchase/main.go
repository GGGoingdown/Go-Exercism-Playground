package main

import "fmt"

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	switch kind {
	case "car", "truck":
		return true
	default:
		return false
	}
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in lexicographical order.
func ChooseVehicle(option1, option2 string) string {
	if option1 > option2 {
		return fmt.Sprintf("%s is clearly the better choice.", option2)
	}
	return fmt.Sprintf("%s is clearly the better choice.", option1)
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
	var resellRate float64
	switch {
	case age >= 10:
		resellRate = 0.5
	case age < 10 && age > 3:
		resellRate = 0.7
	case age <= 3:
		resellRate = 0.8
	}
	return originalPrice * resellRate

}

func main() {
	needLicense := NeedsLicense("motocycle")
	fmt.Println(needLicense)
	price := CalculateResellPrice(1000, 10.0)
	fmt.Println(price)
	vehicle := ChooseVehicle("2018 Bergamont City", "2020 Gazelle Medeo")
	fmt.Println(vehicle)

	a := 10.0
	fmt.Println(a == 10)

}
