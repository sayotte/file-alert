# file-alert
`file-alert` watches a directory for new files, and emails them to
a configurable address when they appear.

## Installation
### From source
```sh
go get github.com/sayotte/file-alert
```
### From pre-built binary
Look at recent [releases](https://github.com/sayotte/file-alert/releases)
and download a binary appropriate for your platform (if it exists).

If you want to invoke this from the command line, place the binary
somewhere in your `$PATH` (on Linux/Unix) or `%PATH%` (on Windows).

If you want to invoke it from e.g. Explorer / Finder, simply put
it somewhere where you can find it easily.

## Configuration
See [example-config.yaml](example-config.yaml), reproduced here:
```yaml
smtp:
  username: myname@gmail.com
  password: "SeriouslyExcellentPassphrase"
  host: smtp.gmail.com
  port: 587
message:
  subject: "New file!"
  from: myname@gmail.com
  fromNickname: "Directory watcher bot"
  to: myname@gmail.com
  toNickname: "myself"
watchPath: . # watches the directory the tool is invoked in
```

## Usage
1. Create a config file by copying [example-config.yaml](example-config.yaml) and
   editing to your needs.
2. Run the program. 
   * From the command-line: `file-alert [-config C:\path\to\config.yaml]`
   * From a file explorer / finder / whatever, just double-click the binary.
      * Note that in this case, your config file MUST be named `config.yaml`
        and MUST be located in the same directory as the binary.
