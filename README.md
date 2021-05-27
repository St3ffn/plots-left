# plots-left

[![Release](https://img.shields.io/github/v/release/St3ffn/plots-left)](https://github.com/St3ffn/plots-left/releases)
[![CI](https://github.com/St3ffn/plots-left/actions/workflows/ci.yml/badge.svg)](https://github.com/St3ffn/plots-left/actions/workflows/ci.yml)
[![License](https://img.shields.io/github/license/st3ffn/plots-left)](/LICENSE)
[![GO](https://img.shields.io/github/go-mod/go-version/St3ffn/plots-left)](https://golang.org/)

![count](https://media.giphy.com/media/3owzW5c1tPq63MPmWk/giphy.gif)

Tiny CLI tool to find out how many chia plots will still fit on your hard disk. In other words to say: 
How many chia plots do I have left on my hard disk. The tool will work fine on all unix based systems (unix, linux, macos)

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

```bash
# find out how many chia plots will still fit on /path/to/plots
> plots-left /path/to/plots
```

To specify the amount of plots to reserve use `--reserve` or `-r`
```bash
# find out how many chia plots will still fit on /path/to/plots
# reserve 1 chia plot
> plots-left -r 1 /path/to/plots
```
Call with `--help` or `-h` to see the help page
```bash
> plots-left -h

NAME:
   plots-left - find out how many plots will still fit on your hard disk

USAGE:
   plots-left [-r RESERVE] PATH
   plots-left -r 1 /plots/nas1

DESCRIPTION:
   Tool will perform the following calculation (AVAILABLE_DISK_SPACE/SINGLE_PLOT_SIZE)-RESERVED_PLOTS.

GLOBAL OPTIONS:
   --reserve RESERVE, -r RESERVE  RESERVE. the amount of plots to reserve. (default: 0)
   --help, -h                     show help (default: false)

COPYRIGHT:
   GNU GPLv3
```

## Kind gestures

If you like the tool, and you are open for a kind gesture. Thanks in advance. 

- XCH Address: xch18s8r9v4kpwdx2y8jks5ma4g2rmff0h9dtr5nkc6zmnk5kj6v0faqer6k9v

