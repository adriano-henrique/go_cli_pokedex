package repl

import (
	"errors"
	"fmt"

	api "github.com/adriano-henrique/internal"
)

func commandMap(mapConfig *MapConfig) error {
	currentUrl := mapConfig.Next
	if currentUrl == nil {
		fmt.Println("End of values, use mapb to get it back")
		return errors.New("NO NEXT VALUE")
	}
	locationAreaObject, err := api.GetInformation(*mapConfig.Next)
	if err != nil {
		return err
	}
	results := locationAreaObject.Results
	for _, result := range results {
		fmt.Println(result.Name)
	}

	mapConfig.Previous = currentUrl
	mapConfig.Next = &locationAreaObject.Next

	return nil
}
