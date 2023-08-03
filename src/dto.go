package src

type Request struct {
	Url string `json:"url"`
}

type Response struct {
	TinyUrl string `json:"tinyurl"`
}
