package main

import (
	"fmt"
	"sync"
	"time"
)

func esperar() {
	var input string
	fmt.Scanln(&input)
}

func main() {
	m := new(sync.Mutex)

	for i := 0; i < 10; i++ {
		go func(i int) {
			m.Lock() // exclusivo acceso a la SC
			fmt.Println(i, "SC line1")
			time.Sleep(time.Second)
			fmt.Println(i, "SC line2")
			m.Unlock()
		}(i)
	}

	esperar()
}
