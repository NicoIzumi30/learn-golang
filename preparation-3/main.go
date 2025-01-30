package main

import (
	"fmt"
)

func sayHello() {
	fmt.Println("Hello World")
}

func main() {

	// Defer
	defer sayHello() // Defer akan di panggil di akhir kode

	// Perkondisian menggunakan if
	age := 50

	if age <= 18 {
		fmt.Println("Masih kecil dek.")
	} else if age > 60 {
		fmt.Println("Tua banget bang")
	} else {
		fmt.Println("Nahh cocok..")
	}

	// Perkondisian menggunakan switch
	grade := "F"
	switch grade {
	case "A":
		fmt.Printf("GG Bang dapet A\n")
	case "B", "C":
		fmt.Printf("Nicee, lanjutkan bang\n")
	case "F":
		fmt.Printf("Gapapaa bang, belum takdirnya..\n")
		fallthrough
	default:
		fmt.Println("Belajar lagi yaa...\n")
	}

	// Perulangan menggunakan for
	for i := 0; i < 5; i++ {
		fmt.Printf("Perulangan ke-%d\n", i)
	}

	// Perulangan menggunakan while
	j := 0
	for j < 5 {
		fmt.Printf("Perulangan ke-%d\n", j)
		j++
	}

	// Perulangan khusus
	i := 0
	for {
		if i >= 5 {
			break
		}
		i++
		if i == 2 {
			continue
		}
		fmt.Printf("Perulangan ke-%d\n", i)
	}

	// Perulangan dalam array
	numbers := []int64{1, 2, 3, 4, 100, 10, 50}
	for i, v := range numbers {
		fmt.Printf("Index ke %d memiliki value %d\n", i, v)
		if v == 100 {
			fmt.Println("Dapat jacpot bang")
		}
	}
}
