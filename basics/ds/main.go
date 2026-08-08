package main

import "fmt"

func main() {

	studentGrades := map[string]int{
		"Naser": 20,
		"Jasem": 66,
	}
	studentGrades["Naser"] = 99

	naser, ok := studentGrades["Naser"]
	if !ok {
		return
	}
	fmt.Printf("Naser: %+v\n", naser)

	if hashem, ok := studentGrades["Hashem"]; ok {
		fmt.Printf("Hashem: %+v\n", hashem)
	}

	delete(studentGrades, "Naser")

	classes := make(map[int]string)
	classes[22] = "Collage"

	fmt.Printf("%+v\n", classes)

}

// Arrays
// var numbers [2]int
// 	numbers[0] = 1
// 	numbers[1] = 2
// 	fmt.Printf("%+v\n", numbers)

// 	nums := [4]int{1, 2, 3, 4}
// 	fmt.Printf("%+v\n", nums)

// 	for _, value := range nums {
// 		fmt.Println(value)

// 	}

// 	var matrix [2][3]int
// 	fmt.Printf("%+v\n", matrix)

// names := []string{"hashem", "saber"}
// names = append(names, "jasem")
// for index := 0; index < len(names); index++ {
// 	names[index] = "Mr" + " " + names[index]

// }
// fmt.Printf("%+v\n", names)
// items := make([]int, 3, 4)
// items[2] = 122
// items = append(items, 4)
// fmt.Printf("%+v", "cap: %d\n", items, cap(items))
