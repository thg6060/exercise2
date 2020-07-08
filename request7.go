package main

import (
	"context"
	"fmt"
	"time"
)

type favContextKey string

//Tạo 1 đối tượng context với 1 value là timestamp hiện tại tính theo `ns`
//chạy qua 1 hàm như sau. Sau 3s in ra hiệu của thời gian của thời gian hiện tại với biến dữ liệu trong context. in ra màn hình kết quả.

func newContext(ctx context.Context, name string, val int64) (newCtx context.Context) {

	k := favContextKey(name)
	ctx = context.WithValue(context.Background(), k, val)
	return ctx
}
func Run7() {
	timestamp := time.Now().UnixNano()
	ctx := context.Background()
	newCtx := newContext(ctx, "tstamp", timestamp)

	time.Sleep(time.Second * 3)

	timestamp = time.Now().UnixNano()
	fmt.Println("Sub time after 3s ", timestamp-newCtx.Value(favContextKey("tstamp")).(int64))
}
