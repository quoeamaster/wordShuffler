package wordShuffler

import (
    "net/http"
    "time"
    "io/ioutil"
)

// run an http request based on the given url, an optional map object could
// be supplied for optional parameters
func RunHttpRequest(url string, optional... map[string]interface{}) ([]byte, error) {
    // set timeout to 2 seconds
    clientPtr := &http.Client{
        Timeout: time.Second * 2,
    }
    response, err := clientPtr.Get(url)
    if err != nil {
        return nil, err
    }
    byteArr, err := ioutil.ReadAll(response.Body)
    if err != nil {
        return nil, err
    }

    return byteArr, nil
}


