package main

import (
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"
)

func getDescription() string {
	resp, err := http.Get("https://workspace-rho.vercel.app/api/description")
	if err != nil {
		// handle error
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	return string(body)
}

func getJobList() string {
	resp, err := http.Get("https://workspace-rho.vercel.app/api/jobs")
	if err != nil {
		// handle error
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	return string(body)
}

type Cache struct {
	DescString string
	JobList    string
	IsFilled   bool
}

// Deklarasi variabel untuk cache kosong
var zeroCache = &Cache{}

func (c *Cache) aggregate(i int) {
	wg := sync.WaitGroup{}

	//Number of goroutines for getDescription() and getJobList()
	wg.Add(2)

	//Waktu awal memulai
	start := time.Now()
	fmt.Println("dari calculate ", start)

	if !c.IsFilled {

		go func() {
			c.DescString = getDescription()
			wg.Done()
		}()

		go func() {
			c.JobList = getJobList()
			wg.Done()
		}()
		wg.Wait()

		c.IsFilled = true
	}

	//Print Description and JobList from API, we only need the runtime of the program, so we disable it
	//fmt.Println(c.DescString)
	//fmt.Println(c.JobList)

	//Waktu sejak operasi aggregate dimulai
	fmt.Printf("took %v\n", time.Since(start))

	//Saat aggregate dipanggil genap akan mengosongkan cache dengan zeroCache -> declared diatas
	if i%2 == 0 {
		*c = *zeroCache

		//Deklarasikan boolean  isFilled jadi false lagi for safe measure
		c.IsFilled = false
	}
}

func main() {

	//Declare cache
	var c Cache

	//Insert number of loops -> berapa kali aggregate() dipanggil
	fmt.Print("Masukkan jumlah looping: ")
	var n int
	fmt.Scanf("%d", &n)

	//For loops of calling method aggregate() n times
	for i := 1; i <= n; i++ {
		c.aggregate(i)
	}

	//Disaat ganjil nilai runtime akan lebih lama dibandingkan saat looping genap, karena pada looping genap sudah memiliki cache
	//Setelah looping genap selesai cache akan dikosongkan sehingga Cache c perlu diisi kembali

	//Contoh output
	// Masukkan jumlah looping: 6
	// dari calculate  2022-11-07 04:05:33.4636118 +0700 +07 m=+0.468313201
	// took 341.3058ms
	// dari calculate  2022-11-07 04:05:33.8049176 +0700 +07 m=+0.809619001
	// took 529.6µs
	// dari calculate  2022-11-07 04:05:33.8054472 +0700 +07 m=+0.810148601
	// took 257.0132ms
	// dari calculate  2022-11-07 04:05:34.0624604 +0700 +07 m=+1.067161801
	// took 561µs
	// dari calculate  2022-11-07 04:05:34.0630214 +0700 +07 m=+1.067722801
	// took 258.0837ms
	// dari calculate  2022-11-07 04:05:34.3216144 +0700 +07 m=+1.326315801
	// took 73.1µs

	//Berkat goroutine, fungsi aggregate tanpa cache sudah termasuk cepat (dalam milisecond)
	//Tetapi bisa lebih cepat bila local system memiliki memori untuk menyimpan value yang dipanggil berulang-ulang (mencapai microsecond)
}
