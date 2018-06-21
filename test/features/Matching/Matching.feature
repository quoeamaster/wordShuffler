Feature: Word Matching (with dictionary reference)
    Provide a sequence of characters; would be able to shuffle the sequences and
    match with a valid word within a dictionary.

    For example, "edmo" => "mode"; steps involved:
    a) shuffle the sequences to form a combination of "words"
        e.g. emod, emdo, edom, edmo, eodm, eomd ... mode ... moed etc
    b) among the words given, cross check with the dictionary entries and
        picked out valid words only (e.g. mode or demo)
    c) to be able to form valid words < length of the sequence => e.g.
        "xfo" => "fox" BUT also "ox"

    Assumptions for the feature test:
    - the longer the length of the sequence, the more computation power is needed
        for the shuffling and dictionary-matching; hence would start out with a
        sequence below 5 (so sequences starting from 2 till 5 is valid)
    - the minimum sequence length should be settable
        (controlling the outcomes to be at least a certain sequence length)
    - the maximum sequence length should be settable too
        (avoid computation overheads)

    Major use cases:
    - parse the input sequence and start digging valid "words"

    Scenario: 1) matching a valid sequence
        (valid here means the sequence members could form at least 1 valid word)

        Given a sequence "odme"
        When the analysis has completed
        Then a list of matched valid words of size "2" should be retrieved
        And 1 of the matched words should be "dome"
        And 1 of the matched words should be "demo"


