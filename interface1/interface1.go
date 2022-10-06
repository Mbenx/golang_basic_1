package main

import "fmt"

func main() {
	var x interface{}

	x = "Interface in string"
	fmt.Println(x)

	x = []string{"abc", "qwewr", "123"}
	fmt.Println(x)

	x = 10.12
	fmt.Println(x)

	xyz := map[string]interface{}{
		"Name":  "Heru",
		"Nilai": 80,
		"Skill": []string{"renang", "lari"},
	}

	fmt.Println(xyz)

}
