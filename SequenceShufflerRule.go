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
func (s *SequenceShufflerRule) Shuffle(sequence string, optionalArgs... map[string]interface{}) ([]string, error) {
    oldCharArray := []rune(sequence)
    return recursivePermute(oldCharArray[1:], []string{ string(oldCharArray[0]) }), nil
}

// method referenced from http://www.golangprograms.com/golang-program-to-print-all-permutations-of-a-given-string.html

// joining the given runes
func join(ins []rune, c rune) (result []string) {
    for i := 0; i <= len(ins); i++ {
        result = append(result, string(ins[:i])+string(c)+string(ins[i:]))
    }
    return
}

// perform permutation based on the given charArray
func recursivePermute(word []rune, permutationResult []string) []string {
    if len(word) == 0 {
        return permutationResult
    } else {
        result := make([]string, 0)
        for _, e := range permutationResult {
            result = append(result, join([]rune(e), word[0])...)
        }
        return recursivePermute(word[1:], result)
    }
}


