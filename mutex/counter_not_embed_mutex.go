// berikut ini merupakan contoh penggunaan mutex tanpa harus di embed ke object instance yang memerlukan lock-unlocking
package mutex

import (
	"fmt"
	"runtime"
	"sync"
)

type pertambahan struct {
	val int
}

func (p *pertambahan) Add(int) {
	p.val++
}

func (p *pertambahan) Value() int {
	return p.val
}

func Run() {
	runtime.GOMAXPROCS(2)
	var wg sync.WaitGroup
	var meter pertambahan
	var mtx sync.Mutex

	jumplahPerulangan := 1000

	for i := 0; i < jumplahPerulangan; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			for i := 0; i < jumplahPerulangan; i++ {

				mtx.Lock()
				meter.Add(1)
				mtx.Unlock()
			}
		}()
	}
	wg.Wait()
	fmt.Println("run", meter.Value())
}
