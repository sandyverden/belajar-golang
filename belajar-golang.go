// belajar golang

package main

import "fmt"

func main() {
	// komentar kode
	// menampilkan pesan hello world

	/*
		Komentar multi line
		menampilan pesan hello world
	*/

	fmt.Println("Hello World!")

	// fmt.println("baris ini tidak akan di eksekusi")

	var firstName string = "John"

	var lastName string
	lastName = "Wick"

	fmt.Printf("Hello %s %s!\n", firstName, lastName)

	// deklasarasi variable tanpa tipe data
	// type interface, mengganti operand = menjadi := dan keyword var dihilangkan

	var firstName2 string = "Sandi"

	// Deklasarasi menggunakan := hanya dapat dilakukan didalam blok fungsi
	lastName2 := "Budiman"

	fmt.Printf("Hello %s %s!\n", firstName2, lastName2)

	// Deklarasi multi variable
	var first, second, third string
	first, second, third = "satu", "dua", "tiga"

	fmt.Println(first, second, third)

	// lebih ringkas
	seventh, eight, ninth := "tujuh", "delapan", "sembilan"

	fmt.Printf("dibilang %s, %s, %s\n", seventh, eight, ninth)

	// variable underscore ( _ )
	name, _ := "john", "wick"
	fmt.Println(name)

}
