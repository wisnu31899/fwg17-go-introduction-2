package main

import "fmt"

//soal 1
func main() {
	fmt.Printf("%.2f\n", pembulatan(4.37))  // Output: 4.40
	fmt.Printf("%.3f\n", pembulatan(4.32))  // Output: 4.30
	fmt.Printf("%.4f\n", pembulatan(4.324)) // Output: 4.30

	fmt.Println(pembulatan(4.37))  // Output: 4.4
	fmt.Println(pembulatan(4.32))  // Output: 4.3
	fmt.Println(pembulatan(4.324)) // Output: 4.3
}

// soal 2
// func main() {
// 	deret := DeretBilangan(10)
// 	deret.Prima()
// 	deret.Ganjil()
// 	deret.Genap()
// 	deret.Fibonacci()
// }
