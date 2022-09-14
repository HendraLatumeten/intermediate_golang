package controller

import (
	"fmt"
	"sync"
)

var mtx sync.Mutex

func GetDiskon() {
	ch := make(chan int, 3)
	wg.Add(3)
	go checkDiskon(ch, "10%", 700.000)
	go checkDiskon(ch, "10%", 300.000)
	go checkDiskon(ch, "50%", 500.000)

	fmt.Println("diskon dari 300.000", "adalah", <-ch, ".000")
	fmt.Println("diskon dari 700.000", "adalah", <-ch, ".000")
	fmt.Println("diskon dari 500.000", "adalah", <-ch, ".000")

	wg.Wait()

}

func checkDiskon(ch chan int, diskon string, harga int) {
	defer wg.Done()
	mtx.Lock()
	if diskon == "10%" {
		disc := harga * 10 / 100
		result := harga - disc
		ch <- result
		//	close(ch)
	} else if diskon == "50%" {
		disc := harga * 50 / 100
		result := harga - disc
		ch <- result
	}
	mtx.Unlock()

}
