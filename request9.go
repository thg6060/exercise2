package main

import (
	"fmt"
	"time"
)

func Run9() {
//Tạo 1 đoạn code sử dụng `time.AfterFunc()`, sau 100ms thì in ra dòng chữ `i'm study`
	durationtime := time.Duration(100) * time.Millisecond

	Timer := time.AfterFunc(durationtime, func() {
		fmt.Println("i'm study")
	})
	defer Timer.Stop()

	time.Sleep(3 * time.Second)

}
