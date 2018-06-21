package wordShuffler

import (
    "fmt"
    "strings"
    "math/rand"
    "time"
)

// rule based on the Cambridge research
type CambridgeRule struct {}

func (r *CambridgeRule) Shuffle(oldText string) (string, error) {
    // TODO: investigate stringBuffer
    newText := ""

    if len(strings.TrimSpace(oldText)) == 0 {
        return "", nil
    }
    runes := []rune(oldText)
    // to convert rune to character/string => strconv.QuoteRune(rune_val)
    idxFirstAlphanumeric := getFirstAlphanumericChar(runes)
    idxLastAlphanumeric := getLastAlphanumericChar(runes)
    if idxLastAlphanumeric == -1 || idxLastAlphanumeric == -1 {
        return "", fmt.Errorf("the word provided doen't contain any alphanumeric characters [%v]", oldText)
    }
    /*
    fmt.Println("first char @", idxFirstAlphanumeric,
        " last char @", idxLastAlphanumeric,
        " actual word => ", oldText,
        " first char => ", strconv.QuoteRune(runes[idxFirstAlphanumeric]),
        " last char => ", strconv.QuoteRune(runes[idxLastAlphanumeric]))
    */
    newText = generateRandomCharSequence(idxFirstAlphanumeric, idxLastAlphanumeric, oldText)

    return newText, nil
}


// helper method to get back the first index of an alphanumeric character
func getFirstAlphanumericChar(charArray []rune) int {
    for idx, char := range charArray {
        // alphanumeric check
        if (char >= 65 && char <= 90) || (char >= 97 && char <= 122) || (char >= 48 && char <= 57) {
            return idx
        }
    }   // loop through the runes of the charArray
    return -1
}
// helper method to get back the last index of an alphanumeric character
func getLastAlphanumericChar(charArray []rune) int {
    runeLen := len(charArray) - 1
    for idx := runeLen; idx >= 0; idx-- {
        char := charArray[idx]
        if (char >= 65 && char <= 90) || (char >= 97 && char <= 122) || (char >= 48 && char <= 57) {
            return idx
        }
    }
    return -1
}

func generateRandomCharSequence(idxFirst, idxLast int, oldText string) string {
    trimmedOldText := strings.TrimSpace(oldText)
    if len(trimmedOldText) >= 0 && len(trimmedOldText) <= 3 {
        return trimmedOldText
    }
    innerString := oldText[(idxFirst + 1): idxLast]
    charArray := []rune(innerString)
    if len(charArray) == 1 {
        return oldText
    }
    destCharArray := make([]rune, len(charArray))
    // default fill with LF (10)
    // fillRuneArrayWithValue(&destCharArray, '\n')

    // generate randomizer
    rGenerator := rand.New(rand.NewSource(time.Now().UnixNano()))
    rGenerator.Seed(time.Now().UnixNano())

    for idx, charInRune := range charArray {
        // get a random idx for the current rune
        // the new idx can't be the same as the original idx
        // check if the new idx already occupied or not
        newIdx := rGenerator.Intn(len(charArray))
        for true {
            if newIdx == idx {
                if idx == (len(charArray) - 1) && destCharArray[idx] == 0 {
                    break
                }
                newIdx = rGenerator.Intn(len(charArray))
            } else {
                if destCharArray[newIdx] != 0 {
                    // fmt.Println("* need to regenerate idx => ", idx, "x", newIdx, "content =", destCharArray)
                    newIdx = rGenerator.Intn(len(charArray))
                } else {
                    // fmt.Println("idx vs newIdx -", idx, "x", newIdx)
                    break
                }
            }
        }
        destCharArray[newIdx] = charInRune
        //fmt.Println("-> idx vs newIdx -", idx, "x", newIdx, "contents -", destCharArray)
    }
    // newText := fmt.Sprintf("%v-%v-%v", oldText[0:(idxFirst + 1)], string(destCharArray), oldText[idxLast:])
    // fmt.Println("bb) transformed innerstring > ", newText, " ori => ", oldText)
    //fmt.Println(destCharArray)
    newText := fmt.Sprintf("%v%v%v", oldText[0:(idxFirst + 1)], string(destCharArray), oldText[idxLast:])

    return newText
}

