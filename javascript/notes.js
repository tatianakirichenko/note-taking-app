const fs = require('fs');
const readline = require('readline');

const NOTES_FILE = 'notes.json';

const rl = readline.createInterface({
    input: process.stdin,
    output: process.stdout
});

function loadNotes() {
    if (!fs.existsSync(NOTES_FILE)) return [];
    const data = fs.readFileSync(NOTES_FILE, 'utf8');
    return JSON.parse(data);
}

function saveNotes(notes) {
    fs.writeFileSync(NOTES_FILE, JSON.stringify(notes, null, 2));
}

function addNote(notes) {
    rl.question('Enter your note: ', (text) => {
        if (!text.trim()) {
            console.log('Note cannot be empty.');
            return showMenu(notes);
        }
        const note = {
            id: notes.length + 1,
            text: text.trim(),
            created: new Date().toISOString()
        };
        notes.push(note);
        saveNotes(notes);
        console.log(`✅ Note #${note.id} added.`);
        showMenu(notes);
    });
}

function listNotes(notes) {
    if (notes.length === 0) {
        console.log('📭 No notes yet.');
    } else {
        console.log('\n📋 Your notes:');
        notes.forEach(note => {
            console.log(`  [${note.id}] ${note.text} (created: ${note.created.slice(0,10)})`);
        });
    }
    showMenu(notes);
}

function deleteNote(notes) {
    if (notes.length === 0) {
        console.log('📭 No notes to delete.');
        return showMenu(notes);
    }
    listNotes(notes);
    rl.question('Enter note ID to delete: ', (answer) => {
        const id = parseInt(answer);
        const index = notes.findIndex(n => n.id === id);
        if (index !== -1) {
            notes.splice(index, 1);
            saveNotes(notes);
            console.log(`🗑️ Note #${id} deleted.`);
        } else {
            console.log('Note not found.');
        }
        showMenu(notes);
    });
}

function showMenu(notes) {
    console.log('\n' + '='.repeat(30));
    console.log('📓 NOTE TAKING APP');
    console.log('1. Add a note');
    console.log('2. List notes');
    console.log('3. Delete a note');
    console.log('4. Exit');
    rl.question('Choose: ', (choice) => {
        switch(choice) {
            case '1':
                addNote(notes);
                break;
            case '2':
                listNotes(notes);
                break;
            case '3':
                deleteNote(notes);
                break;
            case '4':
                console.log('Goodbye! 👋');
                rl.close();
                break;
            default:
                console.log('Invalid choice.');
                showMenu(notes);
        }
    });
}

const notes = loadNotes();
showMenu(notes);
