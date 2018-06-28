package wordShuffler

import (
    "github.com/Jeffail/gabs"
    "fmt"
)

func GetArrayFromJsonByPath(data []byte, path string) ([]*gabs.Container, error) {
    if data == nil || len(data) == 0 {
        return nil, fmt.Errorf("the given data ([]byte) is empty")
    }

    json, err := gabs.ParseJSON(data)
    if err != nil {
        return nil, err
    }

    list, err := json.Search(path).Children()
    if err != nil {
        return nil, err
    }
    return list, nil
}
