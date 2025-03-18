package piscine

func PrintIfNot(str string) string {
	if len(str) > 3 || len(str) == 0 {
		return "Invalid Input\n"
	} else {
		return "G\n"
	}
}
