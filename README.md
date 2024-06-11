# rum

## Overview

**rum** is a command-line interface (CLI) application designed to automate the process of creating pull requests directly in your web browser. At it's current state, it does only one thing and it does it well; opening your pull requests in the default web browser. However, it only provides support for codecommit repositories, but I plan to add support for more repositories soon.

## Usage

To open pull request from a source to destination branch.
```
rum pr <src> <destination>
```

To open a pull request from the currently checked out branch to a destination branch.

```
rum pr <destination>
```

## Installation

To install **rum**, you need to have [Go](https://golang.org/) installed on your system. Follow these steps to install **rum**:

```
go install github.com/prajwalad101/rum
```
This installs the rum binary to $HOME/go on Unix and %USERPROFILE%\go on Windows. Then in a new terminal run

```
rum --help
```
This should display a list of possible commands you can execute.
