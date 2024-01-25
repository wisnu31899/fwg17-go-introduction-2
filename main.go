package main

import "fmt"

//soal 1
// func main() {
// 	fmt.Printf("%.2f\n", pembulatan(4.37))  // Output: 4.40
// 	fmt.Printf("%.3f\n", pembulatan(4.32))  // Output: 4.30
// 	fmt.Printf("%.4f\n", pembulatan(4.324)) // Output: 4.30
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
	// Buat instance Balok
	ukuranBalok := Balok{
		Panjang: 3.0,
		Lebar:   4.0,
		Tinggi:  5.0,
	}

	// Panggil metode cetak dari instance Balok
	cetakInfo(ukuranBalok, "Balok")

	// Buat instance Kubus
	ukuranKubus := Kubus{
		Sisi: 2.0,
	}

	// Panggil metode cetak dari instance Kubus
	cetakInfo(ukuranKubus, "Kubus")

	// Buat instance Tabung
	ukuranTabung := Tabung{
		JariJari: 2.0,
		Tinggi:   5.0,
	}

	// Panggil metode cetak dari instance Tabung
	cetakInfo(ukuranTabung, "Tabung")
}

// Fungsi untuk mencetak informasi menggunakan interface BangunRuang
func cetakInfo(b BangunRuang, nama string) {
	fmt.Printf("==Rumus %s==\n", nama)
	fmt.Printf("Luas: %.2f\n", b.luas())
	fmt.Printf("Keliling: %.2f\n", b.keliling())
	fmt.Printf("Volume: %.2f\n\n", b.volume())
}
