package entry

type Response struct {
	Id     string      `json:"id"`
	En     int         `json:"en"`
	Em     string      `json:"em"`
	Result interface{} `json:"result"`
}
