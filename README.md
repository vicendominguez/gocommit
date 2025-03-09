# GoCommit

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
