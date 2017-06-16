// Package nlbr provides nl to br and br to nl.
package nlbr

import (
	"strings"
)

// Nl2br returns string with <br> inserted before all newlines (\r\n, \n\r, \n and \r).
func Nl2br(s string) string {
	r := strings.NewReplacer(
		"\r", "<br>\r",
		"\n", "<br>\n",
		"<br>\n<br>\r", "<br>\n\r",
		"<br>\r<br>\n", "<br>\r\n")
	return r.Replace(s)
}

// RevNl2br deletes br tag
func RevNl2br(s string) string {
	r := strings.NewReplacer("<br>", "")
	return r.Replace(s)
}

// Br2nl converts br tags to nl
func Br2nl(s string) string {
	r := strings.NewReplacer("<br>", "\n")
	return r.Replace(s)
}
