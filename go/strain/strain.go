package strain

type Ints []int
type Lists [][]int
type Strings []string

// Keep filters a collection of ints
func (i Ints) Keep(strain func(int) bool) Ints {
	var output Ints
	for _, v := range i {
		if strain(v) {
			output = append(output, v)
		}
	}
	return output
}

// Discard filters out the keeps of ints
func (i Ints) Discard(strain func(int) bool) Ints {
	var output Ints
	for _, v := range i {
		if !strain(v) {
			output = append(output, v)
		}
	}
	return output
}

// Keep filters a collection of lists
func (l Lists) Keep(strain func([]int) bool) Lists {
	var output Lists
	for _, v := range l {
		if strain(v) {
			output = append(output, v)
		}
	}
	return output
}
func (s Strings) Keep(strain func(string) bool) Strings {
	var output Strings
	for _, v := range s {
		if strain(v) {
			output = append(output, v)
		}
	}
	return output
}
