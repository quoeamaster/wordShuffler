package wordShuffler

type SequenceShufflerRule struct {
    // the minimum size of the "words" to be created from the given sequence.
    // Theoretically it should be at least "2"
    MinGramSize int
    // optional size threshold for the sequence generation (remember the
    // longer is the sequence, the more computation is required)
    MaxGramSize int

    // the actual sequence for formulation
    sequence string
}

// create a new instance of SequenceShufflerRule
func NewSequenceShufflerRule(minSize, maxSize int, sequence string) SequenceShufflerRule {
    m := new(SequenceShufflerRule)

    m.sequence = sequence
    // minimum sequence is 2!
    if minSize < 2 {
        m.MinGramSize = 2
    } else {
        m.MinGramSize = minSize
    }
    // maxSize check
    if maxSize < m.MinGramSize {
        m.MaxGramSize = m.MinGramSize
    } else {
        m.MaxGramSize = maxSize
    }
    if m.MaxGramSize > len(sequence) {
        // should be max => the same value as the length of the sequence
        m.MaxGramSize = len(sequence)
    }
    return *m
}

// implementation of the ShuffleRule; however the returned string value
// is not meaningful in here. Instead should invoke "GetValidSequences" method
func (s *SequenceShufflerRule) Shuffle(_ string) (string, error) {



    return "", nil
}

/*
// method to match the given word against a "source"; could be a dictionary.
func (s *SequenceShufflerRule) MatchWord(word string) (bool, error) {

    return false, nil
}
*/
