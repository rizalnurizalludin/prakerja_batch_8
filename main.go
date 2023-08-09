package main

import (
	"fmt"
	"prakerja_batch_8/quiz"
)

func main() {

	// age := 17

	// if age == 17 {
	// 	fmt.Println("Seventeen")
	// } else if age > 10 {
	// 	fmt.Println("sepuluh")
	// } else {
	// 	fmt.Println("Incalid")
	// }

	var bil int

	fmt.Println("MENENTUKAN BILANGAN PRIMA DAN KELIPATAN 7")
	fmt.Print("Masukkan angka: ")
	fmt.Scan(&bil)

	quiz.Bil_prima(bil)
	quiz.Bil_kel_7(bil)

	var t, a, b int

	fmt.Println("MENGHITUNG LUAS TRAPESIUM")
	fmt.Print("Masukkan a : ")
	fmt.Scan(&a)
	fmt.Print("Masukkan b : ")
	fmt.Scan(&b)
	fmt.Print("Masukkan t : ")
	fmt.Scan(&t)

	Luas := quiz.Luas_trapesium(a, b, t)
	fmt.Println(Luas)
}
