package main

import (
	"fmt"
	"time"
)

func main() {

	c := time.After(5 * time.Second) //5 snlik timer oluşturuldı
	for {
		b := false

		select {
		case <-c: //kanalı dinle anlamına gelir, buraya bir değer geldiği zaman...
			b = true //5 sn sonra timer a değer basıyor
		default:
			fmt.Println(time.Now())
			time.Sleep(1 * time.Second)
		}

		if b {
			break
		}
	}

}
