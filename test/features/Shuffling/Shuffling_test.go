package Shuffling

import (
    "github.com/DATA-DOG/godog"
    "wordShuffler"
    "strings"
    "fmt"
    "strconv"
)

// instance of shuffler
var shufflerInstance wordShuffler.Shuffler

// step 1
func textLocation(location string) error {
    // using Cambridge Rule
    shufflerInstance = wordShuffler.NewShuffler(location)
    return nil
}

// step 2
func runAnalysis() error {
    _, err := shufflerInstance.ShuffleText()
    if err != nil {
        return err
    }
    return nil
}

// step 3
func getAnalyzedText() error {
    text := shufflerInstance.GetShuffleText()
    if len(strings.TrimSpace(text)) == 0 {
        return fmt.Errorf("something is wrong with the shuffled text, should NOT be zero length")
    }
    return nil
}

// step 4
func characterCountCheck(countInStr string) error {
    iCount, err := strconv.ParseInt(countInStr, 10, 32)
    if err != nil {
        return err
    }
    text := shufflerInstance.GetShuffleText()
    if len(text) != int(iCount) {
        return fmt.Errorf("the count check failed, expected [%v] BUT gained [%v]", iCount, len(text))
    }
    return nil
}

// step 5
func wordMatchTest(idxInStr, word string) error {
    idx, err := strconv.ParseInt(idxInStr, 10, 32)
    if err != nil {
        return err
    }
    grams := shufflerInstance.GetShuffleGrams()
    if strings.Compare(grams[idx], word) == 0 {
        return fmt.Errorf("the word at index %v should NOT match !! [%v] vs [%v]", idx, grams[idx], word)
    }
    return nil
}

func FeatureContext(s *godog.Suite) {
    s.Step(`^a text \/ passage extracted from a file "([^"]*)"$`, textLocation)
    s.Step(`^the analysis has completed$`, runAnalysis)
    s.Step(`^the shuffled text \/ passage could be retrieved$`, getAnalyzedText)
    s.Step(`^the character count is still "([^"]*)" \(including the punctuation marks\)$`, characterCountCheck)
    s.Step(`^the word at index "([^"]*)" doesn\'t equals to "([^"]*)" anymore$`, wordMatchTest)
}
