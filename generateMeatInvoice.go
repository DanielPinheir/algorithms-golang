package main

import "fmt"

func main() {
	var meatType uint16
	var meatQuantity float32
	var creditCard uint16
	var kgMeat float32
	var purchaseValue float32

	fmt.Println("Enter the type of meat:")
	fmt.Scan(&meatType)

	fmt.Println("Enter the amount:")
	fmt.Scan(&meatQuantity)

	fmt.Println("Payment with Tabajara card?: 1=Yes 2=No")
	fmt.Scan(&creditCard)

	switch meatType {
	case 1:
		fmt.Println("Type of meat chosen: Double Filet")
		if meatQuantity > 5 {
			kgMeat = 4.90
		} else {
			kgMeat = 5.80
		}
	case 2:
		fmt.Println("Type of meat chosen: Rump")
		if meatQuantity > 5 {
			kgMeat = 5.90
		} else {
			kgMeat = 6.80
		}
	case 3:
		fmt.Println("Type of meat chosen: Picanha")
		if meatQuantity > 5 {
			kgMeat = 6.90
		} else {
			kgMeat = 7.80
		}
	}

	if creditCard == 1 {
		purchaseValue = (kgMeat * meatQuantity) - ((kgMeat * meatQuantity) * 0.05)
	} else {
		purchaseValue = kgMeat * meatQuantity
	}

	fmt.Println("Total Purchase Value: ", purchaseValue)
}