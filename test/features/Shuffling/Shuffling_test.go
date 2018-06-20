package Shuffling

import (
    "github.com/DATA-DOG/godog"
    "wordShuffler"
    "fmt"
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
    newText, err := shufflerInstance.ShuffleText()
    if err != nil {
        return err
    }
    fmt.Println(newText)

    return nil
}

func getAnalyzedText() error {
    return godog.ErrPending
}

func characterCountCheck(arg1 string) error {
    return godog.ErrPending
}

func wordMatchTest(arg1, arg2 string) error {
    return godog.ErrPending
}

func FeatureContext(s *godog.Suite) {
    s.Step(`^a text \/ passage extracted from a file "([^"]*)"$`, textLocation)
    s.Step(`^the analysis has completed$`, runAnalysis)
    s.Step(`^the shuffled text \/ passage could be retrieved$`, getAnalyzedText)
    s.Step(`^the character count is still "([^"]*)" \(including the punctuation marks\)$`, characterCountCheck)
    s.Step(`^the word at index "([^"]*)" doesn\'t equals to "([^"]*)" anymore$`, wordMatchTest)
}
