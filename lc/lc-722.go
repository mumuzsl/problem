package lc

func removeComments(source []string) []string {
	ans := make([]string, 0)
	cur, status := "", "str"
	for _, s := range source {
		for _, ch := range s {
			if status == "str" {
				if ch == '/' {
					status = "pre"
				} else {
					cur += string(ch)
				}
			} else if status == "pre" {
				if ch == '/' {
					status = "str"
					break
				} else if ch == '*' {
					status = "block"
				} else {
					status = "str"
					cur += string('/')
					cur += string(ch)
				}
			} else if status == "block" && ch == '*' {
				status = "block_end_pre"
			} else if status == "block_end_pre" {
				if ch == '/' {
					status = "str"
				} else if ch != '*' {
					status = "block"
				}
			}
		}
		if status == "pre" {
			cur += string('/')
			status = "str"
		} else if status == "block_end_pre" {
			status = "block"
		}
		if len(cur) != 0 && status == "str" {
			ans = append(ans, cur)
			cur = ""
		}
	}
	return ans
}
