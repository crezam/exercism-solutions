/*
Package greeting solves the exercism Go: Two Fer
*/
package twofer

import "fmt"

// ShareWith returns a Two Fer game string
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}
	return fmt.Sprintf("One for %v, one for me.", name)
}
