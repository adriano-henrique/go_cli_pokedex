package repl

import (
	"errors"
	"fmt"

	api "github.com/adriano-henrique/internal"
)

func parsePreviousUrl(url any) *string {
	switch v := url.(type) {
	case string:
		return &v
	default:
		return nil
	}
}

func commandMapb(mapConfig *MapConfig) error {
	previousUrl := mapConfig.Previous
	if previousUrl == nil {
		fmt.Println("No values back, use map to go forward on the list of areas")
		return errors.New("NO VALUES BACK")
	}
	locationAreaObject, err := api.GetInformation(*previousUrl)
	if err != nil {
		return err
	}

	results := locationAreaObject.Results
	for _, result := range results {
		fmt.Println(result.Name)
	}

	mapConfig.Previous = parsePreviousUrl(locationAreaObject.Previous)
	mapConfig.Next = previousUrl

	return nil
}
