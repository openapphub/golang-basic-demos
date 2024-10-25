package handlers

import (
	"encoding/json"
	"net/http"
	"openapphub/models"
)

func MockScenesHandler(w http.ResponseWriter, r *http.Request) {
	response := models.ResponseScenesData{
		ImgUrl:  string("https://www.baidu.png"),
		Code:    int(0),
		Message: string("Success"),
		Data: []models.ScenesData{
			{
				Name:   string("图片质量要求"),
				Result: string("图片要求说明"),
			},
			{
				Name:   "场景要求",
				Result: "场景要求说明",
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
