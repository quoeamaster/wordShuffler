package wordShuffler

import "fmt"

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
        // 2 chars must be picked to make the algorithm work
        // pick the 1st char
        /*
        for idx1:=0; idx1 < len(g.Sequence); idx1++ {
            ...
        }   // end -- for (pick the "selected" char)
        */
        newGrams, err := g.shuffleRule.Shuffle(g.Sequence, -1, -1)
        if err != nil {
            return err
        }
        if len(newGrams) == 0 {
            return fmt.Errorf("length of the words created after the shuffle should be at least 1~ [%v]", g.Sequence)
        }
        g.populateValidSequenceMap(newGrams)
    }
    // fmt.Println(newGrams)
    fmt.Println(g.validSequences)



    return nil
}

func (g *GramSequencer) populateValidSequenceMap(grams []string) {
    needUpdate := false
    for _, gram := range grams {
        if g.validSequences[gram] == false {
            g.validSequences[gram] = true
            needUpdate = true
        }
    }
    if needUpdate {
        g.convertValidSequenceMapToArray()
    }
}

func (g *GramSequencer) convertValidSequenceMapToArray() []string {
    arr := make([]string, 0)

    for key := range g.validSequences {
        arr = append(arr, key)
    }
    return arr
}

func (g *GramSequencer) GetValidSequences() []string {
    return g.validSequencesArray
}

