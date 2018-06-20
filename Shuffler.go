package wordShuffler

import (
    "strings"
    "fmt"
)

type Shuffler struct {
    // location of the file containing the text (could be empty value "")
    Location string
    // Shuffle Rule
    Rule ShuffleRule

    // old / original words
    oldGrams []string
    // new / shuffled words
    newGrams []string

    // original text
    oldText string
    // shuffled text
    newText string
}

func NewShuffler(location string, shuffleRule... ShuffleRule) Shuffler {
    m := new(Shuffler)

    m.Location = location
    if shuffleRule != nil && len(shuffleRule) > 0 {
        m.Rule = shuffleRule[0]
    } else {
        // default rule is Cambridge Rule
        m.Rule = &CambridgeRule{}
    }

    return *m
}

// shuffle the text. If no text is provided, will try to extract the text
// from the file location WHICH might also be empty ""; if that is the case,
// an error would be thrown
func (s *Shuffler) ShuffleText(oldText... string) (string, error) {
    if oldText != nil && len(oldText) > 0 {
        s.oldText = oldText[0]
    } else {
        if strings.Compare(s.Location, "") == 0 {
            return "", fmt.Errorf("there is NO text for shuffling! Either provide it to the function or through a file (check the Location attribute)")
        } else {
            content, err := ReadFileContent(s.Location)
            if err != nil {
                return "", err
            }
            s.oldText = content
        }
    }   // end -- if (oldText provided?)

    // grams breaking for the oldText
    grams, err := gramsBreaking(s.oldText)
    if err != nil {
        return "", err
    }
    s.oldGrams = grams
    s.newGrams = make([]string, 1)

    // run shuffle by the rule object
    for _, oldGram := range s.oldGrams {
        gram, err := s.Rule.Shuffle(oldGram)
        if err != nil {
            return "", err
        }
        s.newGrams = append(s.newGrams, gram)
    }

    // form the newText
    s.newText = combineGramsToText(s.newGrams)

    return s.newText, nil
}

func gramsBreaking(text string) ([]string, error) {
    grams := strings.Split(text, " ")
    return grams, nil
}

func combineGramsToText(grams []string) string {
    return strings.Join(grams, " ")
}