package tools

import (
	"fmt"
	"time"
)

func Delay(delay int) {
	fmt.Println("Screenshot delayed...")
	for i := 1; i < delay+1; i++ {
		fmt.Println(i)
		time.Sleep(time.Duration(1) * time.Second)
	}
}
