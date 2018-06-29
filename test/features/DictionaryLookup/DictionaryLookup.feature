Feature: Dictionary lookup
    Simply calling api from well known dictionaries to get the meaning of a given word.
    Dictionary(s):
    a) glosbe.com (sort of a collaborated version of dictionary lookups (most importantly... free)
    b) Pearson (have access limit to 1000 request per day? Also deprecated already...)
    c) wordsapi.com (have access limit of 2500 request per day)

    Assumptions for the feature test:
    - the given word is VALID and internet access is available (calling the api)
    - usually there is a pronunciation available for the chosen word; try to seek for libraries to "READ" the word

    Major use cases:
    - find explanations of the words by querying the api (1 of them at least)
    - TODO: read / pronounce the words

    Scenario: 1) find explanation based on a word (slang or contemporary words)
        Given a word "hangry"
        When calling the dictionary api(s), the corresponding explanation is retrieved
        And the explanation MIGHT contain words like "angry,hungry,hunger"

    Scenario: 2) find explanation based on a NON existing word
        Given a word "roda"
        When calling the dictionary api(s), the corresponding explanation is retrieved
        And no explanation should be available

    Scenario: 3) find explanation based on a normal word (non slang)
        Given a word "cock"
        When calling the dictionary api(s), the corresponding explanation is retrieved
        And the explanation should contain words like "penis,chicken,male"

