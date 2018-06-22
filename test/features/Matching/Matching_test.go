package Matching

import (
    "github.com/DATA-DOG/godog"
    "wordShuffler"
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

func validWordsSizeCheck(sizeInStr string) error {
    return godog.ErrPending
}

func validWordMatchCheck(_ int, targetWord string) error {
    return godog.ErrPending
}

func FeatureContext(s *godog.Suite) {
    s.Step(`^a sequence "([^"]*)"$`, sequenceProvided)
    s.Step(`^the analysis has completed$`, analyzeSequence)
    s.Step(`^a list of matched valid words of size "([^"]*)" should be retrieved$`, validWordsSizeCheck)
    s.Step(`^(\d+) of the matched words should be "([^"]*)"$`, validWordMatchCheck)
}
