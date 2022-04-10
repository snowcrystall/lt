package problems

func UniqueMorseRepresentations(words []string) int {
	code := []string{".-","-...","-.-.","-..",".","..-.","--.","....","..",".---","-.-",".-..","--","-.","---",".--.","--.-",".-.","...","-","..-","...-",".--","-..-","-.--","--.."}
	wordmap := map[string]int{}
	for _,word := range words {
		codestring := ""
		for _,v := range  word {
			codestring += code[v-97]
		}
		wordmap[codestring] = 1
	}
	return len(wordmap)
}
