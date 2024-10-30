package handlers

import (
	"net/http"

	"github.com/brianvoe/gofakeit/v7"
)

func MockImageHandle(w http.ResponseWriter, r *http.Request) {
	fake := gofakeit.New(0)

	w.Header().Set("Content-Type", "image/jpeg")
	w.Write(fake.ImageJpeg(800, 600))
}
