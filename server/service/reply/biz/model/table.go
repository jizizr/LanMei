package model

type ReplyTable struct {
	Code int    `json:"code"`
	Data Data   `json:"data"`
	Msg  string `json:"msg"`
}

type Data struct {
	Revision         int          `json:"revision"`
	SpreadsheetToken string       `json:"spreadsheetToken"`
	TotalCells       int          `json:"totalCells"`
	ValueRanges      []ValueRange `json:"valueRanges"`
}

type ValueRange struct {
	MajorDimension string     `json:"majorDimension"`
	Range          string     `json:"range"`
	Revision       int        `json:"revision"`
	Values         [][]string `json:"values"`
}

type TableMetaInfo struct {
	Code int       `json:"code"`
	Data TableData `json:"data"`
	Msg  string    `json:"msg"`
}
type TableData struct {
	Properties       Properties `json:"properties"`
	Sheets           []Sheet    `json:"sheets"`
	SpreadsheetToken string     `json:"spreadsheetToken"`
}
type Properties struct {
	OwnerUser  int64  `json:"ownerUser"`
	Revision   int    `json:"revision"`
	SheetCount int    `json:"sheetCount"`
	Title      string `json:"title"`
}
type Sheet struct {
	ColumnCount    int    `json:"columnCount"`
	FrozenColCount int    `json:"frozenColCount"`
	FrozenRowCount int    `json:"frozenRowCount"`
	Index          int    `json:"index"`
	RowCount       int    `json:"rowCount"`
	SheetId        string `json:"sheetId"`
	Title          string `json:"title"`
}
