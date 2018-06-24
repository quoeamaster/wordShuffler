package wordShuffler

import (
    "io/ioutil"
    "os"
    "strings"
    "fmt"
)

// read the contents of a given file
func ReadFileContent(file string) (string, error) {
    byteArr, err := ioutil.ReadFile(file)
    if err != nil {
        return "", err
    }
    return string(byteArr), nil
}

// method to check if the targeted file exists or not and return the 1st matched path
// which could be the path given or the value under the env-var
func SeekFileLocation(location string, envVariableName string) (string, error) {
    fileInfo, err := os.Stat(location)
    if os.IsNotExist(err) || fileInfo.IsDir() {
        // location is invalid, try env_var
        envVar := os.Environ()
        for _, eVar := range envVar {
            kv := strings.Split(eVar, "=")
            if kv != nil && len(kv) == 2 &&
                strings.Compare(envVariableName, kv[0]) == 0 {

                return kv[1], nil
            }   // end -- if (key-value pair valid with length of 2 - e.g. have "value")
        }   // end -- for (environment variable)

    } else if err != nil {
        return "", err
    }
    return "", fmt.Errorf("something is wrong, seems not found through 'filepath' AND 'environment variable'")
}

