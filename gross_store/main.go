package main

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	units := map[string]int{
		"quarter_of_a_dozen": 3,
		"half_of_a_dozen":    6,
		"dozen":              12,
		"small_gross":        120,
		"gross":              144,
		"great_gross":        1728,
	}
	return units
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	return make(map[string]int)
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	uV, uExist := units[unit]
	if !uExist {
		return false
	}
	bV, bExist := bill[item]
	if bExist {
		bill[item] = bV + uV
		return true
	}
	bill[item] = uV
	return true
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	v, ex1 := bill[item]
	if !ex1 {
		return false
	}

	uV, ex2 := units[unit]
	if !ex2 {
		return false
	}

	newQuantity := v - uV
	if newQuantity < 0 {
		return false
	}
	if newQuantity == 0 {
		delete(bill, item)
		return true
	}
	bill[item] = newQuantity
	return true
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	v, exist := bill[item]
	if !exist {
		return 0, false
	}
	return v, true
}

func main() {
	m := map[string]int{"hello": 10}
	delete(m, "world")
}
