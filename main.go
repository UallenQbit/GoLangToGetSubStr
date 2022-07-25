package GoLangToGetSubStr

import "strings"

func GetSubStr(String string, LeftString string, RightString string) (SubString string) {
	if LeftIndex := strings.Index(String, LeftString); LeftIndex == -1 {
		return ""
	} else {
		LeftIndex = LeftIndex + len(LeftString)
		RightSubString := String[LeftIndex:]
		if RightIndex := strings.Index(RightSubString, RightString); RightIndex == -1 {
			return ""
		} else {
			return RightSubString[:RightIndex]
		}
	}
}
