package wordShuffler

import (
    "fmt"
    "strings"
)

const KeyLangFrom = "key_lang_from"
const KeyLangTo = "key_lang_to"

const defaultLang = "eng"

type DLGlosbe struct {
    url string
}

func NewEngine() *DLGlosbe {
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
        return nil, err
    }
    fmt.Println(explanationList, len(explanationList))


    // try catch block
    defer func() {
        if r := recover(); r != nil {
            fmt.Printf("recovered => %v\n", r)
        }
    }()

    return nil, fmt.Errorf("unknown situation met")
}
