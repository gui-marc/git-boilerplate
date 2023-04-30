![Git-Boilerplate](/imgs/cover.png)

# Git Boilerplate CLI Tool

The Git Boilerplate CLI Tool is a command-line interface for creating new projects from a GitHub repository template. This tool provides a simple way to automate the process of setting up new projects by cloning a repository template and customizing it with your own content.

## Installation

### From Source

To install Git Boilerplate from source, first make sure you have Go installed on your system. Then run the following command:

```
$ go get github.com/gui-marc/git-boilerplate
```

This will download the source code and install the `git-boilerplate` binary in your `$GOPATH/bin` directory.

### Using curl

You can also install Git Boilerplate using the following command:

```
$ curl -sSf https://raw.githubusercontent.com/gui-marc/git-boilerplate/main/install.sh | sh
```

This command will download the latest release of Git Boilerplate and install it in `/usr/local/bin`.

## Usage

To use Git Boilerplate, simply run the following command:

```
$ git-boilerplate <template-repo> <project-name>
```

Replace `<template-repo>` with the URL of the GitHub repository template you want to use, and `<project-name>` with the name of the project you want to create.

The Git Boilerplate CLI Tool will clone the template repository into a temporary directory, create a new project directory with the specified name, copy the contents of the template repository into the project directory, remove the `.git` directory from the project directory, initialize a new Git repository in the project directory, and finally remove the temporary directory.

## Contributing

If you find a bug or have an idea for a new feature, please open an issue or submit a pull request on GitHub. All contributions are welcome!

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.