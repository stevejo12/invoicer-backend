package models

type JsonResponse struct {
	Code    int           `json:"code"`
	Data    []interface{} `json:"data"`
	Message string        `json:"message"`
}
