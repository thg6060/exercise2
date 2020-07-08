package main

import (
	"fmt"
	"time"
)

//. Trong golang mặc định thì thời gian dạng số được sử dụng với các loại mốc đơn vị nào?
func Run6() {
	//Golang dung nanosecond,second,year,Day là dạng số
	now := time.Now()
	fmt.Println(now.Year())
	fmt.Println(now.Month())
	fmt.Println(now.Second())
	fmt.Println(now.Nanosecond())
}
