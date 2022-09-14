package controller

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func Fibonacci() {
	chn1 := make(chan []int, 2)
	chn2 := make(chan int, 2)
	wg.Add(2)
	x := 40

	go func() {
		BilFibonacci := Deret{fibonacci: x}
		BilFibonacci.GetNumFibonacci(chn1)
	}()

	go func() {
		GanjilGenap := Deret{ganjilgenap: <-chn1}
		GanjilGenap.Getganjigenap(chn2)
	}()

	fmt.Println(<-chn2)
	wg.Wait()
}

func (value *Deret) GetNumFibonacci(chn chan []int) {
	defer wg.Done()

	num := value.fibonacci
	var result []int
	a := 0
	b := 1
	c := b

	for true {
		c = b
		b = a + b
		if b >= num {
			break
		}
		a = c
		result = append(result, a)
	}

	chn <- result

}

func (value *Deret) Getganjigenap(chn chan int) {
	defer wg.Done()
	row := value.ganjilgenap
	var result int
	value.mtx.Lock()
	for i := 0; i < len(row); i++ {
		if row[i]%2 == 0 {
			result = row[i]
			fmt.Println("Genap", result)

		}
		if row[i]%2 == 1 {
			result = row[i]
			fmt.Println("Ganjil", result)

		}
	}
	value.mtx.Unlock()
	chn <- result

}
