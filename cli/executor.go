package main

import (
	"fmt"
	"github.com/RedCuckoo/merkle-tree-verifier/client/src/service"
	"os"
	"strconv"
	"strings"
)

type CommandExecutor struct {
	Client *service.ClientService
}

func NewCommandExecutor(client *service.ClientService) *CommandExecutor {
	return &CommandExecutor{Client: client}
}

func (e *CommandExecutor) ExecuteCommand(in string) {
	in = strings.TrimSpace(in)

	blocks := strings.Split(in, " ")
	switch blocks[0] {
	case EXIT_COMMAND:
		os.Exit(0)
		return
	case GENERATE_COMMAND:
		err := e.generateCommand(blocks)
		if err != nil {
			fmt.Println(err.Error(), "while executing %s command", GENERATE_COMMAND)
		}
		return
	case RESET_COMMAND:
		err := e.Client.Reset()
		if err != nil {
			fmt.Println(fmt.Errorf(err.Error(), "while executing %s command", RESET_COMMAND))
		}
		return
	case UNLOAD_COMMAND:
		err := e.Client.Unload()
		if err != nil {
			fmt.Println(fmt.Errorf(err.Error(), "while executing %s command", UNLOAD_COMMAND))
		}
		return
	case LIST_COMMAND:
		err := e.listCommand(blocks)
		if err != nil {
			fmt.Println(fmt.Errorf(err.Error(), "while executing %s command", DOWNLOAD_COMMAND))
		}
		return
	case DOWNLOAD_COMMAND:
		err := e.Client.Download()
		if err != nil {
			fmt.Println(fmt.Errorf(err.Error(), "while executing %s command", DOWNLOAD_COMMAND))
		}
		return
	default:
		fmt.Println("unknown command")
	}
}

func (e *CommandExecutor) generateCommand(command []string) error {
	if len(command) == 0 {
		return nil
	}

	if command[0] != GENERATE_COMMAND {
		return ErrInternal
	}
	if len(command) == 1 || len(command) == 2 && command[1] == HELP_COMMAND {
		fmt.Printf("\nUsage: %s AMOUNT_OF_FILES_TO_GENERATE\n\n", GENERATE_COMMAND)
		return nil
	}

	if len(command) == 2 {
		amount, err := strconv.Atoi(command[1])
		if err != nil {
			return err
		}
		return e.Client.GenerateFiles(amount)
	}

	return nil
}

func (e *CommandExecutor) listCommand(command []string) error {
	if len(command) == 0 {
		return nil
	}

	if command[0] != LIST_COMMAND {
		return ErrInternal
	}
	if len(command) == 1 || len(command) == 2 && command[1] == HELP_COMMAND {
		fmt.Printf("\nUsage: %s [OPTIONS]\n", LIST_COMMAND)
		fmt.Printf("\nOptions:\n" +
			"\t--local  	Display local files\n" +
			"\t--remote 	Display remote files on the server\n")
		return nil
	}

	if len(command) == 2 {
		switch command[1] {
		case "--local":
			return e.Client.ListLocal()
		case "--remote":
			return e.Client.ListRemote()
		}
	}

	return nil
}
