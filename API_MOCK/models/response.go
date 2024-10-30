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

// RequestMainData 主要接口请求数据
type RequestMainData struct {
	ScenceCode string  `json:"scenceCode"`
	ProjectId  string  `json:"projectId"`
	TaskId     string  `json:"taskId"`
	ObjectType string  `json:"objectType"`
	ObjectId   string  `json:"objectId"`
	ImgBase64  string  `json:"imgBase64"`
	Longitude  float64 `json:"longitude"`
	Latitude   float64 `json:"latitude"`
	PhotoTime  string  `json:"photoTime"` // 2024-06-24 10:23:22
}

/* 测试入参
{
	"scenceCode": "SC12345",
	"projectId": "PRJ67890",
	"taskId": "TSK54321",
	"objectType": "Building",
	"objectId": "OBJ98765",
	"imgBase64": "data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAA...",
	"longitude": 123.456789,
	"latitude": 34.567890,
	"photoTime": "2024-06-24 10:23:22"
}
*/

type ResponseMainData struct {
	RequestId       string     `json:"requestId"`
	ImageUrl        string     `json:"imageUrl"`
	Code            int        `json:"code"`
	Message         string     `json:"message"`
	CheckResultList []MainData `json:"checkResultList"`
}
type MainData struct {
	ResultId   string `json:"resultId"`
	CheckName  string `json:"checkName"`
	ResultCode string `json:"resultCode"` // 0 / 1
}
