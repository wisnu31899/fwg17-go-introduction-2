package main

import (
	"fmt"
)

// Struct untuk menyimpan deret bilangan
type Deret struct {
	Limit int
}

// Method untuk membuat objek Deret
func DeretBilangan(limit int) *Deret {
	return &Deret{
		Limit: limit,
	}
}

// Fungsi bantu untuk mengecek bilangan prima
func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	if n <= 3 {
		return true
	}
	if n%2 == 0 || n%3 == 0 {
		return false
	}
	for i := 5; i*i <= n; i = i + 6 {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}
	return true
}

// Method untuk mencetak deret bilangan prima
func (d *Deret) Prima() {
	fmt.Print("Deret bilangan prima: ")
	for i := 2; i <= d.Limit; i++ {
		if isPrime(i) {
			fmt.Print(i, " ")
		}
	}
	fmt.Println()
}

// Method untuk mencetak deret bilangan ganjil
func (d *Deret) Ganjil() {
	fmt.Print("Deret bilangan ganjil: ")
	for i := 1; i <= d.Limit; i += 2 {
		fmt.Print(i, " ")
	}
	fmt.Println()
}

// Method untuk mencetak deret bilangan genap
func (d *Deret) Genap() {
	fmt.Print("Deret bilangan genap: ")
	for i := 0; i <= d.Limit; i += 2 {
		fmt.Print(i, " ")
	}
	fmt.Println()
}

// Method untuk mencetak deret bilangan Fibonacci
func (d *Deret) Fibonacci() {
	fmt.Print("Deret bilangan Fibonacci: ")
	a, b := 0, 1
	for a <= d.Limit {
		fmt.Print(a, " ")
		a, b = b, a+b
	}
	fmt.Println()
}
