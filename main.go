package main

// Imports
import (
    "bufio"
    "fmt"
    "io/ioutil"
    "os"
    "os/exec"
    "strings"
)
//
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
	// For the sake of simplicity, we'll just print the new content. In a real implementation, you would find the note with the given title and update the content.
	fmt.Printf("Editing note with title: %s and new content: %s\n", title, content)
	return nil
}

func deleteNote(title string) error {
	fmt.Printf("Deleting note with title: %s\n", title)
	return nil
}

func sendNote(title, target string) error {
    // Read the encrypted note from the notes/ directory
    encryptedNote, err := ioutil.ReadFile(fmt.Sprintf("notes/%s.txt", title))
    if err != nil {
        return err
    }

    // Create a temporary file to store the encrypted note
    tempFile, err := ioutil.TempFile("", "encrypted_note_*.txt")
    if err != nil {
        return err
    }
    defer tempFile.Close()
    defer os.Remove(tempFile.Name())

    // Write the encrypted note to the temporary file
    _, err = tempFile.Write(encryptedNote)
    if err != nil {
        return err
    }

    // Execute the SCP command to send the note to the target IP address
    cmd := exec.Command("scp", tempFile.Name(), fmt.Sprintf("%s:%s.txt", target, title))
    cmd.Stdin = os.Stdin
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    err = cmd.Run()
    if err != nil {
        return err
    }

    fmt.Printf("Successfully sent note '%s' to %s\n", title, target)
    return nil
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

// main function to parse command-line arguments and execute the corresponding function
func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: note <command> [<args>]")
		return
	}

	command := os.Args[1]
	args := os.Args[2:]

	switch command {
	case "new":

		// TODO: need exception handling for already existing note...
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
	case "send":
    		if len(args) < 2 {
        		fmt.Println("Usage: note send <title> <target>")
        		return
    		}

    		title := args[0]
    		target := args[1]
    		err := sendNote(title, target)
    		if err != nil {
        		fmt.Println(err)
        		return
    		}
    		fmt.Println("Note sent")
	}
}
