# fns
Fuzzy Note Search... for those who have lost their patience for browsing through notes.

## Installation and configuration
The required tools for compilation are:
- [Git](https://git-scm.com/)
- [Go](https://go.dev/)
- [Taskfile](https://taskfile.dev/)

Clone the repository:
```shell
git clone --depth 1 https://github.com/Krzysztofz01/fns && cd fns
```

Build the tool from source:
```shell
task build
```

Move the binary into destination directory (within $PATH):
```shell
cp bin/fns /usr/local/bin/fns
```

You can now seed a new `.fns` dotfile with the default configuration.
```shell
fns config --default > ~/.fns
```

Below is an example of the `.fns` dotfile. **You must provide the notes read and write paths to make the tool function correctly**.
```json
{
  "note-read-directory-paths": [
    "/home/hello/world/first-dir-to-search-for-notes",
    "/home/hello/second-dir-to-search-for-notes"
  ],
  "note-write-directory-path": "/home/hello/store-new-notes-here", // It can be one of the paths above.
  "editor-path": "", // Leave empty to use the one specified by $EDITOR
  "trim-note": true,
  "skip-invalid-note-files": true
}
```

## Usage

### New note
Add new note via:
```shell
fns add my-new-note.md
```
This will create a new note and open it in the config specified editor.


### Search note
To perform a fuzzy note search use:
```shell
fns search
```
The content of the selected note will be printed to the terminal.

If you want to access the note path instead of the content use:
```shell
fns search --path
```

### Edit note
To edit an existing note use:
```shell
fns edit
```
First, the note must be selected. This is done using fuzzy search, just like in the search command.
Once the note is selected, it will be opened in the config specified editor.


### Version check
To check the version of fns use:
```shell
fns version
```

### Configuration check
To check the current configuration
```shell
fns config
```

To check the default configuration
```shell
fns config --default
```

### Backup notes
To create a backup of all notes that are accessible via fns use:
```shell
fns backup
```

A archive in ZIP format will be created in the current working directory, containing all the notes
with the original file and parent directory name.