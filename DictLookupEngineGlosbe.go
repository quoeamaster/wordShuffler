package wordShuffler

import (
    "fmt"
    "strings"
    "github.com/Jeffail/gabs"
)

const KeyLangFrom = "key_lang_from"
const KeyLangTo = "key_lang_to"

const defaultLang = "eng"

type DLGlosbe struct {
    url string
}

func NewGlosbeEngine() *DLGlosbe {
    m := new(DLGlosbe)

    // https://glosbe.com/gapi/translate?from=eng&dest=eng&format=json&phrase=blowjob&pretty=true
    m.url = "https://glosbe.com/gapi/translate?from=%v&dest=%v&format=json&phrase=%v&pretty=false"

    return m
}


func (d *DLGlosbe) Source() string {
    return "glosbe api"
}

func (d *DLGlosbe) Lookup(word string, optionalParams map[string]interface{}) ([]DictionaryLookupResult, error) {
    finalUrl := ""

    if IsWordValid(word) == false {
        return nil, fmt.Errorf("the given word is INVALID, probably an empty string [%v]", word)
    }
    // apply parameters to the api
    if optionalParams != nil {
        // only accept map object (might panic though)
        langFrom := "eng"
        langTo := "eng"

        if v:=optionalParams[KeyLangFrom]; v != nil {
            lang := v.(string)
            if len(strings.TrimSpace(lang)) > 0 {
                langFrom = lang
            }
        }
        if v:=optionalParams[KeyLangTo]; v != nil {
            lang := v.(string)
            if len(strings.TrimSpace(lang)) > 0 {
                langTo = lang
            }
        }
        finalUrl = fmt.Sprintf(d.url, langFrom, langTo, word)

    } else {
        finalUrl = fmt.Sprintf(d.url, defaultLang, defaultLang, word)
    }

    // make an http request to get back the results (through the api)
    resultJsonByteArr, err := RunHttpRequest(finalUrl)
    if err != nil {
        return nil, err
    }
    explanationList, err := GetArrayFromJsonByPath(resultJsonByteArr, "tuc")
    if err != nil {
        // ** probably because there is no explanation (invalid word)
        if strings.Compare(err.Error(), "not an object or array") == 0 {
            return nil, nil
        }
        return nil, err
    }
    // prepare the final list
    finalList, err := d.prepareDictionaryLookupResult(explanationList)
    if err != nil {
        return nil, err
    }

    // try catch block
    defer func() {
        if r := recover(); r != nil {
            fmt.Printf("recovered => %v\n", r)
        }
    }()

    return finalList, nil
}

// prepare / convert the given explanation list into []DictionaryLookupResult
func (d *DLGlosbe) prepareDictionaryLookupResult(list []*gabs.Container) ([]DictionaryLookupResult, error) {
    if list == nil {
        return nil, fmt.Errorf("no VALID result returned")
    }

    finalList := make([]DictionaryLookupResult, 0)
    for _, result := range list {
        meanings := result.Search("meanings")
        if meanings == nil {
            // some might not contain any explanations at all
            continue
        }
        meaningsList, err := meanings.Children()
        if err != nil {
            return nil, fmt.Errorf("could not get the meanings")
        }
        for _, meaning := range meaningsList {
            switch meaning.Data().(type) {
            case map[string]interface{}:
                vMap := meaning.Data().(map[string]interface{})
                finalList = append(finalList, NewDictionaryLookupResult(vMap["text"].(string), vMap["language"].(string)))

            default:
                return nil, fmt.Errorf("non support type met [%v]", meaning.Data())
            }
        }
    }   // end -- for (iteration from the given list, not yet reached the "meaning" clause yet)
    return finalList, nil
}
