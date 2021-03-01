package responses

type MyRespon struct {
	Messages string      `json:"messages"`
	Result   interface{} `json:"result"`
}
