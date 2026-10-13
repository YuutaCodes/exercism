
package twofer

// returns text with name if name is not empty
func ShareWith(name string) string {
	text := ""
    if name != "" {
        text = "One for " + name + ", one for me."
    } else {
        text = "One for you, one for me."
    }
	return text
}
