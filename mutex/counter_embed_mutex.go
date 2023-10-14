package mutex

import "sync"

type counter struct {
	sync.Mutex     // embed sync.Mutex untuk melakukan Lock dan Unlock data agar terhindar dari race condition
	val        int // properti ini akan di konsumsi dan di olah oleh banyak goruntine
}

// Add digunakan untuk increment nilai pada properti val
func (c *counter) Add(int) {
	c.Lock() // digunakan untuk menandai bahwa semua operasi pada baris setelahnya adalah eksklusif,
	// hanya ada 1 buah goruntine saja yang bisa melakukannya dalam 1 waktu

	c.val++

	c.Unlock() // digunakan untuk membuka kembali akses operasi ke properti / variabel yang di lock
	// proses mutual exclusife terjadi di antara method Lock() dan Unlock()
}

// Value digunakan untuk mengembalikan nilai pada properti val
//
// pada method pengambilan nilai val tidak di pasang mutex
// karena pada implementasi struct counter
// pengambilan nilai terjadi setelah semua goruntine selesai
// karena data race bisa terjadi saat pengubahan nilai maupun pengambilan nilai
func (c *counter) Value() int {
	return c.val
}
