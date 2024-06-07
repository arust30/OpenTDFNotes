# OpenTDFNotes
_Small PoC for implementing basic encryption using OpenTDF Library using Go_

## Motivation / Philosophy
This repository is a small Proof-of-Concept for managing private notes on a local device using the command line and encrypting/decrypting them using the OpenTDF standard.

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
_functions intended to include, and those that are primarily pursued during development_
 - Simple CLI interface, flags  
 - Local-first
 - Creating, editing, deleting notes
 - Encryption/Decryption using OpenTDF

## Anti-Features
_functions purposefully omitted, or to be implemented at a further stage of development_
- Syncing/Version Control
- Search
- Sharing
- Text Formatting

## Installation / Setup

## Usage / Demo
