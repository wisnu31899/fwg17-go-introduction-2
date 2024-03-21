package main

import "fmt"

//soal 1
func main() {
	fmt.Println("Soal 1:")
	val1, rem1 := pembulatan(4.37)
	fmt.Printf("Nilai: %.1f, Sisa: %.1f\n", val1, rem1) // Output: 4.4, 0.0

	val2, rem2 := pembulatan(4.32)
	fmt.Printf("Nilai: %.1f, Sisa: %.1f\n", val2, rem2) // Output: 4.3, 0.0

	val3, rem3 := pembulatan(4.324)
	fmt.Printf("Nilai: %.2f, Sisa: %.2f\n", val3, rem3) // Output: 4.30, 0.0

	val4, rem4 := pembulatan(4.356)
	fmt.Printf("Nilai: %.3f, Sisa: %.3f\n", val4, rem4) // Output: 4.400, 0.0
}

// // soal 2
// func main() {
// 	deret := DeretBilangan(40)
// 	deret.Prima()
// 	deret.Ganjil()
// 	deret.Genap()
// 	deret.Fibonacci()
// }

// soal 3
// func main() {
// 	ukuranBalok := Balok{
// 		Panjang: 3.0,
// 		Lebar:   4.0,
// 		Tinggi:  5.0,
// 	}
// 	cetakInfo(ukuranBalok, "Balok")

// 	ukuranKubus := Kubus{
// 		Sisi: 2.0,
// 	}
// 	cetakInfo(ukuranKubus, "Kubus")

// 	ukuranTabung := Tabung{
// 		JariJari: 2.0,
// 		Tinggi:   5.0,
// 	}
// 	cetakInfo(ukuranTabung, "Tabung")
// }

// // Fungsi untuk mencetak informasi menggunakan interface BangunRuang
// func cetakInfo(b BangunRuang, nama string) {
// 	fmt.Printf("==Rumus %s==\n", nama)
// 	fmt.Printf("Luas: %.2f\n", b.luas())
// 	fmt.Printf("Keliling: %.2f\n", b.keliling())
// 	fmt.Printf("Volume: %.2f\n\n", b.volume())
// }
