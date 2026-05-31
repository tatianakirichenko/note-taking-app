#!/usr/bin/env python3
import json
import os
from datetime import datetime

NOTES_FILE = "notes.json"

def load_notes():
    if not os.path.exists(NOTES_FILE):
        return []
    with open(NOTES_FILE, 'r') as f:
        return json.load(f)

def save_notes(notes):
    with open(NOTES_FILE, 'w') as f:
        json.dump(notes, f, indent=2)

def add_note(notes):
    text = input("Enter your note: ").strip()
    if not text:
        print("Note cannot be empty.")
        return
    note = {
        "id": len(notes) + 1,
        "text": text,
        "created": datetime.now().isoformat()
    }
    notes.append(note)
    save_notes(notes)
    print(f"✅ Note #{note['id']} added.")

def list_notes(notes):
    if not notes:
        print("📭 No notes yet.")
        return
    print("\n📋 Your notes:")
    for note in notes:
        print(f"  [{note['id']}] {note['text']} (created: {note['created'][:10]})")

def delete_note(notes):
    list_notes(notes)
    if not notes:
        return
    try:
        idx = int(input("Enter note ID to delete: "))
        note = next((n for n in notes if n['id'] == idx), None)
        if note:
            notes.remove(note)
            save_notes(notes)
            print(f"🗑️ Note #{idx} deleted.")
        else:
            print("Note not found.")
    except ValueError:
        print("Invalid ID.")

def main():
    notes = load_notes()
    while True:
        print("\n" + "="*30)
        print("📓 NOTE TAKING APP")
        print("1. Add a note")
        print("2. List notes")
        print("3. Delete a note")
        print("4. Exit")
        choice = input("Choose: ").strip()
        if choice == '1':
            add_note(notes)
        elif choice == '2':
            list_notes(notes)
        elif choice == '3':
            delete_note(notes)
        elif choice == '4':
            print("Goodbye! 👋")
            break
        else:
            print("Invalid choice.")

if __name__ == "__main__":
    main()
