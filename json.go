package gomine

import (
	"encoding/json"
	"os"
)

// ExportToJSON takes a mapping and writes to filename
func ExportToJSON(filename string, mapping interface{}) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	// convert map[string]interface{} to []byte
	res, err := json.Marshal(mapping)
	if err != nil {
		return err
	}

	// write json data to file
	file.Write(res)
	return nil
}
