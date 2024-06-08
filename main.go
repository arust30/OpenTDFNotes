package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type Note struct {
	Title   string
	Content string
}

func createNote(title, content string) error {
	note := &Note{
		Title:   title,
		Content: content,
	}
	fmt.Printf("Creating note with title: %s and content: %s\n", note.Title, note.Content)
	return nil
}

func editNote(title, content string) error {
	// For the sake of simplicity, we'll just print the new content.
	// In a real implementation, you would find the note with the given title
	// and update its content.
	fmt.Printf("Editing note with title: %s and new content: %s\n", title, content)
	return nil
}

func deleteNote(title string) error {
	fmt.Printf("Deleting note with title: %s\n", title)
	return nil
}

func listNotes() ([]string, error) {
	fmt.Println("Listing all notes (placeholder)")
	return []string{"Note 1", "Note 2", "Note 3"}, nil
}

func encryptContent(content string) (string, error) {
	cmd := exec.Command("otdfctl", "encrypt", "--host", "http://localhost:8888")
	cmd.Stdin = strings.NewReader(content)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("failed to encrypt content: %w", err)
	}

	encryptedContent := strings.TrimSpace(string(output))
	return encryptedContent, nil
}

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
		_, err := listNotes()
		if err != nil {
			fmt.Println(err)
		}
	}
}
