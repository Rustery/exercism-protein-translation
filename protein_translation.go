package protein

import (
	"errors"
)

var ErrStop = errors.New("stop codon found")
var ErrInvalidBase = errors.New("codon not found")

func FromRNA(rna string) ([]string, error) {
	proteinsFound := map[string]int{}
	order := 0
	for i := 0; i < len(rna); i += 3 {
		codon := rna[i : i+3]
		c, err := FromCodon(codon)
		if err == ErrStop {
			break
		} else if err == ErrInvalidBase {
			return []string{}, ErrInvalidBase
		}
		proteinsFound[c] = order
		order++
	}
	result := make([]string, order)
	for protein, pos := range proteinsFound {
		result[pos] = protein
	}
	return result, nil
}

func FromCodon(codon string) (string, error) {
	switch codon {
	case "AUG":
		return "Methionine", nil
	case "UUU", "UUC":
		return "Phenylalanine", nil
	case "UUA", "UUG":
		return "Leucine", nil
	case "UCU", "UCC", "UCA", "UCG":
		return "Serine", nil
	case "UAU", "UAC":
		return "Tyrosine", nil
	case "UGU", "UGC":
		return "Cysteine", nil
	case "UGG":
		return "Tryptophan", nil
	case
		"UAA", "UAG", "UGA":
		return "", ErrStop
	default:
		return "", ErrInvalidBase
	}
}
