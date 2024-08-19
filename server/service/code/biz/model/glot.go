package model

type GlotReq struct {
	Files []File `json:"files"`
}

type File struct {
	Name    string `json:"name"`
	Content string `json:"content"`
}

type GlotResp struct {
	Stdout string `json:"stdout"`
	Error  string `json:"error"`
	Stderr string `json:"stderr"`
}

type Message struct {
	Type string `json:"type"`
	Data Data   `json:"data"`
}

type Data struct {
	Uin     int64     `json:"uin"`
	Content []Content `json:"content"`
}

type Content struct {
	Type string      `json:"type"`
	Data ContentData `json:"data"`
}

type ContentData struct {
	Text string `json:"text"`
}
