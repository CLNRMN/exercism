//Package proverb ,this should be a comment for the package, but i don't know what i should write here
package proverb

import "fmt"

const (
	stanza = "For want of a %s the %s was lost."
	outro  = "And all for the want of a %s."
)

// Proverb should have a comment documenting it. *should* ;)
func Proverb(rhyme []string) []string {
	var output []string
	if len(rhyme) > 0 {
		for i := 0; i < len(rhyme)-1; i++ {
			output = append(output, fmt.Sprintf(stanza, rhyme[i], rhyme[i+1]))
		}
		output = append(output, fmt.Sprintf(outro, rhyme[0]))
	}
	return output
}
