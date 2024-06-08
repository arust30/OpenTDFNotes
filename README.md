``` 
 _____                             _   _       _             
/  ___|                           | \ | |     | |            
\ `--.  ___  ___ _   _ _ __ ___   |  \| | ___ | |_ ___  ___  
 `--. \/ _ \/ __| | | | '__/ _ \  | . ` |/ _ \| __/ _ \/ __| 
/\__/ /  __/ (__| |_| | | |  __/  | |\  | (_) | ||  __/\__ \ 
\____/ \___|\___|\__,_|_|  \___|  \_| \_/\___/ \__\___||___/ 

```                                                          
# OpenTDFNotes

## Video Demo
![](DemoVideo.mp4)
https://youtu.be/rIBTVMCRnlc

## Motivation
_Secure plaintext note sharing using OpenTDF & Go_

This repository is a naive and simple Proof-of-Concept for sharing encrypted notes on a (local) network using the CLI

The primary use-case for this tool would be for an individual to manage their personal notes across multile devices on a home network, perhaps (it is assumed key sharing is secure).

**This naive implementation is not secure and is not meant for personal and/or production usage**

## Directory Structure
_The following is the directory structure of this project along with descriptions_
```
OpenTDFNotes/
├── go.mod
├── go.sum
├── main.go : The primary Go file
├── config.yaml : IP addresses, ports, and other configs go here
├── note : bash script for invoking binary and where configuration settings/env variables may be instantiated
├── note: The main.go code is compiled and the binary is placed here
├── notes/ : Here is where encrypted notes would go
|
├── README.md
```

## Features
_Functions intended to include, and those that are primarily pursued during development_
 - Simple CLI interface 
 - Basic encrypted file sharing
 - Creating, deleting, listing, and sharing encrypted notes
 - Encryption/decryption using OpenTDF

## Anti-Features
_Functions purposefully omitted, or to be implemented at a further stage of development_
 - Syncing/Version Control of shared files, viewable history
 - Editing notes (TODO)
 - Non-plaintext formats

### Requirements:
 - [Go](https://github.com/golang/go)
 - OpenTDF Platform running on machine (requires Docker, among others - see [documentation](https://github.com/opentdf/platform))
 - OpenTDF Command-Line Utility: opentdf/otdfctl - ([GitHub](https://github.com/opentdf/otdfctl/))
    - Install as binary and add to PATH
 - ssh daemon running, and scp utility installed on both client and server devices on network (local)

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

5. Test ssh connection(optional)

`ssh -T 192.168.0.112:22`

## Usage
Here are the availble commands and how they are to be invoked:
 - `note new <title>`: Creates a new encrypted note with the specified title. The note content is read from stdin
 - `note delete <title>`: Deletes an existing encrypted note with the specified title
 - `note send <title> <target>`: Sends the encrpyted note to IP address via SCP (OR ssh, local only, for now)

### Sample:
 - `note new "Sample Note 1"`
 - `note delete "Sample Note 1"`
 - `note new "Sample Note 2"`
 - `note send "Sample Note 2" 192.168.0.112:21`

## TODOs:
 - [ ] Setup/Teardown Bash Scripts
 - [ ] Add `note edit <title>` command to make notes mutable
 - [ ] Add tagging feature
 - [ ] Add robust search feature to filter notes
 - [ ] Browser GUI: locally hosted GUI to manage sending/viewing/editing notes

