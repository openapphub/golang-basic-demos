package main

import (
	"fmt"
	"net/http"
	"openapphub/handlers"
)

func main() {
	// 场景说明
	http.HandleFunc("/api/mock-scenes", handlers.MockScenesHandler)
	// 主要接口
	http.HandleFunc("/api/mock-main", handlers.MockMainHandle)
	// 随机图片
	http.HandleFunc("/api/mock/image", handlers.MockImageHandle)

	// add more
	fmt.Println("Starting server at port 8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
