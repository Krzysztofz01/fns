# fns
Fuzzy Note Search... for those who have lost their patience for browsing through notes.

## Installation and configuration
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

Below is an example of the `.fns` dotfile. You can place such into the home directory or the directory where the binary is placed.
```json
{
  "note-read-directory-paths": [
    "/home/hello/world/first-dir-to-search-for-notes",
    "/home/hello/second-dir-to-search-for-notes"
  ],
  "note-write-directory-path": "/home/hello/store-new-notes-here", // It can be one of the paths above.
  "editor-path": "", // Leave empty to use the one specified by $EDITOR
  "trim-note": true
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


### Version check
To check the version of fns use:
```shell
fns version
```
