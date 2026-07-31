package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println("Hello", i)
	}

	k := 3
	stopWhenReachOne := 1
	for k > 0 {
		fmt.Println("Hellloo", k)
		if k == 1 && stopWhenReachOne > 0 {
			stopWhenReachOne--
			continue
		}
		k--
	}

	items := [3]string{"Hashem", "Ghasem", "Naser"}
	for _, value := range items {
		fmt.Println(value)
	}

}
