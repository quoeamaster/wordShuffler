package wordShuffler

import (
    "fmt"
    "github.com/Jeffail/gabs"
    "strings"
)

type DLPearson struct {
    url string
    targetLang string
}

const LookupOperationType = "lookup_operation_type"
const defaultLookupType = "headword"
var pearsonDicts = [...]string { "ldoce5", "lasde", "wordwise" }


func NewPearsonEngine() *DLPearson {
    m := new(DLPearson)

    // http://api.pearson.com/v2/dictionaries/entries?headword=angry
    m.url = "http://api.pearson.com/v2/dictionaries/entries?%v=%v"

    return m
}

func (d *DLPearson) Source() string {
    return "Pearson api"
}

func (d *DLPearson) Lookup(word string, optionalParams map[string]interface{}) ([]DictionaryLookupResult, error) {
    finalUrl := d.url

    if IsWordValid(word) == false {
        return nil, fmt.Errorf("the given word is INVALID, probably an empty string [%v]", word)
    }
    if optionalParams != nil {
        if val := optionalParams[LookupOperationType]; val != nil {
            finalUrl = fmt.Sprintf(finalUrl, val.(string), word)
        } else {
            finalUrl = fmt.Sprintf(finalUrl, defaultLookupType, word)
        }
    } else {
        finalUrl = fmt.Sprintf(finalUrl, defaultLookupType, word)
    }
    // make http call
    byteArr, err := RunHttpRequest(finalUrl)
    if err != nil {
        return nil, err
    }
    explanationList, err := GetArrayFromJsonByPath(byteArr, "results")
    if err != nil {
// TODO: check for normal no result(s) case
        return nil, err
    }
    finalResults, err := d.prepareDictionaryLookupResults(explanationList)
    if err != nil {
        return nil, err
    }

    return finalResults, nil
}

func (d *DLPearson) prepareDictionaryLookupResults(list []*gabs.Container) ([]DictionaryLookupResult, error) {
    results := make([]DictionaryLookupResult, 0)

    for _, explanation := range list {
        // check if the dictionary entry is correct or not (ldoce5, lasde, wordwise)
        dictMatched := false
        dictList, err := explanation.Search("datasets").Children()
        if err != nil {
            return nil, err
        }
        for _, dict := range dictList {
            switch dict.Data().(type) {
            case string:
                if d.isDictNameMatched(dict.Data().(string)) {
                    dictMatched = true
                }
            default:
                return nil, fmt.Errorf("non support type met for result (dataset)")
            }
            if dictMatched {
                break
            }
        }

        if dictMatched {
            // search for the sub-clause "senses"
            definition := explanation.Search("senses").Search("definition")
            if definition != nil {
                defList, err := definition.Children()
                if err != nil {
                    return nil, err
                }
                if len(defList) > 0 {
                    firstDef := defList[0].Data()
                    switch firstDef.(type) {
                    case []interface{}:
                        vList := firstDef.([]interface{})
                        fmt.Println(vList)
                        results = append(results, NewDictionaryLookupResult(vList[0].(string), "eng"))
                    case string:
                        results = append(results, NewDictionaryLookupResult(firstDef.(string), "eng"))
                    default:
                        return nil, fmt.Errorf("non support type met for definition (neither string or []interface{})")
                    }
                }
            }
        }   // end -- if (dict is a match)
    }
    return results, nil
}

func (d *DLPearson) isDictNameMatched(dictName string) bool {
    finalDictName := strings.ToLower(dictName)
    for _, eDict := range pearsonDicts {
        if strings.Compare(eDict, finalDictName) == 0 {
            return true
        }
    }
    return false
}
