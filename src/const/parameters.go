package _const

type Status struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Stack   []string    `json:"stack"`
	Data    interface{} `json:"data"`
}

type TestRequest struct {
	UUID   string `json:"uuid" form:"uuid"`
	Name   string `json:"name" form:"name"`
	Age    int    `json:"age" form:"age"`
	Gender string `json:"gender" form:"gender"`
}

type TestResponse struct {
	Status *Status `json:"status"`
}
