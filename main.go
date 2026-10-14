package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

// Fonction pour mettre la première lettre en majuscule (cap)
func Capitalize(s string) string {
	if s == "" {
		return ""
	}
	runes := []rune(s)
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)
}

// Fonction pour vérifier si un mot commence par une voyelle ou un h
func startsWithVowelOrH(s string) bool {
	if s == "" {
		return false
	}
	firstChar := unicode.ToLower(rune(s[0]))
	return firstChar == 'a' || firstChar == 'e' || firstChar == 'i' || firstChar == 'o' || firstChar == 'u' || firstChar == 'y' || firstChar == 'h'
}

// Fonction pour nettoyer et corriger la ponctuation
func fixPunctuation(text string) string {
	// 1. Enlever l'espace avant les signes de ponctuation courants (, . ! ? : ;)
	reSpaceBefore := regexp.MustCompile(`\s+([,.;!?:]+)`)
	text = reSpaceBefore.ReplaceAllString(text, "$1")

	// 2. S'assurer qu'il y a un espace APRÈS un signe de ponctuation (sauf si c'est la fin du texte ou un guillemet)
	reSpaceAfter := regexp.MustCompile(`([,.;!?:]+)([^\s'”)\]])`)
	text = reSpaceAfter.ReplaceAllString(text, "$1 $2")

	// 3. Gérer les groupes de points de suspension ou ponctuation multiple (ex: " ... " ou " !? ")
	reDots := regexp.MustCompile(`\s*([.]{3,}|[!]{2,}|[?]{2,}|\!\?)\s*`)
	text = reDots.ReplaceAllString(text, " $1 ")
	// Nettoyer les espaces multiples générés
	reMultiSpaces := regexp.MustCompile(`\s{2,}`)
	text = reMultiSpaces.ReplaceAllString(text, " ")

	// 4. Gérer les guillemets simples ' : coller à l'intérieur
	reQuotes := regexp.MustCompile(`'\s+([^']+?)\s+'`)
	text = reQuotes.ReplaceAllString(text, "'$1'")

	return text
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run main.go <input_file> <output_file>")
		return
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]

	// 1. Lire le fichier
	content, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Printf("Erreur lors de la lecture : %v\n", err)
		return
	}

	text := string(content)

	// 2. Découper le texte en mots
	words := strings.Fields(text)

	// 3. Parcourir les mots pour appliquer les règles
	for i := 0; i < len(words); i++ {
		// Gestion de "a" / "an"
		if (words[i] == "a" || words[i] == "A") && i+1 < len(words) {
			if startsWithVowelOrH(words[i+1]) {
				if words[i] == "A" {
					words[i] = "An"
				} else {
					words[i] = "an"
				}
			}
		}

		// Gestion des balises avec un nombre : (up, <num>), (low, <num>), (cap, <num>)
		if (words[i] == "(up," || words[i] == "(low," || words[i] == "(cap,") && i+1 < len(words) {
			numStr := strings.TrimSuffix(words[i+1], ")")
			count, err := strconv.Atoi(numStr)

			if err == nil {
				action := words[i]
				start := i - count
				if start < 0 {
					start = 0
				}

				for j := start; j < i; j++ {
					if action == "(up," {
						words[j] = strings.ToUpper(words[j])
					} else if action == "(low," {
						words[j] = strings.ToLower(words[j])
					} else if action == "(cap," {
						words[j] = Capitalize(words[j])
					}
				}

				words = append(words[:i], words[i+2:]...)
				i--
				continue
			}
		}

		// Gestion des balises simples sans nombre
		if i > 0 {
			if words[i] == "(cap)" {
				words[i-1] = Capitalize(words[i-1])
				words = append(words[:i], words[i+1:]...)
				i--
			} else if words[i] == "(up)" {
				words[i-1] = strings.ToUpper(words[i-1])
				words = append(words[:i], words[i+1:]...)
				i--
			} else if words[i] == "(low)" {
				words[i-1] = strings.ToLower(words[i-1])
				words = append(words[:i], words[i+1:]...)
				i--
			} else if words[i] == "(hex)" {
				decimal, err := strconv.ParseInt(words[i-1], 16, 64)
				if err == nil {
					words[i-1] = strconv.FormatInt(decimal, 10)
				}
				words = append(words[:i], words[i+1:]...)
				i--
			} else if words[i] == "(bin)" {
				decimal, err := strconv.ParseInt(words[i-1], 2, 64)
				if err == nil {
					words[i-1] = strconv.FormatInt(decimal, 10)
				}
				words = append(words[:i], words[i+1:]...)
				i--
			}
		}
	}

	// Recoller les mots
	newText := strings.Join(words, " ")

	// 5. Appliquer la correction de la ponctuation finale
	newText = fixPunctuation(newText)

	// 4. Écrire le résultat dans le fichier de sortie
	err = os.WriteFile(outputFile, []byte(newText), 0644)
	if err != nil {
		fmt.Printf("Erreur lors de l'écriture : %v\n", err)
		return
	}

	fmt.Println("Projet complet exécuté avec succès ! Toutes les règles ont été appliquées.")
}
