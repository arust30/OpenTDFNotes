# OpenTDFNotes
_Small PoC for implementing basic encryption using OpenTDF and Go for private network sharing of plaintext notes_

## Motivation / Philosophy
This repository is a small Proof-of-Concept for managing private notes on a local network and device using the command line and encrypting/decrypting them using the OpenTDF standard. The primary scripts of this PoC are written in Go with bash scripts for aid in setup/teardown. A primary use-case of this tool would be for managing one's own personal notes across mutliple devices on a home network securely. 

## Directory Structure
```
OpenTDFNotes/
├── main.go
├── config.yaml
├── go.mod
├── go.sum
├── README.md
└── scripts/
    ├── setup.sh
    └── teardown.sh
```

## Features
_Functions intended to include, and those that are primarily pursued during development_
 - Simple CLI interface 
 - Local-first
 - Creating, editing, deleting notes
 - Encryption/Decryption using OpenTDF

## Anti-Features
_Functions purposefully omitted, or to be implemented at a further stage of development_
- Syncing/Version Control
- Search
- Sharing
- Text Formatting
### Future Features:
 - Browser Integration (host a little GUI locally)
 - Flags (tagging notes by category, etc)

## Installation / Setup
### Requirements:
    - Go (see installation doc for you Linux repo or macOS)
    - opentdf/otdfctl ([GitHub](https://github.com/opentdf/otdfctl/))
        - `go get github.com/opentdf/otdfctl@upgrade`
        - or install as binary:
    - make sure to chmod necessary scripts


## Usage / Demo
Here are the availble commands and how they are to be invoked:
 - `note new <title>`: Creates a new encrypted note with the specified title. The note content is read from stdin
 - `note edit <title>`: Edits an existing encrypted note with the specified title. The new note content is read from stdin
 - `note delete <title>`: Deletes an existing encrypted note with the specified title
 - `note list`: Lists the current notes

## Sample Usage
 - `echo "This is a test note" | note new "Sample Note 1"`
 - `echo "This is an update to a note" | note edit "Sample Note 1"`
 - `note delete "Sample Note 1"`


