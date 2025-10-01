# Var Saver

## Overview
This is a small CLI solution to store and retrieve variables. This came about from the pain of having to
look through countless `.env` and config files whenever I needed the URL to X or Y API and similar. This
is a quick and rough solution for a CLI program to handle this. This solution allows you to store variables
organized by projects, and associated with environments (dev/stage/prod/etc).

This program is meant only for local use, since the vars are stored in plaintext, and I just made it to
quickly solve a quick pain point for myself. However, anyone who wants to is free to use this.

## Installation
To install the program, you must first install GoLang developer tools to compile it. Then, clone the repo and
run `go install`. To verify that it has been correctly installed, run `var-saver version`, and your output should
look something like this:
```
---------------------------------------
*Metadata*
Application: var-saver
Version: v0.0.1
Variables Last Updated:
---------------------------------------
```

## Usage

### Adding Variables
To add new a new variable, you can use the `create` command:
```sh
var-saver create --name variable-name --project some-project --value some-value --environment dev

# Shorthand version
var-saver create -n variable-name -p some-project --value some-value -e dev
#
```
The `project` and `environment` flags are optional and default to `common` and  `default` respectively. Variables are uniquely
identified by their `project`, `name`, and `environment`. To overwrite an existing variable, you can pass the `--overwrite` flag.

### List Existing Variables and Projects
To list existing projects, use
```sh
var-saver list projects
```
This will list only the project names and total variables in each. You can also list them with all their associated variables
by passing the `--detail` or `-d` flag.

You can list the variables in a specific project with
```sh
var-saver list variables --project my-project
```

### Deleting
To delete a project and all its associated variables, use
```sh
var-saver delete project --name my-project
```

To delete a specific variable use
```sh
var-saver delete variable --name my-var --environment dev --project my-project
```
