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
	os.Mkdir("completions/", 0755)
	os.Mkdir("completions/bash", 0755)
	os.Mkdir("completions/zsh", 0755)
	os.Mkdir("completions/fish", 0755)
	os.Mkdir("completions/powershell", 0755)
	command.GenBashCompletionFileV2(filepath.Join("completions/bash", commandName), true)
	command.GenZshCompletionFile(filepath.Join("completions/zsh", commandName))
	command.GenFishCompletionFile(filepath.Join("completions/fish", commandName), true)
	command.GenPowerShellCompletionFile(filepath.Join("completions/ps1", commandName))
	return nil
}
