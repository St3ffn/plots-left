# plots-left

[![Release](https://img.shields.io/github/v/release/St3ffn/plots-left)](https://github.com/St3ffn/plots-left/releases)
[![CI](https://github.com/St3ffn/plots-left/actions/workflows/ci.yml/badge.svg)](https://github.com/St3ffn/plots-left/actions/workflows/ci.yml)
[![License](https://img.shields.io/github/license/st3ffn/plots-left)](/LICENSE)
[![GO](https://img.shields.io/github/go-mod/go-version/St3ffn/plots-left)](https://golang.org/)

![count](https://media.giphy.com/media/3owzW5c1tPq63MPmWk/giphy.gif)

Tiny CLI tool to find out how many chia plots will still fit on your hard disk. In other word to say: 
How many chia plots do I have left on my hard disk. The tool will work fine on all unix systems (unix, linux, macos)

## Getting started

### Pre-requisites

- Linux, MacOS or other Unix based System
- `git` installed
- `go 1.16` installed

### Installation 

Clone the repository

```shell
git clone https://github.com/St3ffn/plots-left.git
cd plots-left
```

Build the binary

```shell
go build
```

Now you are ready to go.

### Usage

If your chia executable is located in `$HOME/chia-blockchain/venv/bin/chia`, you can simply run:
```bash
# assumes chia executable is located in $HOME/chia-blockchain/venv/bin/chia
# removes all connections from mars
> plots-left mars
```
To specify a custom path to your chia executable use `--chiaexec` or `-e`
```bash
# custom defined chia executable
# removes all connections from "elon on mars"
> plots-left -e /home/steffen/chia-blockchain/venv/bin/chia elon on mars
```
You can also add another filter to remove all nodes which have a lower or equal down speed (in MiB) than specified. 
This will be independent of the location filter. It can be done via `--down-threshold` or `-d`.
```bash
# custom defined chia executable
# remove all connections with down speed lower or equal than 1.52 MiB
# removes all connections from "elon on mars"
> plots-left -e /home/steffen/chia-blockchain/venv/bin/chia -d 1.52 elon on mars
```
Call with `--help` or `-h` to see the help page 
```bash
> plots-left -h

NAME:
   plots-left - remove unwanted connections from your Chia Node based on Geo IP Location.

USAGE:
   plots-left [-e CHIA-EXECUTABLE] [-d DOWN-THRESHOLD] LOCATION
   plots-left -e /chia-blockchain/venv/bin/chia -d 0.2 mars

DESCRIPTION:
   Tool will lookup connections via 'chia show -c', get ip locations via geoiplookup and remove nodes from specified LOCATION via 'chia show -r'

GLOBAL OPTIONS:
   --chia-exec CHIA-EXECUTABLE, -e CHIA-EXECUTABLE     CHIA-EXECUTABLE. normally located inside the bin folder of your venv directory (default: $HOME/chia-blockchain/venv/bin/chia)
   --down-threshold DOWN-THRESHOLD, -d DOWN-THRESHOLD  DOWN-THRESHOLD defines the additional filter for minimal down speed in MiB for filtering. (default: not active)
   --help, -h                                          show help (default: false)

COPYRIGHT:
   GNU GPLv3
```

## Kind gestures

If you like the tool, and you are open for a kind gesture. Thanks in advance. 

- XCH Address: xch18s8r9v4kpwdx2y8jks5ma4g2rmff0h9dtr5nkc6zmnk5kj6v0faqer6k9v

