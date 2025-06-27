package main

import (
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

func generateCompletions(flag *pflag.FlagSet, commandName string) error {
	command := &cobra.Command{
		Use: "completions",
	}
	command.Flags().AddFlagSet(flag)
	os.MkdirAll("completions/bash", 0755)
	os.MkdirAll("completions/zsh", 0755)
	os.MkdirAll("completions/fish", 0755)
	os.MkdirAll("completions/ps1", 0755)
	command.GenBashCompletionFileV2(filepath.Join("completions/bash", commandName), true)
	command.GenZshCompletionFile(filepath.Join("completions/zsh", commandName))
	command.GenFishCompletionFile(filepath.Join("completions/fish", commandName), true)
	command.GenPowerShellCompletionFile(filepath.Join("completions/ps1", commandName))
	return nil
}
