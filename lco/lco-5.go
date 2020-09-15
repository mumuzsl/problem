package lco

import "strings"

func replaceSpace(s string) string {
	//r := ""
	//for _, v := range s {
	//	k := string(v)
	//	if k == " " {
	//		r += "%20"
	//	} else {
	//		r += k
	//	}
	//}
	//return r
	return strings.ReplaceAll(s, " ", "%20")
}
