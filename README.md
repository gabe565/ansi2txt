# ansi2txt

[![Build](https://github.com/gabe565/ansi2txt/actions/workflows/build.yaml/badge.svg)](https://github.com/gabe565/ansi2txt/actions/workflows/build.yaml)
[![Go Report Card](https://goreportcard.com/badge/github.com/gabe565/ansi2txt)](https://goreportcard.com/report/github.com/gabe565/ansi2txt)

Command-line tool that converts text containing ANSI control codes into plain ASCII text. It works as a filter, reading from stdin or a file, removing all ANSI codes, and sending the output to stdout.

This repo was inspired by [kilobyte/colorized-logs](https://github.com/kilobyte/colorized-logs).

## Installation

```shell
go install github.com/gabe565/ansi2txt@latest
```
