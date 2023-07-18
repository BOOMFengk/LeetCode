func replaceSpace(s string) string {
	str := ""
	for _, v := range s {
		if string(v) != " " {
			str += string(v)
		} else {
			str += "%20"
		}

	}
	return str

}