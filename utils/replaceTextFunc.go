package utils

import "strings"

func TrimOtherChar(message string) string {
	message = strings.Replace(message, "{", "", 1)
	message = strings.Replace(message, "}", "", 1)
	mapList := map[string]string{
		"『":  "",
		"』":  "",
		"「":  "",
		"」":  "",
		"\n": "",
		`"`:  "",
		`]`:  "",
		`[`:  "",
		` `:  "",
		`　`:  "",
		`、`:  "",
		`。`:  "",
		`(`:  "",
		`)`:  "",
		`）`:  "",
		`（`:  "",
		`》`:  "",
		`《`:  "",
		`×`:  "",
		`・`:  "",
		`”`:  "",
		`“`:  "",
		`’`:  "",
		`‘`:  "",
		`：`:  "",
		`；`:  "",
		`｝`:  "",
		`｛`:  "",
		`＜`:  "",
		`＞`:  "",
		`＿`:  "",
		`＃`:  "",
		`＄`:  "",
		`％`:  "",
		`＆`:  "",
		`＝`:  "",
		`～`:  "",
		`＾`:  "",
		`￥`:  "",
		`｜`:  "",
		`☆`:  "",
		`…`:  "",
	}
	for before, after := range mapList {
		message = strings.ReplaceAll(message, before, after)
	}

	return message
}

func TrimNumKanji(trimedHiragana string) string {
	//yahooでは対応していない文字
	mapList := map[string]string{
		"数十":       "すうじゅう",
		"数百":       "すうひゃく",
		"数千":       "すうせん",
		"数万":       "すうまん",
		"数十万":      "すうじゅうまん",
		"数億":       "すうおく",
		"数兆":       "すうちょう",
		"一":        "いち",
		"二":        "に",
		"三":        "さん",
		"四":        "よん",
		"五":        "ご",
		"六":        "ろく",
		"七":        "なな",
		"八":        "はち",
		"九":        "きゅう",
		"十":        "じゅう",
		"百":        "ひゃく",
		"千":        "せん",
		"万":        "まん",
		"億":        "おく",
		"兆":        "ちょう",
		"数":        "すう",
		"\n":       "",
		"わざはぐくみてん": "ぎいくてん",
		"おにめつのは":   "きめつのやいば",
		"ゼロ":       "ぜろ",
		"すみじろう":    "たんじろう",
	}
	for before, after := range mapList {
		trimedHiragana = strings.ReplaceAll(trimedHiragana, before, after)
	}
	return trimedHiragana
}
