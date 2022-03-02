package parameters

type User struct {
	UUID   string `json:"uuid" form:"uuid"`
	Name   string `json:"name" form:"name"`
	Age    int    `json:"age" form:"age"`
	Gender string `json:"gender" form:"gender"`
}
