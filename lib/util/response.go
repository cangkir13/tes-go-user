package util

type Response struct {
	Status bool        `json:"success"`
	Error  interface{} `json:"error"`
	Data   interface{} `json:"data"`
}
