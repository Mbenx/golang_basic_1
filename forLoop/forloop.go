package main

import "fmt"

func main() {
	i := 0

	for i < 5 {
		// code yang akan di eksekusi bila kondisi TRUE
		fmt.Println("Value of i : ", i)
		i++ // increment == i + 1
	}

	for {
		fmt.Println("Value of i : ", i)
		i++

		if i == 5 {
			break
		}
	}

	for {

		fmt.Println("Value of i : ", i)
		if i == 5 {
			break
		}
		i++
	}

	for x := 0; x <= 5; x++ {

		// if x%2 == 1 {
		// 	continue
		// }

		if x == 1 {
			continue
		}

		if x == 3 {
			continue
		}

		// if x == 3 {
		// 	break
		// }

		fmt.Println("Value of x : ", x)
	}
}
