package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"openapphub/models"
	"os"
	"time"

	"github.com/brianvoe/gofakeit/v7"
)

func MockMainHandle(w http.ResponseWriter, r *http.Request) {

	// 1. 检查是否是POST请求
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// 2. 读取POST请求体
	var requestBody models.RequestMainData

	// 3. 解析JSON请求体
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&requestBody); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// 4. 将完整请求保存到文件
	requestJSON, _ := json.MarshalIndent(requestBody, "", "  ")
	filename := fmt.Sprintf("logs/request_%s.json", time.Now().Format("20060102_150405"))
	if err := os.MkdirAll("logs", 0755); err != nil {
		log.Printf("Failed to create logs directory: %v", err)
	}
	if err := os.WriteFile(filename, requestJSON, 0644); err != nil {
		log.Printf("Failed to write request to file: %v", err)
	}

	// 5. 打印请求参数摘要（不包含完整的base64数据）
	log.Printf("Received request, image length: %d bytes", len(requestBody.ImgBase64))
	log.Printf("Request saved to: %s", filename)

	// 6. 生成响应数据
	fake := gofakeit.New(0)

	mainDataList := make([]models.MainData, 4)
	for i := 0; i < 4; i++ {
		mainDataList[i] = models.MainData{
			ResultId:   fake.UUID(),
			CheckName:  fake.Username(),
			ResultCode: fake.RandomString([]string{"0", "1"}),
		}
	}

	response := models.ResponseMainData{
		RequestId:       fake.UUID(),
		ImageUrl:        "/api/mock/image",
		Code:            0,
		Message:         "success",
		CheckResultList: mainDataList,
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}
