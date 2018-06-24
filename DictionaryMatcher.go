package wordShuffler

import (
    "strings"
    "fmt"
    "os"
    "bufio"
    "io"
)

// default dictionary path
const DefaultDictionaryLocation = "words_alpha.txt"

// default environment variable name => "MATCHER_DICT_LOCATION"
const EnvVariableDictLocation = "MATCHER_DICT_LOCATION"

// Matcher implementation based on the Matcher interface
// involves efficient indexing of the chunk of dictionary entries
// (e.g. lazy loading of dictionary sections based word prefix match)
type DictionaryMatcher struct {
    // location of the dictionary file
    dictionaryLocation string

    // dict entry map (lazy cached and should be evicted later on based on LRU rules)
    dictionaryEntriesByLetter map[string][]string
}

// create a new instance of DictionaryMatcher
func NewDictionaryMatcher(location... string) DictionaryMatcher {
    m := new(DictionaryMatcher)

    if location != nil && len(location) > 0 {
        m.dictionaryLocation = location[0]
    } else {
        m.dictionaryLocation = DefaultDictionaryLocation
    }

    m.dictionaryEntriesByLetter = make(map[string][]string, 0)

    return *m
}

func (d *DictionaryMatcher) MatchWord(word string) (bool, error) {
    dictEntries, err := d.getDictionaryEntriesByPrefix(word)
    if err != nil {
        return false, err
    }

    finalWord := strings.ToLower(strings.TrimSpace(word))
    // matching
    for _, entry := range dictEntries {
        finalEntry := strings.ToLower(strings.TrimSpace(entry))
        if strings.Compare(finalWord, finalEntry) == 0 {
            return true, nil
        }
    }
    return false, nil
}

func (d *DictionaryMatcher) getDictionaryEntriesByPrefix(word string) ([]string, error) {
    if len(strings.TrimSpace(word)) == 0 {
        return nil, fmt.Errorf("invalid PREFIX provided, should be non-empty string [%v]", word)
    }
    // TODO: LRU cache eviction required???
    prefixCharInString := string(([]rune(word))[0])
    dictEntries := d.dictionaryEntriesByLetter[prefixCharInString]

    if dictEntries == nil || len(dictEntries) == 0 {
        // could be the full word or just the prefix character (for performance tuning...)
        dictEntries, err := d.seekFileContentsByPrefix(d.dictionaryLocation, word)
        if err != nil {
            return nil, err
        }
        d.dictionaryEntriesByLetter[prefixCharInString] = dictEntries
    }
    return d.dictionaryEntriesByLetter[prefixCharInString], nil
}



func (d *DictionaryMatcher) seekFileContentsByPrefix(fileLocation string, prefix string) ([]string, error) {
    // setup the slice for matched entry(s) (matched prefix)
    entryArr := make([]string, 0)
    charArray := []rune(prefix)
    prefixCharInByte := byte(charArray[0])
    seekStarted := false

    // get back the correct file location (also check with environment variable)
    finalFileLocation, err := SeekFileLocation(d.dictionaryLocation, EnvVariableDictLocation)
    if err != nil {
        return nil, err
    }

    filePtr, err := os.Open(finalFileLocation)
    if err != nil {
        return nil, err
    }
    defer filePtr.Close()
    reader := bufio.NewReader(filePtr)

    for true {
        byteArr, _, err := reader.ReadLine()
        // fmt.Println(string(byteArr))
        if err == io.EOF {
            break
        } else if err != nil {
            return nil, err
        }

        if byteArr[0] == prefixCharInByte {
            seekStarted = true
            entryArr = append(entryArr, string(byteArr))

        } else if seekStarted == true {
            // all prefix matched entries met
            break
        }
    }
    return entryArr, nil
}


