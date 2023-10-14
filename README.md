# sync

## sync.WaitGroup

go menyediakan package `sync` yang berisi cukup banyak API untuk handle masalah multiprocessing (goruntine),
salah satunya adalah `sync.WaitGroup`.

kegunaan `WaitGroup` sendiri adalah untuk sinkronasi goruntine, berbeda dengan `channel`, `sync.WaitGroup` memang dirancang untuk maintaince goruntine

`sync.WaitGroup` merupakan salah satu tipe _thread safe_, kita tidak perlu khawatir terhadap potensi race condition, karena variabel tipe ini aman untuk digunakan di banyak goruntine secara paralel

## sync.Mutex

race condition adalah sebuah kondisi dimana lebih dari satu goruntine mengakses data yang sama pada waktu yang sama juga (benar benar bersama).
Ketika hal ini terjadi, nilai data tersebut menjadi kacau.
Dalam **concurency programming** situasi seperti ini sering terjadi.

**mutex** melakukan pengubahan level akses data menjadi eksklusif,
menjadikan data tersebut hanya dapat di konsumsi (read / write) oleh satu buah goruntine saja.
Jadi ketika terjadi _race condition_ hanya goruntine yang beruntung saja yang dapat mengakses data tersebut, 
goruntine lain (yang waktu running kebetulan bersamaan) akan dipaksa menunggu sampai goruntine yang memanfaatkan data tersebut selesai.

untuk mendeteksi race condition, go menyediakan fitur untuk [deteksi race condition][race], caranya adalah dengan menambahkan flag `-race` pada saat eksekusi aplikasi.

[race]: http://blog.golang.org/race-detector