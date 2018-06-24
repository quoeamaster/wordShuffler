package Matching

import (
    "github.com/DATA-DOG/godog"
    "wordShuffler"
    "strconv"
    "fmt"
    "strings"
)

var instance wordShuffler.GramSequencer

// step 1
func sequenceProvided(sequence string) error {
    instance = wordShuffler.NewGramSequencerSimple(sequence)
    return nil
}

// step 2
func analyzeSequence() error {
    err := instance.GenerateValidSequences()
    if err != nil {
        return err
    }
    return nil
}

// step 3
func validWordsSizeCheck(sizeInStr string) error {
    iSize, err := strconv.ParseInt(sizeInStr, 10, 64)
    if err != nil {
        return err
    }
    validSeq := instance.GetValidSequences()
    fmt.Println("valid sequences >", validSeq)

    if len(validSeq) != int(iSize) {
        return fmt.Errorf("size of the VALID sequences are not equal, expected [%v] BUT got [%v]", iSize, len(validSeq))
    }
    return nil
}

// step 4
func validWordMatchCheck(_ int, targetWord string) error {
    validSeq := instance.GetValidSequences()

    if isWordFound(targetWord, validSeq) == true {
        return nil
    } else {
        return fmt.Errorf("not matched~ expected [%v] BUT not found within sequences [%v]", targetWord, validSeq)
    }
}

func isWordFound(word string, validSeq []string) bool {
    for _, seq := range validSeq {
        if strings.Compare(seq, word) == 0 {
            return true
        }
    }
    return false
}

func FeatureContext(s *godog.Suite) {
    s.Step(`^a sequence "([^"]*)"$`, sequenceProvided)
    s.Step(`^the analysis has completed$`, analyzeSequence)
    s.Step(`^a list of matched valid words of size "([^"]*)" should be retrieved$`, validWordsSizeCheck)
    s.Step(`^(\d+) of the matched words should be "([^"]*)"$`, validWordMatchCheck)
}
