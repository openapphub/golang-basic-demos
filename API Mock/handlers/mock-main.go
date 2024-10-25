package handlers

import (
	"encoding/json"
	"math/rand/v2"
	"net/http"
	"openapphub/models"
	"strconv"
)

func MockMainhandle(w http.ResponseWriter, r *http.Request) {
	response := models.ResponseMainData{
		RequestId: strconv.Itoa(rand.Int()),
		ImageUrl:  string("http"),
		Code:      0,
		Message:   string("success"),
		Data: []models.MainData{
			{
				ResultId:  "123231",
				CheckName: "mnopqrstuvwxyzABCDEFGH",
				Result:    "0",
			},
			{
				ResultId:  "123231",
				CheckName: "mnopqrstuvwxyzABCDEFGH",
				Result:    "0",
			}, {
				ResultId:  "123231",
				CheckName: "mnopqrstuvwxyzABCDEFGH",
				Result:    "0",
			},
			{
				ResultId:  "123231",
				CheckName: "mnopqrstuvwxyzABCDEFGH",
				Result:    "0",
			},
		},
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}
