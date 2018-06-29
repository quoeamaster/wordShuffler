package DictionaryLookup

import (
    "github.com/DATA-DOG/godog"
    "wordShuffler"
    "fmt"
    "strings"
)

var instance wordShuffler.DictionaryLookup
var word string
var results []wordShuffler.DictionaryLookupResult

// **   scenario 1  **

// step 1
func targetWord(targetWord string) error {
    instance = *wordShuffler.NewDictionaryLookup(5)
    word = targetWord
    return nil
}

// step 2
func doLookup() error {
    tempResults, err := instance.Lookup(word, nil)

    if err != nil {
        return err
    }
    /*
    if tempResults == nil || len(tempResults) == 0 {
        return fmt.Errorf("should have at least 1 explanation~")
    }
    */
    results = tempResults

    return nil
}

// step 3
func validateExplanations(wordsInString string) error {
    expectedWordList := strings.Split(wordsInString, ",")
    tempResults := results
    textFound := false

    // cross check (at least 1 match is a success~)
    for _, expectedWord := range expectedWordList {
        for _, result := range tempResults {
            lowerCaseText := strings.ToLower(result.Text)
            // found or not?
            if strings.Index(lowerCaseText, strings.ToLower(expectedWord)) >= 0 {
                textFound = true
                break
            }
        }
        if textFound == true {
            break
        }
    }

    if textFound == true {
        return nil
    } else {
        return fmt.Errorf("there is no match against the words with the explanations~ expected [%v] but got instead [%v]", expectedWordList, tempResults)
    }
}

func explanationMightMatch(wordList string) error {
    // could be empty as slang or contemporary words are not always available
    if results == nil || len(results) == 0 {
        return nil
    } else {
        return validateExplanations(wordList)
    }
}

// **   scenario 2  **

// step 3
func validateNoExplanation() error {
    // check if the results variable is nil or empty
    if results == nil || len(results) == 0 {
        return nil
    }
    return fmt.Errorf("should not have any explanations at all~~~ [%v]", results)
}

func FeatureContext(s *godog.Suite) {
    s.Step(`^a word "([^"]*)"$`, targetWord)
    s.Step(`^calling the dictionary api\(s\), the corresponding explanation is retrieved$`, doLookup)

    s.Step(`^the explanation MIGHT contain words like "([^"]*)"$`, explanationMightMatch)

    s.Step(`^the explanation should contain words like "([^"]*)"$`, validateExplanations)

    s.Step(`^no explanation should be available$`, validateNoExplanation)
}
