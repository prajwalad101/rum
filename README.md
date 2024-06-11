# rum

## Overview

**rum** is a command-line interface (CLI) application designed to automate the process of creating pull requests directly in your web browser. At it's current state, it does only one thing and it does it well; opening your pull requests in the default web browser. However, it only provides support for codecommit repositories, but I plan to add support for more repositories soon.

## Installation

To install **rum**, you need to have [Go](https://golang.org/) installed on your system. Follow these steps to install **rum**:

### Prerequisites

1. **Install Go**: If you haven't installed Go yet, you can download and install it from the [official Go website](https://golang.org/dl/).

2. **Set up Go environment**: Ensure your Go environment is properly set up. Add Go's bin directory to your system's PATH.

### Steps

1. **Download the source code**:

    Clone the repository using `git`:

    ```bash
    git clone https://github.com/prajwal/rum.git
    cd rum
    ```

2. **Build the application**:

    Run the following command to build the application:

    ```bash
    go build -o rum
    ```

    This will generate an executable file named `rum` in the current directory.

3. **Install the application**:

    To install **rum** globally, move the executable to a directory that's in your system's PATH. For example:

    ```bash
    sudo mv rum /usr/local/bin/
    ```

4. **Verify the installation**:

    Verify that **rum** is installed correctly by running:

    ```bash
    rum --version
    ```

    You should see the version information for **rum**.
