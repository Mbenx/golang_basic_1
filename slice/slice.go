package main

import "fmt"

func main() {
	var foods = []string{"egg", "beef", "oat"}
	fmt.Println(foods)
	fmt.Println("Length", len(foods)) // data yang ada, null tidak di hitung
	fmt.Println("cap", cap(foods))    // capacity slice

	var newFoods = foods[:2]
	fmt.Println(newFoods)

	fmt.Println("Length", len(newFoods)) // data yang ada, null tidak di hitung
	fmt.Println("cap", cap(newFoods))    // capacity slice

	// append
	newFoods = append(newFoods, "milk")
	fmt.Println(newFoods)
	fmt.Println("Length", len(newFoods)) // data yang ada, null tidak di hitung
	fmt.Println("cap", cap(newFoods))    // capacity slice

	var newfoods2 = append(foods, "milk")
	fmt.Println(newfoods2)
	fmt.Println("Length", len(newfoods2)) // data yang ada, null tidak di hitung
	fmt.Println("cap", cap(newfoods2))    // capacity slice

}
