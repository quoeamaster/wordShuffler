package wordShuffler

import (
    "github.com/golang/groupcache/lru"
    "strings"
    "fmt"
)

const DefaultDictionaryCacheSize = 20

type DictionaryLookup struct {
    currentWordForLookup string
    // sources of lookup
    lookupSources []DictionaryLookupEngine
    // cache for lookup-ed entry(s)
    lookupCache *lru.Cache
}

type DictionaryLookupResult struct {
    Text string
    Language string
}

// return a new instance of the lookup result
func NewDictionaryLookupResult(text, lang string) DictionaryLookupResult {
    m := new(DictionaryLookupResult)

    m.Text = text
    m.Language = lang

    return *m
}

// create a new instance of DictionaryLookup with an optional cache size;
// if the cache size is not given, the default size of "20" is used
func NewDictionaryLookup(cacheSize... int) *DictionaryLookup {
    m := new(DictionaryLookup)
    finalCacheSize := DefaultDictionaryCacheSize

    if cacheSize != nil && len(cacheSize) > 0 {
        finalCacheSize = cacheSize[0]
    }
    m.lookupCache = lru.New(finalCacheSize)

    // TODO: populate the lookupCache contents
    m.lookupSources = make([]DictionaryLookupEngine, 0)
    m.lookupSources = append(m.lookupSources, NewGlosbeEngine())
    m.lookupSources = append(m.lookupSources, NewPearsonEngine())

    return m
}

func (d *DictionaryLookup) Lookup(word string, optionalParams map[string]interface{}) ([]DictionaryLookupResult, error) {
    for _, engine := range d.lookupSources {
        results, err := engine.Lookup(word, optionalParams)
        if err != nil {
            return nil, err
        }
        // TODO: need to check if results is empty??
        return results, nil
    }
    return nil, fmt.Errorf("unknown situation, probably lookup engines are missing")
}

// handy method to check if the given word is valid (not just space)
func IsWordValid(word string) bool {
    if len(strings.TrimSpace(word)) == 0 {
        return false
    }
    return true
}

