package diff

import (
	"encoding/json"
	"fmt"
)

// PrintPretty prints the differences in a pretty format.
func PrintPretty(differences map[string]interface{}) {
	output, _ := json.MarshalIndent(differences, "", "  ")
	fmt.Println(string(output))
}
