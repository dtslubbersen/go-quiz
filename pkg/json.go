package pkg

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"reflect"
)

func ReadMapFromJsonFile[T any, K comparable](filename string) (map[K]*T, error) {
	file, err := os.ReadFile(filename)

	if err != nil {
		return nil, fmt.Errorf("could not read file %s: %v", filename, err)
	}

	var items []T

	if err := json.Unmarshal(file, &items); err != nil {
		return nil, fmt.Errorf("could not unmarshall json %s: %v", filename, err)
	}

	itemMap := make(map[K]*T)
	for _, item := range items {
		v := reflect.ValueOf(item)
		idField := v.FieldByName("Id")
		if !idField.IsValid() {
			log.Fatalf("Missing 'Id' field in struct %T", item)
		}
		id := idField.Interface()
		itemMap[id.(K)] = &item
	}

	return itemMap, nil
}
