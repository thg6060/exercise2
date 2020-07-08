package main

import "fmt"

func Run4() {
	//4. Tính ra số phút(number_of_minutes) của mốc thời gian sau `1592190294764144364` nano
	ntimestamp := 1592190294764144364
	number_of_minutes := ntimestamp / 10e9 / 60
	fmt.Println(number_of_minutes)
}
