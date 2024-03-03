package main

import (
	"regexp"
	"strings"

	prompt "github.com/c-bata/go-prompt"
)

var cliCommands = New()

var incompleteCommands = regexp.MustCompile(`(?P<command>generate|list|download)\s{1}`)

func getRegexGroups(text string) map[string]string {
	if !incompleteCommands.Match([]byte(text)) {
		return nil
	}

	match := incompleteCommands.FindStringSubmatch(text)
	result := make(map[string]string)
	for i, name := range incompleteCommands.SubexpNames() {
		if i != 0 && name != "" {
			result[name] = match[i]
		}
	}
	return result
}

func completer(d prompt.Document) []prompt.Suggest {
	word := d.GetWordBeforeCursor()
	group := getRegexGroups(d.Text)
	if group != nil {
		command := group["command"]

		if command == GENERATE_COMMAND || command == LIST_COMMAND || command == DOWNLOAD_COMMAND {
			if len(strings.Split(d.Text, " ")) > 2 {
				return []prompt.Suggest{}
			}
		}

		if val, ok := cliCommands.IsCLISubCommand(command); ok {
			return prompt.FilterHasPrefix(val, word, true)
		}
	}
	return prompt.FilterHasPrefix(cliCommands.GetCLISuggestions(), word, true)
}
