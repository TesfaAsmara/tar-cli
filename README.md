# TAR-CLI

A simple command-line interface tool (CLI) for the tar command in a text-based user interface (TUI). Choose from straightforward configuration options to create commands, which will then be automatically copied to your clipboard.

## Installation

#### Build from source

```sh
$ git clone github.com/TesfaAsmara/tar-cli
$ cd tar-cli
$ make all
```

## Usage
Run `make run` in your terminal to start the app.

## Navigation
| Key                      | Description                            |
| -----------------------  | -------------------------------------- |
| <kbd> up </kbd>          | Move up in the current section         |
| <kbd> down </kbd>        | Move down in the current section       |
| <kbd> Enter </kbd>       | Select/toggle current item             |
| <kbd> q </kbd>           | quit                                   |

## Built with
- [Bubbletea](https://github.com/charmbracelet/bubbletea)
- [Lipgloss](https://github.com/charmbracelet/lipgloss)
- [Urfave/cli](https://github.com/urfave/cli/v2)
- [Clipboard](https://github.com/atotto/clipboard)

## Inspired by
- [CHMOD-CLI](https://github.com/Mayowa-Ojo/chmod-cli/tree/main)
- [Bubble Tea Basics Tutorial](https://github.com/charmbracelet/bubbletea/tree/master/tutorials/basics)
- [Lipgloss Tutorial](https://github.com/charmbracelet/lipgloss/blob/master/README.md)
