package smtp

import (
    "encoding/json"
    "io"
    "os"
)

// LoadJSON reads and parses the JSON file into a slice of SMTPConfig
func LoadJSON(path string) ([]SMTPConfig, error) {
    file, err := os.Open(path)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    byteValue, err := io.ReadAll(file)
    if err != nil {
        return nil, err
    }

    var configs []SMTPConfig
    err = json.Unmarshal(byteValue, &configs)
    if err != nil {
        return nil, err
    }

    return configs, nil
}
