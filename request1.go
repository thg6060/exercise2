package main

import (
	"fmt"
	"time"
)

//1. Tạo 1 chương trình cứ 3s in ra dòng chữ: `time now: {milliseconds}`:
//Sau khi in được 3 lần thì in ra dòng chữ `kết thúc`

func Run1() {


	now := time.Now()

	for i := 0; i < 3; i++ {
		fmt.Println("time now:", now.Format(time.StampMilli))
		time.Sleep(3* time.Second)	
	}
	fmt.Println("Ket Thuc")
}
