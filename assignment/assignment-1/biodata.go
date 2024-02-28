/*
	Assignment 1: Session - 3 Function, Method, Pointer, Struct & Exported/Unexported
	Name  : Muhammad Rifqi Setiawan
	Class : Golang-1
	Registration Number : 1957356840-118
*/

package main

import (
	"fmt"
	"os"
	"strconv"
)

type GolangFriends struct {
	nama        string
	alamat      string
	pekerjaan   string
	alasanMasuk string
}

var friendsList = map[int]GolangFriends{
	1: {"Abdul Majid Refindo", "Jakarta", "Software Engineer", "Ingin mempelajari bahasa pemrograman Go"},
	2: {"Muhammad Nahrowi", "Bandung", "Data Scientist", "Tertarik dengan efisiensi dan performa Go dalam pengolahan data besar"},
	3: {"Arif Hakam Hidayat", "Klaten", "Programmer", "Mencari alternatif bahasa pemrograman yang lebih efisien dalam penggunaan resource"},
	4: {"Muhammad Ali Akbar", "Yogyakarta", "Backend Engineering", "Ingin mempelajari backend development dengan Go"},
	5: {"Tamir Gading Hasibuan", "Surabaya", "Fullstack Developer", "Ingin memperdalam pengetahuan tentang Go dan mempelajari best practice dalam pengembangan aplikasi"},
}

func printData(absen int) {
	friends, found := friendsList[absen]

	if !found {
		fmt.Println("Data teman dengan absen", absen, "tidak dapat ditemukan!")
		return
	}

	fmt.Println("Data teman dengan absen nomor", absen, "adalah:")
	fmt.Println("Nama\t\t:", friends.nama)
	fmt.Println("Alamat\t\t:", friends.alamat)
	fmt.Println("Pekerjaan\t:", friends.pekerjaan)
	fmt.Println("Alasan Masuk\t:", friends.alasanMasuk)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run biodata.go <nomor_absen>")
		return
	}

	absenStr := os.Args[1]
	absen, err := strconv.Atoi(absenStr)
	if err != nil {
		fmt.Println("Nomor absen harus berupa bilangan bulat.")
		return
	}

	printData(absen)
}
