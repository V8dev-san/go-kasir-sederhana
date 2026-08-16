package main

import "fmt"

func main() {

	// deklarasi waribale
	var kodeBarang int
    var namaBarang string
    var jumlahBarang int
	var hargaBarang int
	var namaKasir string 
	var namaToko string
	var tanggal string

	// tampikan harga barang
	fmt.Println("=== Fajar Jaya minimarket===")
	fmt.Println("nama kasir : Naruto", namaKasir)
	fmt.Println("Fajar Jaya minimarket", namaToko)
	fmt.Println("tanggal : 28 Mei 2026", tanggal)
	fmt.Println("==================================")

	// kasir input data pembeli
	fmt.Print("Masukan kode barang : ")
	fmt.Scanln(&kodeBarang)

	fmt.Print("Masukan jumlah barang : ")
	fmt.Scanln(&jumlahBarang)

	// switch untuk mentukan harga barang dan jumlah barang
	switch kodeBarang {
	case 101:
		namaBarang = "sabun mandi"
		hargaBarang = 5000
	case 102:
		namaBarang = "indomie"
		hargaBarang = 3000
	case 103:
		namaBarang = "teh javana"
		hargaBarang = 3500
	default:
		namaBarang = "barang tidak di temukan"
		hargaBarang = 0
	}

	// hitung total belanja
    var totalBelanja = jumlahBarang * hargaBarang 
	totalBelanja = jumlahBarang * hargaBarang

	// cetak struk belanja gunaka printf
	fmt.Printf("nama barang : %s\n", namaBarang)
	fmt.Printf("jumlah barang : %d\n", jumlahBarang)
	fmt.Printf("harga barang : Rp.%d\n", hargaBarang)
	fmt.Println("------------------------------")
	fmt.Printf("total belanja : Rp.%d\n", totalBelanja)
	fmt.Println("===============================")
	fmt.Println("Terima kasih telah berbelanja di Minimarket Fajar")


}