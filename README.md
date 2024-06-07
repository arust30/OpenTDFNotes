# OpenTDFNotes
_Small PoC for implementing basic encryption using OpenTDF and Go for private network sharing of plaintext notes_

## Motivation / Philosophy
This repository is a small Proof-of-Concept for managing private notes on a local network and device using the command line and encrypting/decrypting them using the OpenTDF standard. The primary scripts of this PoC are written in Go with bash scripts for aid in setup/teardown. A primary use-case of this tool would be for managing one's own personal notes across mutliple devices on a home network securely. 

## Directory Structure
```
OpenTDFNotes/
├── cmd/
│   └── secure-notes/
│       └── main.go
├── pkg/
│   └── notes/
│       ├── note.go
│       └── notes.go
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


## Usage / Demo
 - Commands:
    - `note new <title>`
    - `note edit <title>`
    - `note delete <title>`


