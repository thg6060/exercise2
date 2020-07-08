package main

import (
	"fmt"
	"time"
)
//5. Tính ra số ngày trong tuần(dạng string và number) của mốc thời gian sau `1592190385` second
func Run5() {
	stimestamp := 1592190385
	currentday := time.Unix(int64(stimestamp), 0)
	fmt.Println("Day of timestamp(1592190385) ", currentday)
	fmt.Println("Today is (string) ", currentday.Weekday())
	fmt.Println("Today is (number) ", int(currentday.Weekday()))

}
