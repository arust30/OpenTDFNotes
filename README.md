# OpenTDFNotes
_Secure plaintext note sharing using OpenTDF & Go_

## Motivation
This repository is a naive and simple Proof-of-Concept for sharing encrypted notes on a (local) network using the CLI

The primary use-case for this tool would be for an individual to manage their personal notes across multile devices on a home network, perhaps (it is assumed key sharing is secure).

**This naive implementation is not secure and is not meant for personal and/or production usage**

## Directory Structure
_The following is the directory structure of this project along with descriptions_
```
OpenTDFNotes/
├── main.go
├── config.yaml
├── go.mod
├── go.sum
├── note `: bash script for invoking binary and where configuration settings/env variables may be instantiated
├── note`: The main.go code is compiled and the binary is placed here
├── notes/ `: Here is where encrypted notes would go
├── README.md
```

## Features
_Functions intended to include, and those that are primarily pursued during development_
 - Simple CLI interface 
 - Local-first
 - Creating, editing, deleting notes
 - Encryption/Decryption using OpenTDF

## Anti-Features
_Functions purposefully omitted, or to be implemented at a further stage of development_
 - Syncing/Version Control of shared files, viewable history
 - More robust search feature
 - Browser Integration (host a little GUI locally)
 - Flags (tagging notes/files by category, etc)

### Requirements:
 - [Go](https://github.com/golang/go)
 - OpenTDF Platform running on machine (requires Docker, among others - see [documentation](https://github.com/opentdf/platform))
 - OpenTDF Command-Line Utility: opentdf/otdfctl - ([GitHub](https://github.com/opentdf/otdfctl/))
    - Install as binary and add to PATH

## Setup Validation & Installation Steps
1. To ensure you have all the prerequisites, invoke the following commands:

`go --version`: Making sure Go is installed

`docker ps`: You should have a few running, make sure none of the daemons are stopped

`otdfctl --help`: Should print out help console for otdfctl utility

2. Then you can clone this repo via http:

```
git clone https://github.com/arust30/OpenTDFNotes
```

3. Then you may have to change permission on the bash executable called `note`:

```
chmod +x note
```

4. Compile Go script

`go build -o bin/`: builds main.go and puts executable binary into bin/ folder

Now you should be able to invoke the CLI as outlined below:

## Usage
Here are the availble commands and how they are to be invoked:
 - `note new <title>`: Creates a new encrypted note with the specified title. The note content is read from stdin
 - `note delete <title>`: Deletes an existing encrypted note with the specified title
 - `note list`: Lists the current notes
 - `note send <target>`: Sends the encrpyted note to IP address via FTP (local only, for now)

### Sample:
 - `echo "This is a test note" | note new "Sample Note 1"`
 - `note delete "Sample Note 1"`
 - `note list`

## TODOs:
 - [ ] Setup/Teardown Bash Scripts
 - [ ] adding `note send <target>` command to send note to device on local network
 - [ ] adding `note edit <title>` command to make notes mutable
 - [ ]


