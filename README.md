SOAL NOMOR 1
Fungsi factorial(n int) int:
Menghitung faktorial dari bilangan n menggunakan perulangan dari 2 hingga n.
Nilai faktorial dihitung dengan mengalikan semua bilangan dari 1 hingga n.
fungsi permutasi(n, r int) int:
Menghitung permutasi P(n, r)
Memanfaatkan fungsi factorial untuk menghitung faktorial dari n dan (n-r).
ungsi combinasi(n, r int) int:
Menghitung kombinasi C(n, r)
Memanfaatkan fungsi factorial untuk menghitung faktorial dari n, r, dan (n-r).
fungsi main():
Meminta input dua nilai integer n dan r dari pengguna.
Jika n lebih besar atau sama dengan r, program akan menghitung dan mencetak nilai permutasi dan kombinasi menggunakan fungsi yang sudah didefinisikan.

SOAL NOMOR 2
Fungsi hitungSkor(waktuSoal [8]int) (int, int):
Fungsi menghitung jumlah soal yang diselesaikan dan total waktu yang dibutuhkan peserta.
Jika waktu pengerjaan soal lebih dari 301 menit (waktu maksimal), dianggap tidak terjawab, dan waktu yang dihitung adalah 301 menit untuk soal tersebut.
Fungsi mengembalikan dua nilai: jumlah soal yang diselesaikan dan total waktu pengerjaan.
Fungsi main():
Program meminta input waktu pengerjaan 8 soal dari dua peserta: Astuti dan Bertha.
Input diterima dalam bentuk array [8]int untuk setiap peserta.
Program memanggil fungsi hitungSkor untuk menghitung skor dan waktu total untuk masing-masing peserta.

SOAL NOMOR 3
Fungsi cetakDeret(n int):
Fungsi ini mencetak deret bilangan dari nilai awal n hingga mencapai 1.
Aturan Collatz diterapkan pada setiap bilangan dalam deret: jika n genap, bagi dengan 2, dan jika ganjil, kalikan dengan 3 lalu tambahkan 1.
Bilangan terakhir yang selalu dicetak adalah 1.
Fungsi main():
Meminta input bilangan n dari pengguna.
Memanggil fungsi cetakDeret untuk mencetak deret bilangan dari n.


