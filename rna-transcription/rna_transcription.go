package strand

func ToRNA(dna string) string {

	var mappings = map[string]string{"G": "C", "C": "G", "T": "A", "A": "U"}
	var rnaString string
	for k := range dna {
		rnaString += mappings[string(dna[k])]

	}
	return rnaString
}
