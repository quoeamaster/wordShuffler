package wordShuffler

type GramSequencer struct {
    // the sequence involved for "word" formation and matching later on
    Sequence string
    // the rule implementation on formulation of "words" based on the given sequence
    shuffleRule ShuffleRule
    // the rule implementation on matching words based on a "source"
    matcherRule MatcherRule

    // the valid sequences / "words" AFTER matching
    validSequences []string
}

// method to create an instance of GramSequencer
func NewGramSequencer(sequence string, minSeqSize, maxSeqSize int,
    matcherRule MatcherRule, shuffleRule ShuffleRule) GramSequencer {

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

    return *m
}

// method to create an instance of GramSequencer
func NewGramSequencerSimple(sequence string) GramSequencer {
    return NewGramSequencer(sequence, -1, -1, nil, nil)
}

func (g *GramSequencer) GenerateValidSequences() error {
    return nil
}

func (g *GramSequencer) GetValidSequences() []string {
    return g.validSequences
}

