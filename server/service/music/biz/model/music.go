package model

type Ck struct {
	Data string `json:"data"`
}

type CkResp struct {
	Result int    `json:"result"`
	Data   string `json:"data"`
}

type MusicInfo struct {
	Result int       `json:"result"`
	Data   MusicData `json:"data"`
}

type singer struct {
	Name string `json:"name"`
}

type list struct {
	MediaMid    string   `json:"media_mid"`
	PubTime     int64    `json:"pubtime"`
	Singers     []singer `json:"singer"`
	SongMid     string   `json:"songmid"`
	SongName    string   `json:"songname"`
	StrMediaMid string   `json:"strMediaMid"`
}

type MusicData struct {
	Lists []list `json:"list"`
}

type MusicStream struct {
	Data   string `json:"data"`
	Result int    `json:"result"`
}
