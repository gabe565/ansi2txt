# ansi2txt

Command-line tool that converts text containing ANSI control codes into plain ASCII text. It works as a filter, reading from stdin, removing
all ANSI codes, and sending the output to stdout.

This repo was inspired by [kilobyte/colorized-logs](https://github.com/kilobyte/colorized-logs).

## Installation

```shell
go install github.com/gabe565/ansi2txt@latest
```
