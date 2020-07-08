// Package strand contains functions to transcript your DNA
package strand

//ToRNA does a transcription of your dna in RNA
func ToRNA(dna string) string {
	transcription := map[string]string{
		"G": "C",
		"C": "G",
		"T": "A",
		"A": "U",
	}
	var output string
	for i := 0; i < len(dna); i++ {
		// Check if the character exists in the transcription map
		if _, exists := transcription[string(dna[i])]; exists {
			output += string(transcription[string(dna[i])])
		} else {
			output += string(dna[i])
		}
	}
	return output
}
