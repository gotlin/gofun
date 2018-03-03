package LongestCommonPrefix

func LongestCommonPrefix(strs []string) string {

	row := len(strs)

	if row == 0 {
		return ""
	}

	if row == 1 {
		return strs[0]
	}

	minLen := row

	for _, v := range strs {
		if len(v) < minLen {
			minLen = len(v)
		}
	}

	res := make([]rune, 0)

	for i, char := range strs[0] {

		comm := true
		for _, v := range strs[1:] {

			if chars := []rune(v); len(chars) < i+1 || char != chars[i] {
				comm = false
				break
			}

		}

		if comm {
			res = append(res, char)
		} else {
			break
		}

	}

	return string(res)

}
