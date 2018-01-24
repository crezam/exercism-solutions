package raindrops

import "strconv"

// Convert converts an integer to a Raindrop string
func Convert(i int) string {
	output := ""
	if i%3 == 0 {
		output = "Pling"
	}
	if i%5 == 0 {
		output = output + "Plang"
	}
	if i%7 == 0 {
		output = output + "Plong"
	}
	if output == "" {
		return strconv.Itoa(i)
	}
	return output
}
