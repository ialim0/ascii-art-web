package link

import (
	"strings"
)

// PrintAscii lit une chaîne de caractères en entrée et affiche l'art ASCII correspondant.
func PrintAscii(input string, tab [][]string) string {
	count := 0
	var final string
	// Parcourir chaque ligne de la chaîne de caractères
	for i, mots := range strings.Split(input, "\\n") {
		linenumber := 0
		// Si la ligne n'est pas vide, afficher chaque caractère de la ligne sous forme d'art ASCII
		if len(mots) != 0 {
			for linenumber < 8 {
				for _, bit := range []byte(mots) {
					if bit >= 31 && bit <= 127 {
						// Afficher l'art ASCII correspondant à chaque caractère
						final = final + (tab[bit-32][linenumber])
					} else {
						return string(bit)
					}
				}
				linenumber++
				final = final + "\n"
			}
			count++
			// Si la ligne est vide mais qu'il y a eu au moins une ligne non vide avant, ajouter une ligne vide
		} else if (i != len(strings.Split(input, "\\n"))-1) || count != 0 {
			final = final + "\n"
		}
	}
	return final
}
