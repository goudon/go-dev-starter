package main

import "fmt"

func main() {
	// memo : 配列へのアクセス
	// colors := [...]string{"Red", "Green", "Blue"}
	// for i, color := range colors {
	// 	fmt.Printf("%d: %s\n", i, color)
	// }
	fmt.Println("Start it !")
	const max_number = 100
	const fizz = "Fizz"
	const buzz = "Buzz"
	for i := 0; i < max_number; i++ { 
		// -- memo: if文を用いたもの
		// if i % 3 == 0 && i % 5 == 0 {
		// 	fmt.Println(fizz+buzz)
		// } else if i % 3 == 0 {
		// 	fmt.Println(fizz)
		// } else if i % 5 == 0 {
		// 	fmt.Println(buzz)
		// } else {
		// 	fmt.Println(i)
		// }

		// -- memo : case文（Break文がない）
		switch {
		case i % 3 == 0 && i % 5 == 0 :
			fmt.Println(fizz+buzz)
		case i % 3 == 0:
			fmt.Println(fizz)
		case i % 5 == 0:
			fmt.Println(buzz)
		default:
			fmt.Println(i)
		}
	}
	
    fmt.Println("End it !")

}