package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/sorahashiroi/difson/pkg/diff"
	"github.com/spf13/cobra"
)

var (
	color       bool
	brief       bool
	pretty      bool
	completions bool // この変数を追加
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "difson",
	Short: "Compare two JSON files and show differences.",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		file1 := args[0]
		file2 := args[1]
		obj1, err := diff.LoadJSON(file1)
		if err != nil {
			fmt.Println("Error loading JSON file1:", err)
			os.Exit(1)
		}
		obj2, err := diff.LoadJSON(file2)
		if err != nil {
			fmt.Println("Error loading JSON file2:", err)
			os.Exit(1)
		}

		// breifフラグが指定されている場合、簡易比較を行う
		if brief {
			same := diff.CompareJSONBrief(obj1, obj2)
			// if err != nil {
			// 	fmt.Println("Error:", err)
			// 	return
			// }
			if same {
				fmt.Println("Files are the same.")
			} else {
				fmt.Println("Files are different.")
			}
			return
		}

		differences := diff.GetDiff(obj1, obj2)

		if pretty {
			// prettyフラグが指定されている場合、差分を整形して表示
			diff.PrintPretty(differences)
		} else if color {
			diff.ColorPrint(differences)
		} else {
			diffJSON, err := json.MarshalIndent(differences, "", "  ")
			if err != nil {
				fmt.Println("Error formatting differences:", err)
			} else {
				fmt.Println(string(diffJSON))
			}
		}

	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolVarP(&color, "color", "c", false, "Show colorized output")
	rootCmd.Flags().BoolVarP(&brief, "brief", "b", false, "Print a simple message")
	rootCmd.Flags().BoolVarP(&pretty, "pretty", "p", false, "Print diff with indentation and formatting")
	rootCmd.Flags().BoolVarP(&completions, "generate-completions", "", false, "generate completions")
	rootCmd.Flags().MarkHidden("generate-completions")
}
