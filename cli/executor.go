package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/RedCuckoo/merkle-tree-verifier/client/src/client"
)

type CommandExecutor struct {
	Client *client.Service
	logger *log.Logger
}

func NewCommandExecutor(client *client.Service) *CommandExecutor {
	logger := log.New(os.Stdout, "[cli] ", log.LstdFlags)
	return &CommandExecutor{Client: client, logger: logger}
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
			e.logCommandExecutionError(err, GENERATE_COMMAND)
		}
		return
	case RESET_COMMAND:
		err := e.resetCommand(blocks)
		if err != nil {
			e.logCommandExecutionError(err, RESET_COMMAND)
		}
		return
	case UNLOAD_COMMAND:
		err := e.unloadCommand(blocks)
		if err != nil {
			e.logCommandExecutionError(err, UNLOAD_COMMAND)
		}
		return
	case LIST_COMMAND:
		err := e.listCommand(blocks)
		if err != nil {
			e.logCommandExecutionError(err, LIST_COMMAND)
		}
		return
	case DOWNLOAD_COMMAND:
		err := e.downloadCommand(blocks)
		if err != nil {
			e.logCommandExecutionError(err, DOWNLOAD_COMMAND)
		}
		return
	default:
		e.logger.Println("unknown command")
	}
}

func (e *CommandExecutor) logCommandExecutionError(err error, command string) {
	e.logger.Println(
		err.Error(),
		fmt.Sprintf("; while executing \"%s\" command", command),
	)
}

func (e *CommandExecutor) generateCommand(command []string) error {
	if len(command) == 0 {
		return nil
	}

	if command[0] != GENERATE_COMMAND {
		return ErrInternal
	}
	if len(command) == 1 || len(command) == 2 && command[1] == HELP_COMMAND {
		fmt.Fprintf(os.Stdout, "\nUsage: %s AMOUNT_OF_FILES_TO_GENERATE\n\n", GENERATE_COMMAND)
		return nil
	}

	if len(command) == 2 {
		amount, err := strconv.Atoi(command[1])
		if err != nil {
			fmt.Fprintf(os.Stdout, "\nUsage: %s AMOUNT_OF_FILES_TO_GENERATE\n\n", GENERATE_COMMAND)
			return ErrInvalidRequest
		}

		if amount <= 0 {
			fmt.Fprintf(os.Stdout, "\nUsage: %s AMOUNT_OF_FILES_TO_GENERATE\n"+
				"\tAMOUNT_OF_FILES_TO_GENERATE > 0\n\n", GENERATE_COMMAND)
			return ErrInvalidRequest
		}

		return e.Client.GenerateFiles(amount)
	}

	return ErrInvalidRequest
}

func (e *CommandExecutor) listCommand(command []string) error {
	if len(command) == 0 {
		return nil
	}

	if command[0] != LIST_COMMAND {
		return ErrInternal
	}
	if len(command) == 1 || len(command) == 2 && command[1] == HELP_COMMAND {
		fmt.Fprintf(os.Stdout, "\nUsage: %s [OPTIONS]\n", LIST_COMMAND)
		fmt.Fprintf(os.Stdout, "\nOptions:\n"+
			"\t--local  	Display local files\n"+
			"\t--remote 	Display remote files on the server\n\n")
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

	return ErrInvalidRequest
}

func (e *CommandExecutor) downloadCommand(command []string) error {
	if len(command) == 0 {
		return nil
	}

	if command[0] != DOWNLOAD_COMMAND {
		return ErrInternal
	}
	if len(command) == 1 || len(command) == 2 && command[1] == HELP_COMMAND {
		fmt.Fprintf(os.Stdout, "\nUsage: %s FILE_INDEX\n\n", DOWNLOAD_COMMAND)
		return nil
	}

	if len(command) == 2 {
		amount, err := strconv.Atoi(command[1])
		if err != nil {
			return err
		}

		if amount <= 0 {
			fmt.Fprintf(os.Stdout, "\nUsage: %s FILE_INDEX\n"+
				"\tFILE_INDEX > 0 (according to indexing from \"list\" command\n\n", DOWNLOAD_COMMAND)
			return ErrInvalidRequest
		}

		return e.Client.Download(uint64(amount))
	}

	return ErrInvalidRequest
}

func (e *CommandExecutor) resetCommand(command []string) error {
	if len(command) == 0 {
		return nil
	}

	if command[0] != RESET_COMMAND {
		return ErrInternal
	}

	if len(command) == 1 {
		return e.Client.Reset()
	}

	return ErrInvalidRequest
}

func (e *CommandExecutor) unloadCommand(command []string) error {
	if len(command) == 0 {
		return nil
	}

	if command[0] != UNLOAD_COMMAND {
		return ErrInternal
	}

	if len(command) == 1 {
		return e.Client.Unload()
	}

	return ErrInvalidRequest
}
