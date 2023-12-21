package util

const (
	Charset = "abcdefghijklmnopqrstuvwxyz"
)

func initCharHashMap(charset string) (map[rune]int, map[int]rune) {
	charToIndex := make(map[rune]int)
	indexToChar := make(map[int]rune)

	for i, char := range charset {
		charToIndex[char] = i
		indexToChar[i] = char
	}

	return charToIndex, indexToChar
}

func Rot13Encode(plaintext string) string {
	charToIndex, indexToChar := initCharHashMap(Charset)

	encodedMessage := ""

	// fmt.Printf("Charset length: %v\n", len(charset))

	for _, char := range plaintext {
		newCharIndex := charToIndex[char] + 13
		if newCharIndex >= len(Charset) {
			newCharIndex = newCharIndex - len(Charset)
		}
		// fmt.Printf("char: %c, oldIndex: %v,  newIndex: %v\n", char, charToIndex[char], newCharIndex)
		encodedChar := indexToChar[newCharIndex]
		encodedMessage += string(encodedChar)
	}

	return encodedMessage
}
