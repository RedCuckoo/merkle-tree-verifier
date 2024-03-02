package main

import (
	"github.com/RedCuckoo/merkle-tree-verifier/client/src/service"
	"github.com/c-bata/go-prompt"
)

func main() {
	client := service.NewClientService()
	executor := NewCommandExecutor(client)

	for {
		p := prompt.Input(
			"merkle-tree-verifier> ",
			completer,
			prompt.OptionTitle("merkle-tree-verifier CLI"),
			prompt.OptionPrefixBackgroundColor(prompt.Purple),
		)

		executor.ExecuteCommand(p)

	}
}
