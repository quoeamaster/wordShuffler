package wordShuffler

// interface encapsulating shuffling rules
type ShuffleRule interface {
    // shuffle the given / old text based on its unique rules
    Shuffle(oldText string) (string, error)
}

type AdvanceSuffleRule interface {
    // shuffle the given sequence and return a list of possible "words"
    Shuffle(sequence string, optionalArgs... map[string]interface{}) ([]string, error)
}

// interface encapsulating matching rules with a "source" (e.g. dictionary)
type MatcherRule interface {
    // method to match the given word against a "source"; could be a dictionary.
    MatchWord(word string) (bool, error)
}

// interface for a dictionary lookup engine / approach
type DictionaryLookupEngine interface {
    // returns an arbitrary string to describe the lookup engine (e.g. webster dictionary api)
    Source() string
    // do the lookup for the given "word", wraps the results into
    // an array of DictionaryLookupResult
    Lookup(word string, optionalParams map[string]interface{}) ([]DictionaryLookupResult, error)
}
