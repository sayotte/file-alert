# file-alert
`file-alert` watches a directory for new files, and emails them to
a configurable address when they appear.

## Installation
### From source
```sh
go get github.com/sayotte/file-alert
```

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
2. `file-alert -config C:\path\to\config.yaml`
