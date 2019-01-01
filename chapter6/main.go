package main

import "fmt"

func main() {
	// var x [5]float64
	// x[0] = 98
	// x[1] = 93
	// x[2] = 77
	// x[3] = 82
	// x[4] = 53 // set the 5th element of the array x to 83

	// Slices with Append
	slice1 := []int{1, 2, 3}
	slice2 := append(slice1, 4, 5)
	fmt.Println(slice1, slice2)

	// Slices with Copy
	slice3 := []int{1, 2, 3}
	slice4 := make([]int, 2)
	copy(slice4, slice3)
	fmt.Println(slice3, slice4)

	// Maps
	z := make(map[string]int)
	z["key"] = 10
	fmt.Println(z["key"])

	w := make(map[int]int) // key type of `int`
	w[1] = 27
	fmt.Println(w[1])

	// elements := map[string]string{
	// 	"H":  "Hydrogen",
	// 	"He": "Helium",
	// 	"Li": "Lithium",
	// 	"Be": "Beryllium",
	// 	"B":  "Boron",
	// 	"C":  "Carbon",
	// 	"N":  "Nitrogen",
	// 	"O":  "Oxygen",
	// 	"F":  "Fluorine",
	// 	"Ne": "Neon",
	elements := map[string]map[string]string{
		"H": map[string]string{
			"name":  "Hydrogen",
			"state": "gas",
		},
		"He": map[string]string{
			"name":  "Helium",
			"state": "gas",
		},
		"Li": map[string]string{
			"name":  "Lithium",
			"state": "solid",
		},
		"Be": map[string]string{
			"name":  "Beryllium",
			"state": "solid",
		},
		"B": map[string]string{
			"name":  "Boron",
			"state": "solid",
		},
		"C": map[string]string{
			"name":  "Carbon",
			"state": "solid",
		},
		"N": map[string]string{
			"name":  "Nitrogen",
			"state": "gas",
		},
		"O": map[string]string{
			"name":  "Oxygen",
			"state": "gas",
		},
		"F": map[string]string{
			"name":  "Fluorine",
			"state": "gas",
		},
		"Ne": map[string]string{
			"name":  "Neon",
			"state": "gas",
		},
	}
	if el, ok := elements["Li"]; ok {
		fmt.Println(el["name"], el["state"])
	}

	// Arrays! (Beginning of chapter)
	x := [5]float64{98, 93, 77, 82, 83} // shorter way to create arrays
	var total float64 = 0
	// for i := 0; i < 5; i++ {
	// for i := 0; i < len(x); i++ {
	// total += x[i]
	for _, value := range x { // the underscore is used to tell Go we don't NEED the iterator variable
		total += value
	}
	fmt.Println(total / float64(len(x)))
}
