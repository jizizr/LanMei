package util

import (
	"github.com/go-ego/gse"
	"github.com/go-ego/gse/hmm/pos"
	"unicode/utf8"
)

var seg gse.Segmenter
var posSeg pos.Segmenter
var wordClass = map[string]struct{}{"v": {}, "l": {}, "n": {}, "nr": {}, "a": {}, "vd": {}, "nz": {}, "PER": {}, "f": {}, "ns": {}, "LOC": {}, "s": {}, "nt": {}, "ORG": {}, "nw": {}, "vn": {}}

func init() {
	err := seg.LoadDict("./data/s_1.txt, ./data/t_1.txt")
	if err != nil {
		panic(err)
	}
	posSeg.WithGse(seg)
}

func Cut(sentence string) map[string]uint {
	poss := posSeg.Cut(sentence, true)
	words := make(map[string]uint)
	for _, po := range poss {
		if _, ok := wordClass[po.Pos]; !ok {
			continue
		}
		if utf8.RuneCountInString(po.Text) < 2 {
			continue
		}
		words[po.Text]++
	}
	return words
}
