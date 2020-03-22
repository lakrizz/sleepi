package stringhelper

func Map(s []string, fun func(s string) string) []string {
	rs := make([]string, 0)
	for _, v := range s {
		rs = append(rs, fun(v))
	}
	return rs
}
