package main

import (
	"fmt"
	"time"
)
//2. Việt 1 đoạn chương trình tính ra ngày hiện tại theo timestamp.
//Điểm mốc từ mức 0  tại 1/1/1970

func Run2() {
	now := time.Now()
	currentDay := time.Unix(now.Unix(), 0)
	fmt.Println(currentDay)
}
