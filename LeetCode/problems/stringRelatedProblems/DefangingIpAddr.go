/*Given a valid (IPv4) IP address, return a defanged version of that IP address.

A defanged IP address replaces every period "." with "[.]".

*/

package problems

import "strings"

func DefangIPaddr(address string) string {

	var builder strings.Builder // for efficient string concatenation

	for _, char := range address {
		if char == '.' {
			builder.WriteString("[.]")
		} else {
			builder.WriteRune(char)
		}
	}
	return builder.String()

}
