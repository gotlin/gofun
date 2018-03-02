package LongestCommonPrefix

func LongestCommonPrefix(strs []string) string {

	res := make([]rune, 1)

	for i, char := range strs[0] {
		comm := true
		for _, v := range strs[1:] {

			chars := []rune(v)
			if char != chars[i] {
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
