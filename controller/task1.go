package controller

import (
	"fmt"
	"strconv"
	"sync"
)

type Deret struct {
	prima       int
	ganjil      int
	genap       int
	fibonacci   int
	ganjilgenap []int
	mtx         sync.Mutex
}

func DeretBilangan(x int) {
	chn1 := make(chan string)
	chn2 := make(chan string)
	chn3 := make(chan string)
	chn4 := make(chan string)

	wg.Add(4)

	go func() {
		BilPrima := Deret{prima: x}
		BilPrima.GetPrima(chn1)
	}()

	go func() {
		BilGanjil := Deret{ganjil: x}
		BilGanjil.GetGanjil(chn2)
	}()

	go func() {
		BilGenap := Deret{genap: x}
		BilGenap.GetGenap(chn3)
	}()

	go func() {
		BilFibonacci := Deret{fibonacci: x}
		BilFibonacci.GetFibonacci(chn4)
	}()
	fmt.Println("Bilangan Prima :", <-chn1)
	fmt.Println("Bilangan ganjil :", <-chn2)
	fmt.Println("Bilangan Genap  :", <-chn3)
	fmt.Println("Bilangan Fibonacci :", <-chn4)

	wg.Wait()

}

func (d *Deret) GetPrima(chn chan string) {
	defer wg.Done()
	x := d.prima
	var result string
	for bil := 1; bil < x; bil++ {
		i := 0
		for bag := 1; bag < x; bag++ {
			if bil%bag == 0 {
				i++
			}
		}
		if (i == 2) && (bil != 1) {
			result += strconv.Itoa(bil) + " "
		}
	}
	chn <- result

}

func (value *Deret) GetGanjil(chn chan string) {
	defer wg.Done()
	x := value.ganjil
	var result string

	for i := 0; i < x; i++ {
		if i%2 == 1 {
			result += strconv.Itoa(i) + " "

		}
	}
	chn <- result
}
func (value *Deret) GetGenap(chn chan string) {
	defer wg.Done()
	x := value.genap
	var result string
	for i := 0; i < x; i++ {
		if i%2 == 0 {
			result += strconv.Itoa(i) + " "
		}
	}
	chn <- result
}
func (value *Deret) GetFibonacci(chn chan string) {
	defer wg.Done()
	num := value.fibonacci
	var result string
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
		result += strconv.Itoa(b) + " "
	}
	chn <- result
}
