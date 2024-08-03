package model

type TableStyle struct {
	AppendStyle AppendStyle `json:"appendStyle"`
}

type AppendStyle struct {
	Range string `json:"range"`
	Style Style  `json:"style"`
}

type Style struct {
	BackColor string `json:"backColor"`
}

type StyleResp struct {
	Code int       `json:"code"`
	Data StyleData `json:"data"`
	Msg  string    `json:"msg"`
}

type StyleData struct {
	Updates Updates `json:"updates"`
}

type Updates struct {
	Revision         int    `json:"revision"`
	SpreadsheetToken string `json:"spreadsheetToken"`
	UpdatedCells     int    `json:"updatedCells"`
	UpdatedColumns   int    `json:"updatedColumns"`
	UpdatedRange     string `json:"updatedRange"`
	UpdatedRows      int    `json:"updatedRows"`
}
