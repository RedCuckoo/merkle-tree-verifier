package main

import (
	prompt "github.com/c-bata/go-prompt"
)

const (
	GENERATE_COMMAND = "generate"
	RESET_COMMAND    = "reset"
	UNLOAD_COMMAND   = "unload"
	DOWNLOAD_COMMAND = "download"
	LIST_COMMAND     = "list"
	HELP_COMMAND     = "help"
	EXIT_COMMAND     = "exit"
)

type Commands struct {
	CLISuggestions    []prompt.Suggest
	CLISubSuggestions map[string][]prompt.Suggest
	HelpSuggestion    []prompt.Suggest
}

func New() Commands {
	return Commands{
		CLISuggestions: []prompt.Suggest{
			{Text: GENERATE_COMMAND, Description: "Generate multiple files"},
			{
				Text:        UNLOAD_COMMAND,
				Description: "Upload local files to the remote server and delete local copies",
			},
			{Text: LIST_COMMAND, Description: "List local or remote files"},
			{
				Text:        DOWNLOAD_COMMAND,
				Description: "Download selected file from the remote server and verify integrity",
			},
			{Text: RESET_COMMAND, Description: "Reset client and server to initial state (empty)"},
			{Text: EXIT_COMMAND, Description: "Exit"},
		},
		CLISubSuggestions: map[string][]prompt.Suggest{
			GENERATE_COMMAND: {
				prompt.Suggest{Text: HELP_COMMAND, Description: "Display help"},
			},
			LIST_COMMAND: {
				prompt.Suggest{Text: HELP_COMMAND, Description: "Display help"},
				prompt.Suggest{Text: "--local", Description: "List local files"},
				prompt.Suggest{Text: "--remote", Description: "List remote files"},
			},
			DOWNLOAD_COMMAND: {
				prompt.Suggest{Text: HELP_COMMAND, Description: "Display help"},
			},
		},
		HelpSuggestion: []prompt.Suggest{
			{HELP_COMMAND, "get help on command"},
		},
	}
}

func (c *Commands) GetCLISuggestions() []prompt.Suggest {
	return c.CLISuggestions
}

func (c *Commands) GetCLISubSuggestions() map[string][]prompt.Suggest {
	return c.CLISubSuggestions
}

func (c *Commands) GetHelpSuggestion() []prompt.Suggest {
	return c.HelpSuggestion
}

func (c *Commands) IsCLICommand(kw string) bool {
	for _, cmd := range c.CLISuggestions {
		if cmd.Text == kw {
			return true
		}
	}

	return false
}

func (c *Commands) IsCLISubCommand(kw string) ([]prompt.Suggest, bool) {
	val, ok := c.CLISubSuggestions[kw]
	return val, ok
}
