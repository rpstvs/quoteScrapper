package scrapper

import "github.com/abadojack/whatlanggo"

func LangDetection(s string) bool {
	quoteLang := whatlanggo.Detect(s)

	return quoteLang.Lang == whatlanggo.Eng

}
