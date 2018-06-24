package wordShuffler

import (
    "fmt"
    "sort"
)

type GramSequencer struct {
    // the sequence involved for "word" formation and matching later on
    Sequence string
    // the rule implementation on formulation of "words" based on the given sequence
    shuffleRule AdvanceSuffleRule
    // the rule implementation on matching words based on a "source"
    matcherRule MatcherRule

    // the valid sequences / "words" AFTER matching; need to convert back to []string
    validSequences map[string]bool
    validSequencesArray []string
}

// method to create an instance of GramSequencer
func NewGramSequencer(sequence string, minSeqSize, maxSeqSize int,
    matcherRule MatcherRule, shuffleRule AdvanceSuffleRule) GramSequencer {

    m := new(GramSequencer)

    m.Sequence = sequence
    // shuffeRule
    if shuffleRule == nil {
        rule := NewSequenceShufflerRule(minSeqSize, maxSeqSize, sequence)
        m.shuffleRule = &rule
    } else {
        m.shuffleRule = shuffleRule
    }
    // matcherRule
    if matcherRule == nil {
        rule := NewDictionaryMatcher()
        m.matcherRule = &rule
    } else {
        m.matcherRule = matcherRule
    }

    // create slice
    m.validSequences = make(map[string]bool, 0)

    return *m
}

// method to create an instance of GramSequencer
func NewGramSequencerSimple(sequence string) GramSequencer {
    return NewGramSequencer(sequence, -1, -1, nil, nil)
}

// method to generate "valid" sequences created from the given sequence.
// Valid or not depends on the implementation of the Matcher
func (g *GramSequencer) GenerateValidSequences() error {
    var newGrams []string

    // special handling for characters of length of "1"
    if len(g.Sequence) > 0 && len(g.Sequence) == 1 {
        newGrams = append(newGrams, g.Sequence)
        g.populateValidSequenceMap(newGrams)

    } else if len(g.Sequence) == 2 {
        newGrams = make([]string, 2)
        newGrams = append(newGrams, g.Sequence)
        newGrams = append(newGrams, fmt.Sprintf("%v%v", g.Sequence[1:], g.Sequence[0:1]))
        g.populateValidSequenceMap(newGrams)

    } else {
        newGrams, err := g.shuffleRule.Shuffle(g.Sequence)
        if err != nil {
            return err
        }
        if len(newGrams) == 0 {
            return fmt.Errorf("length of the words created after the shuffle should be at least 1~ [%v]", g.Sequence)
        }
        err = g.populateValidSequenceMap(newGrams)
        if err != nil {
            return err
        }
    }
    return nil
}

// populate a map from the given words slice / array; de-duplication logic
// applied as well.
func (g *GramSequencer) populateValidSequenceMap(grams []string) error {
    needUpdate := false
    for _, gram := range grams {
        if g.validSequences[gram] == false {
            g.validSequences[gram] = true
            needUpdate = true
        }
    }
    if needUpdate {
        // sorted
        uniqueSeqArr := g.convertValidSequenceMapToArray()
        sort.Strings(uniqueSeqArr)
        // do matching
        // reset the validSequencesArray first (0 length)
        g.validSequencesArray = make([]string, 0)
        if uniqueSeqArr != nil {
            for _, seq := range uniqueSeqArr {
                bMatched, err := g.matcherRule.MatchWord(seq)

                if err != nil {
                    return err
                }
                if bMatched == true {
                    g.validSequencesArray = append(g.validSequencesArray, seq)
                }
            }   // end -- for (per entry within uniqueSeqArr)
        }   // end -- if (uniqueSeqArr VALID)
    }
    return nil
}

// convert a map back to a string slice / array
func (g *GramSequencer) convertValidSequenceMapToArray() []string {
    arr := make([]string, 0)

    for key := range g.validSequences {
        arr = append(arr, key)
    }
    return arr
}

// simple getter
func (g *GramSequencer) GetValidSequences() []string {
    return g.validSequencesArray
}

