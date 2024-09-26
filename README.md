# ansi2txt
[![GitHub release (latest by date)](https://img.shields.io/github/v/release/gabe565/ansi2txt)](https://github.com/gabe565/ansi2txt/releases)
[![Build](https://github.com/gabe565/ansi2txt/actions/workflows/build.yaml/badge.svg)](https://github.com/gabe565/ansi2txt/actions/workflows/build.yaml)
[![Go Report Card](https://goreportcard.com/badge/github.com/gabe565/ansi2txt)](https://goreportcard.com/report/github.com/gabe565/ansi2txt)

Command-line tool that converts text containing ANSI control codes into plain ASCII text. It works as a filter, reading from stdin or a file, removing all ANSI codes, and sending the output to stdout.

This repo was inspired by [kilobyte/colorized-logs](https://github.com/kilobyte/colorized-logs).

## Installation

### APT (Ubuntu, Debian)

<details>
  <summary>Click to expand</summary>

1. If you don't have it already, install the `ca-certificates` package
   ```shell
   sudo apt install ca-certificates
   ```

2. Add gabe565 apt repository
   ```
   echo 'deb [trusted=yes] https://apt.gabe565.com /' | sudo tee /etc/apt/sources.list.d/gabe565.list
   ```

3. Update apt repositories
   ```shell
   sudo apt update
   ```

4. Install ansi2txt
   ```shell
   sudo apt install ansi2txt
   ```
</details>

### RPM (CentOS, RHEL)

<details>
  <summary>Click to expand</summary>

1. If you don't have it already, install the `ca-certificates` package
   ```shell
   sudo dnf install ca-certificates
   ```

2. Add gabe565 rpm repository to `/etc/yum.repos.d/gabe565.repo`
   ```ini
   [gabe565]
   name=gabe565
   baseurl=https://rpm.gabe565.com
   enabled=1
   gpgcheck=0
   ```

3. Install ansi2txt
   ```shell
   sudo dnf install ansi2txt
   ```
</details>

### AUR (Arch Linux)

<details>
  <summary>Click to expand</summary>

Install [ansi2txt-bin](https://aur.archlinux.org/packages/ansi2txt-bin) with your [AUR helper](https://wiki.archlinux.org/index.php/AUR_helpers) of choice.
</details>

### Homebrew (macOS, Linux)

<details>
  <summary>Click to expand</summary>

Install ansi2txt from [gabe565/homebrew-tap](https://github.com/gabe565/homebrew-tap):
```shell
brew install gabe565/tap/ansi2txt
```
</details>

### Manual Installation

<details>
  <summary>Click to expand</summary>

1. Download and run the [latest release binary](https://github.com/gabe565/ansi2txt/releases/latest) for your system and architecture.
2. Extract the binary and place it in the desired directory.
</details>
