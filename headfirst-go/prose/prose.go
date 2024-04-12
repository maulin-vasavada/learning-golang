package prose

import "strings"

func JoinWithSeparator(phrases []string, separator string) string {
	result := strings.Join(phrases[:len(phrases)-1], separator+" ")
	result += ", and "
	result += phrases[len(phrases)-1]
	return result
}
