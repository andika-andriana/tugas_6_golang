package library

import "fmt"

type mahasiswa struct {
	nama string
	umur int
}

func Tampil_mahasiswa() {
	var s1 = mahasiswa{"Andika Andriana", 24}

	fmt.Println("Nama Mahasiswa\t: ", s1.nama)
	fmt.Println("Umur Mahasiswa\t: ", s1.umur)
}
