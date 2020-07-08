package main

import (
	"context"
	"fmt"
	"time"
)

//3. Thực hiện 1 chương trình với 1 vòng lặp for và 3 lần sleep mỗi lần sleep 3sec
//Nhưng sau 3s thì kết thúc hàm đấy
//Sử dụng và tìm hiểu context. Nêu được tác dụng của context trong chương trình.

func Run3() error{
	//Cho deadline là 3s kể từ thời điểm hiện tại

	deadline := time.Now().Add(3 * time.Second)
	ctx, cancel := context.WithDeadline(context.Background(), deadline)
	defer cancel()

	for i := 1; i <= 3; i++ {
		select {
		//Sau khi hết duration deadline context hoàn thành
		case <-ctx.Done():
			return ctx.Err()

		default:
			//Chưa hết deadline thì nhảy vào default
			time.Sleep(3 * time.Second)
			fmt.Printf("Context %ds\n", i*3)
		}

	}
	return nil
//Tác dụng của context ở đây tạo 1 deadline để hủy context tại một thời điểm xác định trong tương lai
}
