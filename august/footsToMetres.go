package main

import "fmt"

const Foot = 0.3048

var UnitOfMeasurement string
var Input float64

func main() {
	fmt.Println("Enter the unit of measurement")
	fmt.Scan(&UnitOfMeasurement)

	KindOfUnit()
}

func KindOfUnit() {
	if UnitOfMeasurement == "Foot" {
		fmt.Println("You have enabled the function conversion foots to metres")
		ConversionToMetres()
	} else if UnitOfMeasurement == "Metre" {
		fmt.Println("You have enabled the function conversion metres to foots")
		ConversionToFoot()
	}
}

func ConversionToMetres() {
	fmt.Println("Enter the number:")
	fmt.Scan(&Input)
	Metres := Input * Foot
	fmt.Println("Result:", Metres)

}

func ConversionToFoot() {
	fmt.Println("Enter the number:")
	fmt.Scan(&Input)
	Foots := Input / Foot
	fmt.Println("Result:", Foots)
}
