package util

import (
	"fmt"
	"github.com/jizizr/LanMei/server/common"
	"github.com/jizizr/LanMei/server/service/code/biz/model"
	"github.com/jizizr/LanMei/server/service/code/conf"
)

const baseUrl = "https://glot.io/api/run"

var client = common.DefaultHttpReq(baseUrl).SetCommonHeaders(map[string]string{
	"Content-Type":  "application/json",
	"Authorization": fmt.Sprintf("Token %s", conf.GetConf().Glot.Token),
})

var CodeType = map[string][]string{
	"py":           {"python", "py"},
	"cpp":          {"cpp", "cpp"},
	"java":         {"java", "java"},
	"php":          {"php", "php"},
	"js":           {"javascript", "js"},
	"c":            {"c", "c"},
	"c#":           {"csharp", "cs"},
	"go":           {"go", "go"},
	"asm":          {"assembly", "asm"},
	"ats":          {"ats", "dats"},
	"bash":         {"bash", "sh"},
	"clisp":        {"clisp", "lsp"},
	"clojure":      {"clojure", "clj"},
	"cobol":        {"cobol", "cob"},
	"coffeescript": {"coffeescript", "coffee"},
	"crystal":      {"crystal", "cr"},
	"D":            {"D", "d"},
	"elixir":       {"elixir", "ex"},
	"elm":          {"elm", "elm"},
	"erlang":       {"erlang", "erl"},
	"fsharp":       {"fsharp", "fs"},
	"groovy":       {"groovy", "groovy"},
	"guile":        {"guile", "scm"},
	"hare":         {"hare", "ha"},
	"haskell":      {"haskell", "hs"},
	"idris":        {"idris", "idr"},
	"julia":        {"julia", "jl"},
	"kotlin":       {"kotlin", "kt"},
	"lua":          {"lua", "lua"},
	"mercury":      {"mercury", "m"},
	"nim":          {"nim", "nim"},
	"nix":          {"nix", "nix"},
	"ocaml":        {"ocaml", "ml"},
	"pascal":       {"pascal", "pp"},
	"perl":         {"perl", "pl"},
	"raku":         {"raku", "raku"},
	"ruby":         {"ruby", "rb"},
	"rust":         {"rust", "rs"},
	"sac":          {"sac", "sac"},
	"scala":        {"scala", "scala"},
	"swift":        {"swift", "swift"},
	"typescript":   {"typescript", "ts"},
	"zig":          {"zig", "zig"},
	"plaintext":    {"plaintext", "txt"},
}

func Run(code string, language string) (string, error) {
	glotReq := model.GlotReq{
		Files: []model.File{
			{
				Name:    fmt.Sprintf("main.%s", CodeType[language][1]),
				Content: code,
			},
		},
	}
	var glotResp model.GlotResp
	r, err := client.R().SetSuccessResult(&glotResp).
		SetBody(glotReq).
		Post(fmt.Sprintf("/%s/latest", CodeType[language][0]))
	if err != nil {
		return "", err
	}
	if !r.IsSuccessState() {
		return "处理失败", nil
	}
	if glotResp.Stdout != "" {
		return glotResp.Stdout, nil
	} else {
		return glotResp.Stderr, nil
	}
}
