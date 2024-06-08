package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: note <command> [<args>]")
		return
	}

	command := os.Args[1]
	args := os.Args[2:]

	switch command {
	case "new":
		if len(args) < 1 {
			fmt.Println("Usage: note new <title>")
			return
		}

		title := args[0]
		fmt.Print("Enter note content: ")
		content, _ := bufio.NewReader(os.Stdin).ReadString('\n')

		encryptedContent, err := encryptContent(content)
		if err != nil {
			fmt.Println(err)
			return
		}

		err = createNote(title, encryptedContent)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("Note created")

	case "edit":
		if len(args) < 1 {
			fmt.Println("Usage: note edit <title>")
			return
		}

		title := args[0]
		fmt.Print("Enter new note content: ")
		content, _ := bufio.NewReader(os.Stdin).ReadString('\n')

		encryptedContent, err := encryptContent(content)
		if err != nil {
			fmt.Println(err)
			return
		}

		err = editNote(title, encryptedContent)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("Note updated")

	case "delete":
		if len(args) < 1 {
			fmt.Println("Usage: note delete <title>")
			return
		}

		title := args[0]
		err := deleteNote(title)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("Note deleted")

	case "list":
		notes, err := listNotes()
		if err != nil {
			fmt.Println(err)
			return
		}

		for _, note := range notes {
			fmt.Println(note)
		}

	default:
		fmt.Printf("Unknown command: %s\n", command)
	}
}

func encryptContent(content string) (string, error) {
	cmd := exec.Command("opentdf", "encrypt", "--plaintext", content)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("failed to encrypt content: %w", err)
	}

	encryptedContent := strings.TrimSpace(string(output))
	return encryptedContent, nil
}

func createNote(title, content string) error {
	// TODO: Implement the logic to create a note with the given title and content
	return nil
}

func editNote(title, content string) error {
	// TODO: Implement the logic to edit a note with the given title and content
	return nil
}

func deleteNote(title string) error {
	// TODO: Implement the logic to delete a note with the given title
	return nil
}

func listNotes() ([]string, error) {
	// TODO: Implement the logic to list all the notes
	return nil, nil
}

