package prose

import "strings"

func JoinWithSeparator(phrases []string, separator string) string {
	if len(phrases) == 1 {
		return phrases[0]
	} else if len(phrases) == 2 {
		return phrases[0] + " and " + phrases[1]
	} else {
		result := strings.Join(phrases[:len(phrases)-1], separator+" ")
		result += ", and "
		result += phrases[len(phrases)-1]
		return result
	}
}
