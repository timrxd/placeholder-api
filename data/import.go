package data

import (
	"encoding/json"
	"os"

	"github.com/timrxd/placeholder-api/api"
)

// ImportData reads in data from file and saves it to server database
func ImportData(filename string) (map[int]*api.Item, error) {

	// Read in json from file
	file, _ := os.ReadFile(filename)
	var data []*api.Item
	err := json.Unmarshal(file, &data)
	if err != nil {
		return nil, err
	}

	// Convert list to map
	newDB := map[int]*api.Item{}
	for _, v := range data {
		newDB[v.ID] = v
	}

	return newDB, nil
}
