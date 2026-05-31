package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Note struct {
	ID      int    `json:"id"`
	Text    string `json:"text"`
	Created string `json:"created"`
}

const notesFile = "notes.json"

func loadNotes() ([]Note, error) {
	if _, err := os.Stat(notesFile); os.IsNotExist(err) {
		return []Note{}, nil
	}
	data, err := os.ReadFile(notesFile)
	if err != nil {
		return nil, err
	}
	var notes []Note
	err = json.Unmarshal(data, &notes)
	return notes, err
}

func saveNotes(notes []Note) error {
	data, err := json.MarshalIndent(notes, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(notesFile, data, 0644)
}

func addNote(notes []Note, reader *bufio.Reader) ([]Note, error) {
	fmt.Print("Enter your note: ")
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)
	if text == "" {
		fmt.Println("Note cannot be empty.")
		return notes, nil
	}
	note := Note{
		ID:      len(notes) + 1,
		Text:    text,
		Created: time.Now().Format(time.RFC3339),
	}
	notes = append(notes, note)
	err := saveNotes(notes)
	if err != nil {
		return notes, err
	}
	fmt.Printf("✅ Note #%d added.\n", note.ID)
	return notes, nil
}

func listNotes(notes []Note) {
	if len(notes) == 0 {
		fmt.Println("📭 No notes yet.")
		return
	}
	fmt.Println("\n📋 Your notes:")
	for _, note := range notes {
		fmt.Printf("  [%d] %s (created: %s)\n", note.ID, note.Text, note.Created[:10])
	}
}

func deleteNote(notes []Note, reader *bufio.Reader) ([]Note, error) {
	if len(notes) == 0 {
		fmt.Println("📭 No notes to delete.")
		return notes, nil
	}
	listNotes(notes)
	fmt.Print("Enter note ID to delete: ")
	input, _ := reader.ReadString('\n')
	id, err := strconv.Atoi(strings.TrimSpace(input))
	if err != nil {
		fmt.Println("Invalid ID.")
		return notes, nil
	}
	index := -1
	for i, n := range notes {
		if n.ID == id {
			index = i
			break
		}
	}
	if index == -1 {
		fmt.Println("Note not found.")
		return notes, nil
	}
	notes = append(notes[:index], notes[index+1:]...)
	err = saveNotes(notes)
	if err != nil {
		return notes, err
	}
	fmt.Printf("🗑️ Note #%d deleted.\n", id)
	return notes, nil
}

func main() {
	notes, err := loadNotes()
	if err != nil {
		fmt.Println("Error loading notes:", err)
		return
	}
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\n" + strings.Repeat("=", 30))
		fmt.Println("📓 NOTE TAKING APP")
		fmt.Println("1. Add a note")
		fmt.Println("2. List notes")
		fmt.Println("3. Delete a note")
		fmt.Println("4. Exit")
		fmt.Print("Choose: ")
		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		switch choice {
		case "1":
			notes, err = addNote(notes, reader)
			if err != nil {
				fmt.Println("Error:", err)
			}
		case "2":
			listNotes(notes)
		case "3":
			notes, err = deleteNote(notes, reader)
			if err != nil {
				fmt.Println("Error:", err)
			}
		case "4":
			fmt.Println("Goodbye! 👋")
			return
		default:
			fmt.Println("Invalid choice.")
		}
	}
}
