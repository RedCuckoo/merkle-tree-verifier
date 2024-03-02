package main

import (
	"github.com/c-bata/go-prompt"
	"regexp"
	"strings"
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

		if command == GENERATE_COMMAND {
			if len(strings.Split(d.Text, " ")) > 2 {
				return []prompt.Suggest{}
			}
			return cliCommands.GetHelpSuggestion()
		}

		if command == LIST_COMMAND {
			if len(strings.Split(d.Text, " ")) > 2 {
				return []prompt.Suggest{}
			}
			if val, ok := cliCommands.IsCLISubCommand(command); ok {
				return prompt.FilterHasPrefix(val, word, true)
			}
		}

		//if command == "start" {
		//	return containerListCompleter(true)
		//}
		//
		//if command == "run" {
		//	if word == "-p" {
		//		if len(portMappingSuggestions) > 0 {
		//			return portMappingSuggestions
		//		}
		//
		//		return portMappingSuggestion()
		//	}
		//
		//	if len(suggestedImages) > 0 {
		//		return suggestedImages
		//	}
		//
		//	return imagesSuggestion()
		//}
		//
		//if command == "pull" {
		//	if strings.Index(word, ":") != -1 || strings.Index(word, "@") != -1 {
		//		return []prompt.Suggest{}
		//	}
		//
		//	if word == "" || len(word) > 2 {
		//		if len(strings.Split(d.Text, " ")) > 2 {
		//			return []prompt.Suggest{}
		//		}
		//		return getFromCache(word)
		//	}
		//
		//	return []prompt.Suggest{}
		//}
		if val, ok := cliCommands.IsCLISubCommand(command); ok {
			return prompt.FilterHasPrefix(val, word, true)
		}
	}
	return prompt.FilterHasPrefix(cliCommands.GetCLISuggestions(), word, true)
}
