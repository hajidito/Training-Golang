package method

import (
	"assignment/function"
	"fmt"
	"strings"
)

func TampilkanDataTeman(nama string) {
	for i, teman := range function.DataTeman() {
		if strings.ToLower(teman.Nama) == strings.ToLower(nama) {
			fmt.Printf("Data teman dengan nama %s (absen no %d):\n", teman.Nama, i+1)
			fmt.Printf("Nama: %s\n", teman.Nama)
			fmt.Printf("Alamat: %s\n", teman.Alamat)
			fmt.Printf("Pekerjaan: %s\n", teman.Pekerjaan)
			fmt.Printf("Alasan memilih kelas Golang: %s\n", teman.Alasan)
			return
		}
	}
	fmt.Printf("Teman dengan nama %s tidak ditemukan.\n", nama)
}