package mutex

import (
	"fmt"
	"runtime"
	"sync"
)

// pastikan jumplah core pada processor lebih dari satu
// GOMAXPROCS > 1

// Race merupakan contoh program yang didalamnya kemungkinan terjadi race condition
func Race() {
	runtime.GOMAXPROCS(2) // membatasi jumplah cpu yang digunakan untuk eksekusi program sebanyak 2

	// wg digunakan untuk sinkronasi antar goruntine
	var wg sync.WaitGroup

	var meter counter

	jumplahPerulangan := 1000

	for i := 0; i < jumplahPerulangan; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			for j := 0; j < jumplahPerulangan; j++ {
				meter.Add(1)
			}

		}()
	}

	// ekspektasi nilai akhir meter.val = 1000000
	wg.Wait()
	fmt.Println("race", meter.Value())
}

// pada contoh diatas penggunaan mutex diterapkan dengan cara embed struct Mutex kedalam object instance counter,
// dimana object instance counter ini memerlukan proses lock-unlocking,
// menjadikan object instance mutex tersebut ekskusif untuk object instance counter itu saja
// cara ini juga merupakan cara yang dianjurkan
