package protein

import (
	"errors"
)

var (
	ErrStop        error = errors.New("STOP")
	ErrInvalidBase error = errors.New("invalid codon, terminate process")
)

//FromCodon returns the Protein from a specific Codon
func FromCodon(codon string) (string, error) {
	switch codon {
	case "AUG":
		return "Methionine", nil
	case "UUU", "UUC":
		return "Phenylalanine", nil
	case "UUA", "UUG":
		return "Leucine", nil
	case "UCG", "UCU", "UCC", "UCA":
		return "Serine", nil
	case "UAU", "UAC":
		return "Tyrosine", nil
	case "UGU", "UGC":
		return "Cysteine", nil
	case "UGG":
		return "Tryptophan", nil
	case "UAA", "UAG", "UGA":
		return "", ErrStop
	default:
		return "", ErrInvalidBase
	}

}

// FromRNA returns a list of string with all the proteins
func FromRNA(rna string) ([]string, error) {
	var rnaList []string
	for i := 0; i < len(rna); i += 3 {
		protein, err := FromCodon(rna[i : i+3])
		if err == ErrStop {
			break
		}
		if err != nil {
			return rnaList, err
		}
		rnaList = append(rnaList, protein)
	}
	return rnaList, nil
}
