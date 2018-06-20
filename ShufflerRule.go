package wordShuffler

import (
    "fmt"
    "strconv"
)

// interface encapsulating shuffling rules
type ShuffleRule interface {
    // shuffle the given / old text based on its unique rules
    Shuffle(oldText string) (string, error)
}

// rule based on the Cambridge research
type CambridgeRule struct {}

func (r *CambridgeRule) Shuffle(oldText string) (string, error) {
    // TODO: investigate stringBuffer
    newText := ""

    runes := []rune(oldText)
    for idx, char := range runes {
        // to convert rune to character/string => strconv.QuoteRune(rune_val)
        // fmt.Println(fmt.Sprintf("runes at idx (%v) => %v -- actual char : %v", idx, char, strconv.QuoteRune(char)))
        
    }
    /*
    for idx := 0; idx < len(oldText); idx++ {
        char := oldText[idx: idx+1]
        shouldIgnore, err := isCharacterIgnorable(char)
        if err != nil {
            return "", err
        }
        fmt.Println(idx, " => ", char, ", can ignore? ", shouldIgnore)
        // TODO: build the newText
    } // end -- for (character extraction of oldText)
    */
    return newText, nil
}


func isCharacterIgnorable(char string) (bool, error) {
    isIgnorable := false
    // is it 1 character wide?
    if len(char) != 1 {
        return false, fmt.Errorf("char MUST be of length of 1")
    }



    /*
    switch char {
    case "\"", "'", ",", "!", ".", " ", "\n", "\r", "“", "”":
        isIgnorable = true
    default:
        isIgnorable = false
    }
    */
    return isIgnorable, nil
}