# GoCommit

This creates commits with automatically generated messages from the `git diff --cached` using ollama

## Installation

### Precompiled Binaries

You can download the precompiled binaries from the [releases](https://github.com/vicendominguez/your-repo/releases) page on GitHub.

#### Linux (amd64)

1. Download the Linux binary from the releases page.
2. Make the binary executable:

   ```bash
   chmod +x gocommit-vX.X.X-linux-amd64
   sudo mv gocommit-vX.X.X-linux-amd64 /usr/local/bin/gocommit
   ```

#### macOS (darwin/amd64 and darwin/arm64)

1. Download the macOS binary from the releases page.
2. Make the binary executable:

   ```bash
   chmod +x gocommit-vX.X.X-darwin-arm64
   sudo mv gocommit-vX.X.X-darwin-arm64 /usr/local/bin/gocommit
   ```

#### Debian Package (amd64)

1. Download the .deb package from the releases page.

2. Install the package using dpkg:

```bash
sudo dpkg -i gocommit_X.X.X_amd64.deb
```

## Usage

```
ds-1❯ vim README.md
ds-1❯ git add README.md
ds-1❯ gocommit
 SUCCESS  Commit created successfully: [ds-1] Automated commit message generation via ollama added to README
 INFO  Commit created with hash: fbb5bdf4169bef87658afe3e15317f547d2a780c
ds-1❯
```

```
gocommit --help
Usage: gocommit [options]
  -help
    	Show available flags
  -no-prefix
    	Disable the prefix in the commit message
  -prefix string
    	Define a custom prefix for the commit message
  -version
    	Show the version of the application

```


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

## Building from Source

  ```bash
  git clone https://github.com/your-username/your-repo.git
  cd your-repo
  go build -o gocommit ./cmd/gocommit
  sudo mv gocommit /usr/local/bin/ 
  ```
