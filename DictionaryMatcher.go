package wordShuffler

// default dictionary path
const DefaultDictionaryLocation = "words_alpha.txt"

// Matcher implementation based on the Matcher interface
// involves efficient indexing of the chunk of dictionary entries
// (e.g. lazy loading of dictionary sections based word prefix match)
type DictionaryMatcher struct {
    dictionaryLocation string
}

// create a new instance of DictionaryMatcher
func NewDictionaryMatcher(location... string) DictionaryMatcher {
    m := new(DictionaryMatcher)

    if location != nil && len(location) > 0 {
        m.dictionaryLocation = location[0]
    } else {
        m.dictionaryLocation = DefaultDictionaryLocation
    }

    return *m
}

func (d *DictionaryMatcher) MatchWord(word string) (bool, error) {
    // TODO: implement it...
    return false, nil
}
