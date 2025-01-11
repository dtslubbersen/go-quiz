package pkg

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"reflect"
)

func ReadMapFromJsonFile[T any, K comparable](fileName string, keyFieldName string) (map[K]*T, error) {
	file, err := os.ReadFile(fileName)

	if err != nil {
		return nil, fmt.Errorf("could not read file %s: %v", fileName, err)
	}

	var items []T

	if err := json.Unmarshal(file, &items); err != nil {
		return nil, fmt.Errorf("could not unmarshall json %s: %v", fileName, err)
	}

	itemsMap := make(map[K]*T)
	for _, item := range items {
		v := reflect.ValueOf(item)
		field := v.FieldByName(keyFieldName)
		if !field.IsValid() {
			log.Fatalf("Missing '%s' field in struct %T", keyFieldName, item)
		}

		mapKeyValue := field.Interface()
		itemsMap[mapKeyValue.(K)] = &item
	}

	return itemsMap, nil
}
