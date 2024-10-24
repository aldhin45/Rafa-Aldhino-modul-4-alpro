package main

import "fmt"

func hitungSkor(waktuSoal [8]int) (int, int) {
	const waktuMaksimal = 301
	totalWaktu := 0
	soalDiselesaikan := 0

	for i := 0; i < 8; i++ {
		if waktuSoal[i] <= waktuMaksimal {
			soalDiselesaikan++
			totalWaktu += waktuSoal[i]
		} else {
			totalWaktu += waktuMaksimal
		}
	}
	return soalDiselesaikan, totalWaktu
}

func main() {
	var waktuPeserta1, waktuPeserta2 [8]int

	fmt.Println("Astuti:")
	for i := 0; i < 8; i++ {
		fmt.Printf("%d: ", i+1)
		fmt.Scan(&waktuPeserta1[i])
	}

	fmt.Println("Bertha:")
	for i := 0; i < 8; i++ {
		fmt.Printf(" %d: ", i+1)
		fmt.Scan(&waktuPeserta2[i])
	}

	selesai1, waktu1 := hitungSkor(waktuPeserta1)
	selesai2, waktu2 := hitungSkor(waktuPeserta2)

	fmt.Printf("Astuti %d %d \n", selesai1, waktu1)
	fmt.Printf("Bertha %d %d \n", selesai2, waktu2)

	if selesai1 > selesai2 || (selesai1 == selesai2 && waktu1 < waktu2) {
		fmt.Printf("Pemenangnya adalah Astuti dengan %d soal dan waktu %d menit\n", selesai1, waktu1)
	} else {
		fmt.Printf("Pemenangnya adalah Bertha dengan %d soal dan waktu %d menit\n", selesai2, waktu2)
	}
}
