package quiz

import "fmt"

func Bil_prima(bil int) {
	faktor := 0

	for i := 1; i <= bil; i++ {
		if bil%i == 0 {
			faktor++
		}
	}
	if faktor == 2 {
		fmt.Println("Bilangan Prima")
	} else {
		fmt.Println("Bukan Bilangan Prima")
	}

}

func Bil_kel_7(bil int) {
	if bil%7 == 0 {
		fmt.Println("Bilangan Kelipatan 7")
	} else {
		fmt.Println("Bukan Kelipatan 7")
	}
}

func Luas_trapesium(a int, b int, t int) int {
	var Luas int
	Luas = (t * (a + b)) / 2
	return Luas
}
