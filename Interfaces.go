package wordShuffler

// interface encapsulating shuffling rules
type ShuffleRule interface {
    // shuffle the given / old text based on its unique rules
    Shuffle(oldText string) (string, error)
}

// interface encapsulating matching rules with a "source" (e.g. dictionary)
type MatcherRule interface {
    // method to match the given word against a "source"; could be a dictionary.
    MatchWord(word string) (bool, error)
}