package model

type MusicInfo struct {
	Data MusicData `json:"data"`
}

type MusicData struct {
	Id       int    `json:"id"`
	Mid      string `json:"mid"`
	Vid      string `json:"vid"`
	Song     string `json:"song"`
	Subtitle string `json:"subtitle"`
	Singer   string `json:"singer"`
	Album    string `json:"album"`
	Pay      string `json:"pay"`
	Time     string `json:"time"`
	Bpm      int    `json:"bpm"`
	Quality  string `json:"quality"`
	Interval string `json:"interval"`
	Size     string `json:"size"`
	Kbps     string `json:"kbps"`
	Cover    string `json:"cover"`
	Link     string `json:"link"`
	Url      string `json:"url"`
}
