package diff

import (
	"fmt"

	"github.com/fatih/color"
)

// ColorPrint prints the differences in color.
func ColorPrint(differences map[string]interface{}) {
	yellow := color.New(color.FgYellow).SprintFunc()
	red := color.New(color.FgRed).SprintFunc()
	green := color.New(color.FgGreen).SprintFunc()

	for k, v := range differences {
		diffMap, ok := v.(map[string]interface{})
		if !ok {
			fmt.Printf("%s - Diff: %v\n", yellow(k), v)
			continue
		}
		fromVal := diffMap["from"]
		toVal := diffMap["to"]

		fmt.Printf("%s - Diff: from=%s to=%s\n",
			yellow(k),
			red(fmt.Sprintf("%v", fromVal)),
			green(fmt.Sprintf("%v", toVal)),
		)
	}
}
