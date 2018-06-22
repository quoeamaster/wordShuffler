package wordShuffler

import (
    "fmt"
    "math/rand"
    "time"
)

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
func (s *SequenceShufflerRule) Shuffle(sequence string, charIdx1, charIdx2 int) ([]string, error) {
    fmt.Println(sequence, " idx1 ", charIdx1, ", idx2 ", charIdx2)
    // key => string seq of the runes; value => true / false (default is false)
    runeSeqMap := make(map[string]bool)
    oldCharArray := []rune(s.sequence)

    maxSeqSize := getMaxChoicesSize(len(oldCharArray) - 1)
    rGenerator := rand.New(rand.NewSource(time.Now().UnixNano()))
    rGenerator.Seed(time.Now().UnixNano())
    // fmt.Println("bb > max seq =", maxSeqSize, " oldcharArray length >", len(oldCharArray))

    // get all the possible word combos (maxSeqSize)
    for i := 0; i < maxSeqSize; i++ {
        newCharArray := make([]rune, len(oldCharArray))
        // set the 2 chosen chars values
        newCharArray[charIdx1] = oldCharArray[charIdx1]

        for true {
            // records down which rune has been used...
            seqMapEntry := make(map[int]bool, len(oldCharArray))
            seqMapEntry[charIdx1] = true

            for j := range oldCharArray {
                currentChar := oldCharArray[j]
                if j == charIdx1 {
                    continue
                }
                randIdx := s.GetValidRandomIdx(j, len(oldCharArray), &seqMapEntry, rGenerator)
                newCharArray[randIdx] = currentChar
                // fmt.Println(randIdx, " =>", string(newCharArray), newCharArray)
            }   // end -- for (per rune of the oldCharArray)
            newWord := string(newCharArray)
            // fmt.Println("b1) "+newWord)
            if runeSeqMap[newWord] == true {
                continue
            }
            runeSeqMap[newWord] = true
            break
        }
    }   // end -- for (max seq for a particular word length)

    // translate map back to []string
    wordArray := make([]string, 0)
    for key := range runeSeqMap {
        wordArray = append(wordArray, key)
    }
fmt.Println("cc)", wordArray)
    return wordArray, nil
}

// get the max size for the sequence...
func getMaxChoicesSize(gramLength int) int {
    if gramLength == 1 {
        return 1
    }
    return gramLength * getMaxChoicesSize(gramLength - 1)
}

func (s *SequenceShufflerRule) GetValidRandomIdx(
    targetIdx, oldCharArraySize int, seqMap *map[int]bool, rGenerator *rand.Rand) int {

    // randomly set the other runes
    randIdx := rGenerator.Intn(oldCharArraySize)
    for true {
        // non occupied entry
        if (*seqMap)[randIdx] == false {
            (*seqMap)[randIdx] = true
            break
        } else {
            randIdx = rGenerator.Intn(oldCharArraySize)
        }
    }
    return randIdx
}

func (s *SequenceShufflerRule) handleSequenceWithLengthOf2() error {

    return nil
}

/*
// method to match the given word against a "source"; could be a dictionary.
func (s *SequenceShufflerRule) MatchWord(word string) (bool, error) {

    return false, nil
}
*/
