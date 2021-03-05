package ansi

import "regexp"

// const reg = "[\u001B\u009B]\\[.*?m"

// based on chalk/ansi-regex
const reg = "[\u001B\u009B][[\\]()#;?]*(?:(?:(?:[a-zA-Z\\d]*(?:;[a-zA-Z\\d]*)*)?\u0007)|(?:(?:\\d{1,9}(?:;\\d{0,9})*)?[\\dA-PRZcf-ntqry=><~]))"

var re = regexp.MustCompile(reg)

// Strip strip all ANSI escape sequences.
// given a string, this will strip all ANSI escape codes.
func Strip(str string) string {
	return re.ReplaceAllString(str, "")
}
