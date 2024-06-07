package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
//	"strings"
)

func main() {
	// Example usage of the encryptNote() and decryptNote() functions
	note := "This is a secret note"
	encryptedNote, err := encryptNote(note)
	if err != nil {
		fmt.Println("Error encrypting note:", err)
		return
	}
	fmt.Println("Encrypted note:", encryptedNote)

	decryptedNote, err := decryptNote(encryptedNote)
	if err != nil {
		fmt.Println("Error decrypting note:", err)
		return
	}
	fmt.Println("Decrypted note:", decryptedNote)
}

// encryptNote() encrypts a note using the otdfctl binary and returns the encrypted note as a string
func encryptNote(note string) (string, error) {
	// Create a temporary file to store the note
	tempFile, err := ioutil.TempFile("", "note")
	if err != nil {
		return "", err
	}
	defer os.Remove(tempFile.Name()) // Clean up the temporary file when we're done

	// Write the note to the temporary file
	_, err = tempFile.WriteString(note)
	if err != nil {
		return "", err
	}

	// Create a new command to execute the otdfctl binary and encrypt the note
	cmd := exec.Command("otdfctl", "encrypt", tempFile.Name())

	// Run the command and capture the output
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}

	// Return the encrypted note as a string
	return string(output), nil
}

// decryptNote() decrypts an encrypted note using the otdfctl binary and returns the decrypted note as a string
func decryptNote(encryptedNote string) (string, error) {
	// Create a temporary file to store the encrypted note
	tempFile, err := ioutil.TempFile("", "note.tdf")
	if err != nil {
		return "", err
	}
	defer os.Remove(tempFile.Name()) // Clean up the temporary file when we're done

	// Write the encrypted note to the temporary file
	_, err = tempFile.WriteString(encryptedNote)
	if err != nil {
		return "", err
	}

	// Create a new command to execute the otdfctl binary and decrypt the note
	cmd := exec.Command("otdfctl", "decrypt", tempFile.Name())

	// Run the command and capture the output
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}

	// Return the decrypted note as a string
	return string(output), nil
}

