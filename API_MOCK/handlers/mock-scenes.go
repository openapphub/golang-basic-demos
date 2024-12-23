package handlers

import (
	"encoding/json"
	"net/http"
	"openapphub/models"
	"strings"

	"github.com/brianvoe/gofakeit/v7"
)

func MockScenesHandler(w http.ResponseWriter, r *http.Request) {
	fake := gofakeit.New(0)

	response := models.ResponseScenesData{
		ImgUrl:  "/api/mock/image",
		Code:    0,
		Message: "Success",
		Data: []models.ScenesData{
			{
				// Name:   fake.Word(),
				Name:   "1",
				Result: strings.Replace(fake.Paragraph(2, 4, 10, "."), ". ", ".\n", -1),
			},
			{
				// Name:   fake.Word(),
				Name:   "2",
				Result: fake.Sentence(10),
			},
		},
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}
