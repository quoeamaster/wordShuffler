Feature: Word Shuffling / Scrambling
    By the theory of a Cambridge research, the human brain can interpret
    words with certain typos and structure.
    For example, "tohuhgt" => "thought"; but to make it work, the text would
    need to fulfill certain conditions:
    a) it must be a passage, also the words would create meaningful content
        (simply you can't provide random text to form this passage)
    b) the words actually is shuffling by this rule =>
        the 1st and last character must be accurate,
        for the rest just shuffle them

    Assumptions for the feature test:
    - a meaningful passage is provided or extracted from some sources (e.g. news reportings)

    Major use cases:
    - parse the input text / passage correctly and apply the shuffling
    - BONUS => read the text / passage by sphinx api (Carnegie Mellon University research)

    Scenario: 1) shuffling of text / passage
        Given a text / passage extracted from a file "readings_01.txt"
        When the analysis has completed
        Then the shuffled text / passage could be retrieved
        And the character count is still "708" (including the punctuation marks)
        And the word at index "2" doesn't equals to "department" anymore

