# GoCommit

## Installation

working on it....

## Usage

1. **Stage Your Changes**:
   - Use `git add` to stage your changes:
     ```bash
     git add -A
     ```

2. **Run GoCommit**:
   - Invoke the tool to automatically generate and create a commit:
     ```bash
     gocommit
     ```

## Rules

- Ensure you have a branch checked out before running `gocommit`. If you don't have a branch, create one using:
  ```bash
  git branch <branch-name>
  git checkout <branch-name>

## Work in progress

BUGS to fix:

- [ ] Currently you must be in your git root directory
- [ ] The fucking `git diff --cached` is  not implemented using the `go-git` library. Currently is a shitty exec.
