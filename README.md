# saferm
A command-line tool for removing files from a directory in a safe fashion. Written in Go.

## Installation
First, install or update the `saferm` binary into you `$GOPATH` like so:
```bash
$ go get -u github.com/dansyuqri/saferm
```

## Usage
You can immediately use the `saferm` command on your cli.
```bash
$ saferm
```
By default, the `saferm` works on the current working directory. In order to specify another directory, you will need to supplement the `-p` flag like so:
```bash
$ saferm -p /home/user/Downloads
```

The above command immediately starts giving you a filename for you to input in order to delete the file:

```bash
$ saferm -p /home/user/Downloads
Enter filename to confirm deletion:
testfile.txt

```

Once you have entered the exact filename, it will proceed to delete the file. This painstaking process of typing the exact name of the file will assure you that the deletion is by choice, not accident (similar to how deletion of a reposity on GitHub works).