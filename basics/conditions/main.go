package main

import "fmt"

var productPrices = map[string]float64{
	"T-shirt": 30.03,
	"Mug":     20.18,
	"Book":    21.21,
}

func calculateItemPrice(item string) (float64, bool) {
	price, found := productPrices[item]
	if !found {
		return 0.00, false
	}
	return price * 0.95, true
}

func main() {
	fmt.Println(
		calculateItemPrice("T-shirt"),
	)
	// temp := 30
	// if temp > 30 {
	// 	fmt.Println("Too hot brother!")
	// } else if temp == 30 {
	// 	fmt.Println("good")

	// } else {
	// 	fmt.Println("cool weather")

	// }

	// userAccess := map[string]bool{
	// 	"Hashem": false,
	// 	"Naser":  true,
	// }
	// if userAccess["Naser"] {
	// 	fmt.Println("Naser has access")
	// }

}
