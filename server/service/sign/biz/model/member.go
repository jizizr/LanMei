package model

type Member struct {
	Status  string `json:"status"`
	Retcode int    `json:"retcode"`
	Data    struct {
		Nickname string `json:"nickname"`
		Card     string `json:"card"`
	} `json:"data"`
}
