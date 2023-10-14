# sync

go menyediakan package `sync` yang berisi cukup banyak API untuk handle masalah multiprocessing (goruntine),
salah satunya adalah `sync.WaitGroup`.

kegunaan `WaitGroup` sendiri adalah untuk sinkronasi goruntine, berbeda dengan `channel`, `sync.WaitGroup` memang dirancang untuk maintaince goruntine

`sync.WaitGroup` merupakan salah satu tipe _thread safe_, kita tidak perlu khawatir terhadap potensi race condition, karena variabel tipe ini aman untuk digunakan di banyak goruntine secara paralel