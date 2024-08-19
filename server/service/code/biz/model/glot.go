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
