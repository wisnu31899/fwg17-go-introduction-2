package main

import "fmt"

//soal 1
// func main() {
// 	fmt.Printf("%.0f\n", pembulatan(4.37))  // Output: 4.4
// 	fmt.Printf("%.1f\n", pembulatan(4.32))  // Output: 4.3
// 	fmt.Printf("%.2f\n", pembulatan(4.324)) // Output: 4.30
// 	fmt.Printf("%.3f\n", pembulatan(4.356)) // Output: 4.40
// }

// // soal 2
// func main() {
// 	deret := DeretBilangan(10)
// 	deret.Prima()
// 	deret.Ganjil()
// 	deret.Genap()
// 	deret.Fibonacci()
// }

// soal 3
func main() {
	ukuranBalok := Balok{
		Panjang: 3.0,
		Lebar:   4.0,
		Tinggi:  5.0,
	}
	cetakInfo(ukuranBalok, "Balok")

	ukuranKubus := Kubus{
		Sisi: 2.0,
	}
	cetakInfo(ukuranKubus, "Kubus")

	ukuranTabung := Tabung{
		JariJari: 2.0,
		Tinggi:   5.0,
	}
	cetakInfo(ukuranTabung, "Tabung")
}

// Fungsi untuk mencetak informasi menggunakan interface BangunRuang
func cetakInfo(b BangunRuang, nama string) {
	fmt.Printf("==Rumus %s==\n", nama)
	fmt.Printf("Luas: %.2f\n", b.luas())
	fmt.Printf("Keliling: %.2f\n", b.keliling())
	fmt.Printf("Volume: %.2f\n\n", b.volume())
}
