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

    Scenario: 1) matching a valid sequence (length of PLUS double characters)

        Given a sequence "pap"
        When the analysis has completed
        Then a list of matched valid words of size "3" should be retrieved
        And 1 of the matched words should be "app"

    Scenario: 2) matching a valid sequence
        (valid here means the sequence members could form at least 1 valid word)

        Given a sequence "odme"
        When the analysis has completed
        Then a list of matched valid words of size "4" should be retrieved
        And 1 of the matched words should be "dome"
        And 1 of the matched words should be "demo"

    Scenario: 3) matching a valid sequence (length of 5)

        Given a sequence "cpsir"
        When the analysis has completed
        Then a list of matched valid words of size "3" should be retrieved
        And 1 of the matched words should be "crisp"

    Scenario: 4) matching a valid sequence (length of PLUS double characters)

        Given a sequence "phapy"
        When the analysis has completed
        Then a list of matched valid words of size "1" should be retrieved
        And 1 of the matched words should be "happy"

    Scenario: 5) matching a valid sequence (length of PLUS triple characters)

        Given a sequence "banana"
        When the analysis has completed
        Then a list of matched valid words of size "1" should be retrieved
        And 1 of the matched words should be "banana"

    Scenario: 6) matching a valid sequence (length of PLUS double characters)

        Given a sequence "anspe"
        When the analysis has completed
        Then a list of matched valid words of size "10" should be retrieved
        And 1 of the matched words should be "panse"
        And 1 of the matched words should be "peans"