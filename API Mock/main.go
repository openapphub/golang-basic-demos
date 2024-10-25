package main

import (
	"fmt"
	"net/http"
	"openapphub/handlers"
)

func main() {
	http.HandleFunc("/api/mock-scenes", handlers.MockScenesHandler)
	http.HandleFunc("/api/mock-main", handlers.MockMainhandle)
	// add more
	fmt.Println("Starting server at port 8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
