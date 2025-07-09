package cmd

import (
    "os"

    "github.com/spf13/cobra"
)

func init() {
    rootCmd.AddCommand(completionCmd)
}

var completionCmd = &cobra.Command{
    Use:   "completion",
    Short: "Generate shell completion scripts",
}

func init() {
    completionCmd.AddCommand(&cobra.Command{
        Use:   "zsh",
        Short: "Generate zsh completion script",
        Run: func(cmd *cobra.Command, args []string) {
            rootCmd.GenZshCompletion(os.Stdout)
        },
    })
    // 必要なら bash, fish, powershell も同様に追加できます
}