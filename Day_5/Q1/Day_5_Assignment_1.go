package main

import (
	"fmt"
)


//Fungsi untuk soal 1.a
func list(pembelian_cap float64, barang map[int]string, price map[int]float64) {
	var j int = 0
	var container float64 = 0
	var key_ref int = 0
	var list_barang [10]string
	var list_harga [10]float64
	
	//Membuat object map baru, agar tidak mengubah original
	var price_clone = make(map[int]float64)
	for key, value := range price{
		price_clone[key] = value
	}

	//Looping untuk menemukan barang dengan harga yang paling kecil sampai "container" lebih besar atau sama dengan pembelian_cap
	for j = 1; j<=len(barang); j++ {
		key_ref, list_harga[j], list_barang[j] = murah(barang, price_clone)

		//Bila container belum memenuhi kapasitas
		if container < pembelian_cap {
			container = container + list_harga[j]
			fmt.Printf("%s dengan harga: Rp.%.2f\n", list_barang[j], list_harga[j])
			
			//Delete map price_clone agar harga terendah tidak diulang
			delete(price_clone, key_ref)
		} else {
			fmt.Printf("\nJumlah total: Rp.%.2f\n", container)
			break
		}
	}
}
	
//Fungsi untuk soal 1.b
func murah(barang map[int]string, price map[int]float64) (key int, minimum float64, produk_min string){
	var j int = 0
	for j = 1; j<=len(barang); j++{
		if price[j] == price[1]{
			//Menyimpan value awal
			minimum = price[j]
		} else if price[j] > minimum{
			continue
		} else if price[j] < minimum && price[j] != 0{
			
			//Menyimpan value terendah
			minimum = price[j]
			produk_min = barang[j]
			
			//Digunakan untuk func list
			key = j
		}
	} 
	return key, minimum, produk_min
}

//Fungsi untuk soal 1.b
func mahal(barang map[int]string, price map[int]float64) (maksimum float64, produk_maks string){
	var j int = 0
	for j = 1; j<=len(barang); j++{
		if price[j] == price[1]{
			//Menyimpan value awal
			maksimum = price[j]
		} else if price[j] < maksimum{
			continue
		} else if price[j] > maksimum{
			
			//Menyimpan value tertinggi
			maksimum = price[j]
			produk_maks = barang[j]
		}
	} 
	return maksimum, produk_maks
}

//Fungsi untuk soal 1.c
func only_10k(barang map[int]string, price map[int]float64){
	var j int = 0
	for j = 1; j<=len(barang); j++{
		if price[j] == 10000{
			fmt.Printf("%s - %.2f\n", barang[j], price[j])
		}
	} 
}

func main() {
	
	//Deklarasi kapasitas pembelian total
	var pembelian_cap float64 = 100000

	//Deklarasi list barang
	var barang = map[int]string{1:"Benih Lele",2:"Pakan lele cap menara",3:"Probiotik A",4:"Probiotik Nila B",5:"Pakan Nila",
	6:"Benih Nila",7:"Cupang",8:"Benih Nila",9:"Benih Cupang",10:"Probiotik B"}

	//Deklarasi list harga
	var price = map[int]float64{1:50000 , 2:25000, 3:75000, 4:10000, 5:20000,
	6:20000, 7:5000, 8:30000, 9:10000, 10:10000}
	
	fmt.Println("===============")
	//1.a	Memanggil fungsi untuk menginformasikan barang yang dapat dibeli dengan kapasitas pembelian total (100k)
	fmt.Printf("Total produk dengan harga dibawah Rp.%.2f adalah:\n\n", pembelian_cap)
	list(pembelian_cap, barang, price)
	fmt.Println("===============")

	//1.b	Minimum price and product
	min_key, min_price, barang_min := murah(barang, price)
	_= min_key
	fmt.Printf("Harga terendah adalah Rp.%.2f untuk produk bernama %s\n", min_price, barang_min)

	//1.b	Maximum price and product
	maks_price, barang_maks := mahal(barang, price)
	fmt.Printf("Harga termahal adalah Rp.%.2f untuk produk bernama %s\n", maks_price, barang_maks)
	fmt.Println("===============")

	//1.c	Barang dengan harga 10k
	fmt.Println("Produk dengan harga Rp. 10000.00 adalah:\n")
	only_10k(barang, price)
}