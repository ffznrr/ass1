package main

import (
	"fmt"
	"os"
	"strconv"
)

type Bio struct {
	nama, alamat, pekerjaan, alasan string
}

func main() {

	// var nama = [2]string{"fauzan", "maman"}
	// fmt.Println(nama)
	args := os.Args

	var isi = [...]string{args[0], args[1]}
	var ket = isi[1]
	n, _ := strconv.Atoi(ket)

	var student = [...]Bio{
		{nama: "fauzan", alamat: "Jakarta", pekerjaan: "belum ada", alasan: "karena coding itu keren"},
		{nama: "maman", alamat: "Bandung", pekerjaan: "Admin", alasan: "karena coding itu mantapp"},
		{nama: "Ridwan", alamat: "Tangerang", pekerjaan: "mahasiswa", alasan: "toppp"},
	}

	var hitung = n - 1

	fmt.Println(student[hitung])

}
