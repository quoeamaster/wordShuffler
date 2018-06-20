package wordShuffler

import "io/ioutil"

// read the contents of a given file
func ReadFileContent(file string) (string, error) {
    byteArr, err := ioutil.ReadFile(file)
    if err != nil {
        return "", err
    }
    return string(byteArr), nil
}
