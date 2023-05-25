package payload

type Response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type LoginResponse struct {
	Username string `json:"username"`
	Token    string `json:"token"`
}

type UpdateItemResponse struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Stock       uint   `json:"stock"`
	Price       uint   `json:"price"`
}
