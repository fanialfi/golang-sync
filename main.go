package main

import (
	"github.com/fanialfi/golang-sync/mutex"
	"github.com/fanialfi/golang-sync/wg"
)

func main() {
	wg.WG()
	mutex.Race()
	mutex.Run()
}
