package main

import (
	"fmt"
	"time"
)

func Run8() {
	//Tạo 1 interval time sao cho cứ 100ms in ra dùng chữ `${time.Now().Unix()} done`
	var Timer *time.Timer

	for i := 1; i <= 10; i++ {
		durationtime := time.Duration(100*i) * time.Millisecond
		Timer = time.AfterFunc(durationtime, func() {
			fmt.Println("Now is: ", time.Now().Unix())
		})

	}
	defer Timer.Stop()
	time.Sleep(30 * time.Second)
}
