package DictionaryLookup

import (
    "github.com/DATA-DOG/godog"
    "wordShuffler"
    "fmt"
)

var instance wordShuffler.DictionaryLookup
var word string
var results []wordShuffler.DictionaryLookupResult

func targetWord(targetWord string) error {
    instance = *wordShuffler.NewDictionaryLookup(5)
    word = targetWord
    return nil
}

func doLookup() error {
    tempResults, err := instance.Lookup(word, nil)

    if err != nil {
        return err
    }
    if tempResults == nil || len(tempResults) == 0 {
        return fmt.Errorf("should have at least 1 explanation~")
    }
    results = tempResults

    return nil
}

func validateExplanations(wordsInString string) error {
    return godog.ErrPending
}

func FeatureContext(s *godog.Suite) {
    s.Step(`^a word "([^"]*)"$`, targetWord)
    s.Step(`^calling the dictionary api\(s\), the corresponding explanation is retrieved$`, doLookup)
    s.Step(`^the explanation should contain words like "([^"]*)"$`, validateExplanations)
}
