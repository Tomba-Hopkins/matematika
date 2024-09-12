package util

import "fmt"

// Penjumlahan, pengurangan, perkalian dan teman teman
func Add(a, b int) int {
	return a + b
}



// Matematika deret


// versi A sudah diketahui / Un = 1 :
// rumus 1
func FindNthTerm(suku1, nilaiSuku1, beda, sukuN int) (int, string) {
	if suku1 > 1 || suku1 < 1 {
		return 0, "\nSuku yang diketahui harus 1, jika lebih pakailah rumus lanjutan"
	}
	jawaban := nilaiSuku1 + ((sukuN - 1) * beda)
	return jawaban, ""
}

// rumus 2
func FindNthTerm2(suku1, nilaiSuku1, beda, sukuN int) (int, string) {
	if suku1 > 1 || suku1 < 1 {
		return 0, "\nSuku yang diketahui harus 1, jika lebih pakailah rumus lanjutan"
	}
	jawaban := beda * sukuN

	if beda > nilaiSuku1 {
		jawaban -= beda - nilaiSuku1
	} else if beda < nilaiSuku1 {
		jawaban += nilaiSuku1 - beda
	}

	return jawaban, ""
}

// versi A belum diketahui / Un > 1 :
// rumus 1
func FindNthTermWithUnknownA(sukuDiketahui, nilaiSukuDiketahui, beda, sukuN int) int {
	a := 0
	if sukuDiketahui > 1 {
		temp := (sukuDiketahui - 1) * beda
		a = nilaiSukuDiketahui - temp
	} else {
		return nilaiSukuDiketahui + (sukuN - 1) * beda
	}
	fmt.Println("Nilai awal: ",a)
	jawaban := a + ((sukuN - 1) * beda)
	return jawaban
}