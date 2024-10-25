package models

// ResponseDate JSON

type ResponseScenesData struct {
	ImgUrl  string       `json:"imgUrl"`
	Code    int          `json:"code"`
	Message string       `json:"message"`
	Data    []ScenesData `json:"data"`
}
type ScenesData struct {
	Name   string `json:"name"`
	Result string `json:"result"`
}

type ResponseMainData struct {
	RequestId string     `json:"requestId"`
	ImageUrl  string     `json:"imageUrl"`
	Code      int        `json:"code"`
	Message   string     `json:"message"`
	Data      []MainData `json:"data"`
}
type MainData struct {
	ResultId  string `json:"resultId"`
	CheckName string `json:"checkName"`
	Result    string `json:"result"` // 0 / 1
}
